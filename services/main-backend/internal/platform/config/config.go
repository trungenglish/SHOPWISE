package config

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/joho/godotenv"
)

type Config struct {
	Port           string
	GinMode        string
	AllowedOrigins []string
	DatabaseURL    string
	RedisURL       string
	LogLevel       string
	AppName        string

	JWTSecret      string
	JWTAccessTTL   time.Duration
	JWTRefreshTTL  time.Duration
	GoogleClientID string
	GoogleSecret   string
	GoogleRedirect string
	WebAppURL      string

	AIRuntimeURL string
	ZaloOAID     string
	ZaloAPIToken string

	SMTPHost string
	SMTPPort string
	SMTPUser string
	SMTPPass string
	SMTPFrom string

	StoragePath string

	LangfusePublicKey string
	LangfuseSecretKey string
	LangfuseHost      string

	CheckoutAuthBypass bool
}

func Load() (*Config, error) {
	loadEnvFiles()

	accessTTL, err := time.ParseDuration(getEnv("JWT_ACCESS_TTL", "15m"))
	if err != nil {
		return nil, fmt.Errorf("JWT_ACCESS_TTL: %w", err)
	}
	refreshTTL, err := time.ParseDuration(getEnv("JWT_REFRESH_TTL", "168h"))
	if err != nil {
		return nil, fmt.Errorf("JWT_REFRESH_TTL: %w", err)
	}
	checkoutAuthBypass, err := strconv.ParseBool(getEnv("CHECKOUT_AUTH_BYPASS", "false"))
	if err != nil {
		return nil, fmt.Errorf("CHECKOUT_AUTH_BYPASS: %w", err)
	}
	ginMode := getEnv("GIN_MODE", "debug")

	cfg := &Config{
		Port:           getEnv("PORT", "18080"),
		GinMode:        ginMode,
		LogLevel:       getEnv("LOG_LEVEL", "info"),
		AppName:        getEnv("APP_NAME", "server"),
		DatabaseURL:    os.Getenv("DATABASE_URL"),
		RedisURL:       os.Getenv("REDIS_URL"),
		JWTSecret:      os.Getenv("JWT_SECRET"),
		JWTAccessTTL:   accessTTL,
		JWTRefreshTTL:  refreshTTL,
		GoogleClientID: os.Getenv("GOOGLE_CLIENT_ID"),
		GoogleSecret:   os.Getenv("GOOGLE_CLIENT_SECRET"),
		GoogleRedirect: getEnv("GOOGLE_REDIRECT_URI", "http://localhost:18080/api/v1/identity/google/callback"),
		WebAppURL:      getEnv("WEB_APP_URL", "http://localhost:3001"),
		AIRuntimeURL:   getEnv("AI_RUNTIME_URL", "http://localhost:8000"),
		ZaloOAID:       os.Getenv("ZALO_OA_ID"),
		ZaloAPIToken:   os.Getenv("ZALO_API_TOKEN"),
		SMTPHost:       getEnv("SMTP_HOST", "localhost"),
		SMTPPort:       getEnv("SMTP_PORT", "1025"),
		SMTPUser:       os.Getenv("SMTP_USER"),
		SMTPPass:       os.Getenv("SMTP_PASS"),
		SMTPFrom:       getEnv("SMTP_FROM", "noreply@shopwise.local"),
		StoragePath:    getEnv("STORAGE_PATH", "./storage"),

		LangfusePublicKey: os.Getenv("LANGFUSE_PUBLIC_KEY"),
		LangfuseSecretKey: os.Getenv("LANGFUSE_SECRET_KEY"),
		LangfuseHost:      getEnv("LANGFUSE_HOST", "https://cloud.langfuse.com"),

		CheckoutAuthBypass: checkoutAuthBypass && strings.EqualFold(ginMode, "debug"),
	}

	origins := getEnv("ALLOWED_ORIGINS", "http://localhost:3001")
	cfg.AllowedOrigins = splitAndTrim(origins)

	if cfg.DatabaseURL == "" {
		return nil, fmt.Errorf("DATABASE_URL is required")
	}
	if cfg.RedisURL == "" {
		return nil, fmt.Errorf("REDIS_URL is required")
	}
	if cfg.JWTSecret == "" {
		return nil, fmt.Errorf("JWT_SECRET is required")
	}

	return cfg, nil
}

func loadEnvFiles() {
	_ = godotenv.Load(".env")

	if os.Getenv("DATABASE_URL") == "" || os.Getenv("REDIS_URL") == "" {
		_ = godotenv.Load(".env.example")
	}
}

func getEnv(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}

func splitAndTrim(value string) []string {
	parts := strings.Split(value, ",")
	result := make([]string, 0, len(parts))
	for _, part := range parts {
		trimmed := strings.TrimSpace(part)
		if trimmed != "" {
			result = append(result, trimmed)
		}
	}
	return result
}

func (c *Config) GoogleOAuthEnabled() bool {
	return c.GoogleClientID != "" && c.GoogleSecret != ""
}
