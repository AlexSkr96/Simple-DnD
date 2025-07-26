package middleware

import (
	httpint "github.com/AlexSkr96/Simple-DnD/pkg/http"
	"github.com/AlexSkr96/Simple-DnD/pkg/logging"
	"net/http"
)

// TODO: на основе чего аутентификация/авторизация - кука, выписанный токен?
func AuthenticateEndpoint(logger logging.Logger, origin string) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			reqID := r.Header.Get("X-Request-ID")
			userIDStr := r.Header.Get("Authorization")

			if userIDStr == "" {
				httpint.WriteErrorResponse(logger, w, reqID, origin, "Authorization header is missing", http.StatusUnauthorized)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}
