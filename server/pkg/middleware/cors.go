package middleware

import (
	"net/http"

	"github.com/go-chi/cors"
)

func CORSHandler() func(http.Handler) http.Handler {
	return cors.AllowAll().Handler
}
