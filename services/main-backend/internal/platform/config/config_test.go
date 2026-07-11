package config

import "testing"

func TestLoadEnablesCheckoutAuthBypassOnlyInDebugMode(t *testing.T) {
	tests := []struct {
		name     string
		ginMode  string
		expected bool
	}{
		{name: "debug", ginMode: "debug", expected: true},
		{name: "release", ginMode: "release", expected: false},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Setenv("DATABASE_URL", "postgres://example")
			t.Setenv("REDIS_URL", "redis://example")
			t.Setenv("JWT_SECRET", "test-secret")
			t.Setenv("GIN_MODE", test.ginMode)
			t.Setenv("CHECKOUT_AUTH_BYPASS", "true")

			configuration, err := Load()
			if err != nil {
				t.Fatalf("Load() error = %v", err)
			}
			if configuration.CheckoutAuthBypass != test.expected {
				t.Fatalf("CheckoutAuthBypass = %t, want %t", configuration.CheckoutAuthBypass, test.expected)
			}
		})
	}
}

func TestLoadRejectsInvalidCheckoutAuthBypass(t *testing.T) {
	t.Setenv("DATABASE_URL", "postgres://example")
	t.Setenv("REDIS_URL", "redis://example")
	t.Setenv("JWT_SECRET", "test-secret")
	t.Setenv("GIN_MODE", "debug")
	t.Setenv("CHECKOUT_AUTH_BYPASS", "sometimes")

	if _, err := Load(); err == nil {
		t.Fatal("Load() error = nil, want invalid boolean error")
	}
}
