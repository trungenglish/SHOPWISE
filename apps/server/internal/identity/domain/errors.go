package domain

import "errors"

var (
	ErrDuplicateEmail           = errors.New("email already registered")
	ErrInvalidCredentials       = errors.New("invalid credentials")
	ErrInvalidRefreshToken      = errors.New("invalid refresh token")
	ErrGoogleNotConfigured      = errors.New("google oauth not configured")
	ErrInvalidVerificationToken = errors.New("invalid verification token")
	ErrVerificationTokenExpired = errors.New("verification token expired")
	ErrVerificationTokenUsed    = errors.New("verification token already used")
)
