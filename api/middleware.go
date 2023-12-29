// api/middleware.go
package api

import (
	"net/http"
	"os"
	"mux"
)

// Authenticate is a middleware to authenticate requests
func Authenticate(expectedToken string) mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			token := r.Header.Get("Authorization")
			if token != "Bearer "+expectedToken {
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
				return
			}
			next.ServeHTTP(w, r)
		})
	}
}

// GetToken retrieves the token from environment variables or uses a default value
func GetToken() string {
	return getEnv("TOKEN", "your_secret_token")
}

// getEnv retrieves the value of an environment variable or uses a default value
func getEnv(key, defaultValue string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		return defaultValue
	}
	return value
}

