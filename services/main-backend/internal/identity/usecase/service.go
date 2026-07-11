package usecase

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"log/slog"
	"strings"
	"time"

	"shopwise/retail/internal/identity/domain"
	"shopwise/retail/internal/platform/apperror"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type Service struct {
	users      UserAuthRepository
	identities AuthIdentityRepository
	refresh    RefreshTokenRepository
	verify     EmailVerificationRepository
	jwt        *JWTService
	refreshTTL time.Duration
	google     GoogleTokenExchanger
	devMode    bool
	log        *slog.Logger
}

func NewService(
	users UserAuthRepository,
	identities AuthIdentityRepository,
	refresh RefreshTokenRepository,
	verify EmailVerificationRepository,
	jwt *JWTService,
	refreshTTL time.Duration,
	google GoogleTokenExchanger,
) *Service {
	return &Service{
		users:      users,
		identities: identities,
		refresh:    refresh,
		verify:     verify,
		jwt:        jwt,
		refreshTTL: refreshTTL,
		google:     google,
	}
}

func (s *Service) WithDevLogging(devMode bool, log *slog.Logger) *Service {
	s.devMode = devMode
	s.log = log
	return s
}

type RegisterInput struct {
	Email    string
	Password string
	Name     string
}

func (s *Service) Register(ctx context.Context, input RegisterInput) error {
	email := strings.TrimSpace(input.Email)
	name := strings.TrimSpace(input.Name)
	if email == "" || name == "" || len(input.Password) < 8 {
		return apperror.Validation("invalid registration input", nil)
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		return apperror.Internal("failed to hash password", err)
	}

	user, err := s.users.CreateWithPassword(ctx, email, name, string(hash))
	if err != nil {
		if errors.Is(err, domain.ErrDuplicateEmail) {
			return apperror.Conflict("email already registered", err)
		}
		return apperror.Internal("failed to create user", err)
	}

	rawToken, tokenHash, expiresAt, err := newVerificationToken()
	if err != nil {
		return apperror.Internal("failed to create verification token", err)
	}
	if s.devMode && s.log != nil {
		s.log.Info("email verification token (development only)",
			slog.String("email", email),
			slog.String("token", rawToken),
		)
	}
	if err := s.verify.Create(ctx, user.ID, tokenHash, expiresAt); err != nil {
		return apperror.Internal("failed to store verification token", err)
	}

	return nil
}

type LoginInput struct {
	Email    string
	Password string
}

func (s *Service) Login(ctx context.Context, input LoginInput) (*TokenPair, error) {
	user, err := s.users.GetByEmail(ctx, input.Email)
	if err != nil {
		if errors.Is(err, domain.ErrInvalidCredentials) {
			return nil, apperror.Unauthorized("invalid credentials", err)
		}
		return nil, apperror.Internal("login failed", err)
	}
	if user.PasswordHash == nil {
		return nil, apperror.Unauthorized("invalid credentials", domain.ErrInvalidCredentials)
	}
	if err := bcrypt.CompareHashAndPassword([]byte(*user.PasswordHash), []byte(input.Password)); err != nil {
		return nil, apperror.Unauthorized("invalid credentials", domain.ErrInvalidCredentials)
	}
	return s.issueTokenPair(ctx, user.ID)
}

type RefreshInput struct {
	RefreshToken string
}

func (s *Service) Refresh(ctx context.Context, input RefreshInput) (*TokenPair, error) {
	if strings.TrimSpace(input.RefreshToken) == "" {
		return nil, apperror.Validation("refresh token required", nil)
	}
	now := time.Now().UTC()
	hash := domain.HashToken(input.RefreshToken)
	userID, err := s.refresh.FindValid(ctx, hash, now)
	if err != nil {
		if errors.Is(err, domain.ErrInvalidRefreshToken) {
			return nil, apperror.Unauthorized("invalid refresh token", err)
		}
		return nil, apperror.Internal("refresh failed", err)
	}
	if err := s.refresh.Revoke(ctx, hash, now); err != nil {
		return nil, apperror.Internal("refresh failed", err)
	}
	return s.issueTokenPair(ctx, userID)
}

func (s *Service) GoogleAuthURL(state string) (string, error) {
	if s.google == nil {
		return "", apperror.Internal("google oauth not configured", domain.ErrGoogleNotConfigured)
	}
	return s.google.AuthCodeURL(state), nil
}

func (s *Service) GoogleCallback(ctx context.Context, code string) (*TokenPair, error) {
	if s.google == nil {
		return nil, apperror.Internal("google oauth not configured", domain.ErrGoogleNotConfigured)
	}
	info, err := s.google.Exchange(ctx, code)
	if err != nil {
		return nil, apperror.Unauthorized("google authentication failed", err)
	}
	var verifiedPtr *time.Time
	if info.EmailVerified {
		verifiedAt := time.Now().UTC()
		verifiedPtr = &verifiedAt
	}
	user, err := s.identities.CreateGoogleUser(ctx, info.Email, info.Name, "google", info.Subject, verifiedPtr)
	if err != nil {
		if errors.Is(err, domain.ErrDuplicateEmail) {
			return nil, apperror.Conflict("email already registered", err)
		}
		return nil, apperror.Internal("google sign-in failed", err)
	}
	if info.EmailVerified && user.EmailVerifiedAt == nil {
		if err := s.users.MarkEmailVerified(ctx, user.ID, time.Now().UTC()); err != nil {
			return nil, apperror.Internal("google sign-in failed", err)
		}
	}
	return s.issueTokenPair(ctx, user.ID)
}

func (s *Service) VerifyEmail(ctx context.Context, rawToken string) error {
	token := strings.TrimSpace(rawToken)
	if token == "" {
		return apperror.Validation("verification token required", nil)
	}
	now := time.Now().UTC()
	hash := domain.HashToken(token)
	userID, err := s.verify.FindValidByHash(ctx, hash, now)
	if err != nil {
		switch {
		case errors.Is(err, domain.ErrVerificationTokenExpired),
			errors.Is(err, domain.ErrVerificationTokenUsed),
			errors.Is(err, domain.ErrInvalidVerificationToken):
			return apperror.Validation("invalid or expired verification token", err)
		default:
			return apperror.Internal("email verification failed", err)
		}
	}
	user, err := s.users.GetByID(ctx, userID)
	if err != nil {
		return apperror.Internal("email verification failed", err)
	}
	if user.EmailVerifiedAt != nil {
		_ = s.verify.MarkUsed(ctx, hash, now)
		return nil
	}
	if err := s.verify.MarkUsed(ctx, hash, now); err != nil {
		return apperror.Internal("email verification failed", err)
	}
	if err := s.users.MarkEmailVerified(ctx, userID, now); err != nil {
		return apperror.Internal("email verification failed", err)
	}
	return nil
}

func (s *Service) issueTokenPair(ctx context.Context, userID uuid.UUID) (*TokenPair, error) {
	accessToken, expiresIn, err := s.jwt.IssueAccessToken(userID)
	if err != nil {
		return nil, apperror.Internal("failed to issue access token", err)
	}
	rawRefresh, err := newRefreshToken()
	if err != nil {
		return nil, apperror.Internal("failed to issue refresh token", err)
	}
	expiresAt := time.Now().UTC().Add(s.refreshTTL)
	if err := s.refresh.Store(ctx, userID, domain.HashToken(rawRefresh), expiresAt); err != nil {
		return nil, apperror.Internal("failed to store refresh token", err)
	}
	return &TokenPair{
		AccessToken:  accessToken,
		ExpiresIn:    expiresIn,
		RefreshToken: rawRefresh,
	}, nil
}

func newRefreshToken() (string, error) {
	buf := make([]byte, 32)
	if _, err := rand.Read(buf); err != nil {
		return "", err
	}
	return base64.RawURLEncoding.EncodeToString(buf), nil
}

func newVerificationToken() (raw string, hash string, expiresAt time.Time, err error) {
	buf := make([]byte, 32)
	if _, err = rand.Read(buf); err != nil {
		return "", "", time.Time{}, err
	}
	raw = base64.RawURLEncoding.EncodeToString(buf)
	hash = domain.HashToken(raw)
	expiresAt = time.Now().UTC().Add(24 * time.Hour)
	return raw, hash, expiresAt, nil
}
