package middleware

import (
	"context"
	"net/http"
)

func HTTPSMiddleware(enabled bool) Middleware {
	return func(handler http.HandlerFunc) http.HandlerFunc {

		if !enabled {
			return handler
		}
		return func(rw http.ResponseWriter, r *http.Request) {
			if r.URL.Scheme != "https" && r.Header.Get("X-Forwarded-Proto") != "https" {
				url := r.URL
				if url.Host == "" {
					url.Host = r.Host
				}
				url.Scheme = "https"
				rw.Header().Set(
					"Location",
					url.String(),
				)
				rw.WriteHeader(http.StatusPermanentRedirect)
				r = r.WithContext(
					context.WithValue(r.Context(), responseWrittenKey, true),
				)
				handler(rw, r)
			} else {
				handler(rw, r)
			}
		}

	}
}
