package middleware

// CORSConfig is configuration for CORS headers
import (
	"context"
	"fmt"
	"net/http"
	"strings"
)

type CORSConfig struct {
	AllowedCredentials bool
	AllowedHeaders     []string
	AllowedMethods     []string
	AllowedOrigins     []string
}

// CORSMiddleware adds CORS headers
func CORSMiddleware(config CORSConfig) Middleware {
	allowedCredentials := fmt.Sprintf("%t", config.AllowedCredentials)
	allowedMethods := strings.Join(config.AllowedMethods, ", ")
	allowedHeaders := strings.Join(config.AllowedHeaders, ", ")
	return func(handler http.HandlerFunc) http.HandlerFunc {
		return func(rw http.ResponseWriter, r *http.Request) {
			rw.Header().Add("Access-Control-Allow-Credentials", allowedCredentials)

			rw.Header().Add("Access-Control-Allow-Methods", allowedMethods)

			requestedOrigin := r.Header.Get("Origin")
			if requestedOrigin != "" {
				for _, origin := range config.AllowedOrigins {
					if origin == requestedOrigin || origin == "*" {
						rw.Header().Add("Access-Control-Allow-Origin", requestedOrigin)
						break
					}
				}
			}

			rw.Header().Add("Access-Control-Allow-Headers", allowedHeaders)
			if r.Method == "OPTIONS" {
				r = r.WithContext(
					context.WithValue(r.Context(), responseWrittenKey, true),
				)
			}

			handler(rw, r)
		}
	}
}
