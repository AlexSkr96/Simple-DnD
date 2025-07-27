package middleware

import (
	"context"
	"github.com/AlexSkr96/Simple-DnD/internal/models"
	"github.com/AlexSkr96/Simple-DnD/internal/services/auth"
	httpint "github.com/AlexSkr96/Simple-DnD/pkg/http"
	"github.com/AlexSkr96/Simple-DnD/pkg/logging"
	"net/http"
	"strings"
)

const (
	pathDocs     = "/docs"
	pathOpenAPI  = "/openapi.json"
	pathRegister = "/api/v1/auth/register"
	pathLogin    = "/api/v1/auth/login"
)

type userContextKey struct{}

func GetUserFromContext(ctx context.Context) (*models.User, bool) {
	user, ok := ctx.Value(userContextKey{}).(*models.User)
	return user, ok
}

func SetUserInContext(ctx context.Context, user *models.User) context.Context {
	return context.WithValue(ctx, userContextKey{}, user)
}

func AuthenticateEndpoint(logger logging.Logger, authService *auth.Service, origin string) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if isPublicEndpoint(r.URL.Path) {
				next.ServeHTTP(w, r)
				return
			}

			reqID := r.Header.Get("X-Request-ID")
			authHeader := r.Header.Get("Authorization")

			if authHeader == "" {
				httpint.WriteErrorResponse(logger, w, reqID, origin, "Authorization header is missing", http.StatusUnauthorized)
				return
			}

			if !strings.HasPrefix(authHeader, "Bearer ") {
				httpint.WriteErrorResponse(logger, w, reqID, origin, "Invalid authorization header format", http.StatusUnauthorized)
				return
			}

			token := strings.TrimPrefix(authHeader, "Bearer ")
			if token == "" {
				httpint.WriteErrorResponse(logger, w, reqID, origin, "Token is missing", http.StatusUnauthorized)
				return
			}

			user, err := authService.ValidateToken(r.Context(), token)
			if err != nil {
				logger.WithContext(r.Context()).WithError(err).Warning("Invalid token")
				httpint.WriteErrorResponse(logger, w, reqID, origin, "Invalid or expired token", http.StatusUnauthorized)
				return
			}

			// Добавляем пользователя в контекст
			ctx := SetUserInContext(r.Context(), user)
			r = r.WithContext(ctx)

			next.ServeHTTP(w, r)
		})
	}
}

// isPublicEndpoint проверяет, является ли эндпоинт публичным
func isPublicEndpoint(path string) bool {
	publicPaths := []string{
		pathDocs,
		pathOpenAPI,
		pathRegister,
		pathLogin,
	}

	for _, publicPath := range publicPaths {
		if path == publicPath {
			return true
		}
	}

	return false
}
