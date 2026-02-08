package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	t.Run("returns api key when header is valid", func(t *testing.T) {
		h := make(http.Header)
		h.Set("Authorization", "ApiKey abc123")

		key, err := GetAPIKey(h)
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}
		if key != "abc123" {
			t.Fatalf("expected key %q, got %q", "abc123", key)
		}
	})

	t.Run("returns error when header is missing", func(t *testing.T) {
		h := make(http.Header)

		_, err := GetAPIKey(h)
		if err != ErrNoAuthHeaderIncluded {
			t.Fatalf("expected ErrNoAuthHeaderIncluded, got %v", err)
		}
	})
}
