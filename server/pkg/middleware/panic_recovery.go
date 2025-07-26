package middleware

import (
	"net/http"
	"runtime"
	"strconv"
	"strings"

	httpint "github.com/AlexSkr96/Simple-DnD/pkg/http"
	"github.com/AlexSkr96/Simple-DnD/pkg/logging"
)

func NewPanicRecoverer(
	logger logging.Logger, origin string,
) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			defer func() {
				rvr := recover()
				if rvr == nil || rvr == http.ErrAbortHandler { //nolint
					return
				}

				stack := make([]byte, 1024) // nolint: mnd

				_, file, line, _ := runtime.Caller(2) // nolint: mnd
				if strings.Contains(file, "panic") {
					_, file, line, _ = runtime.Caller(3) // nolint: mnd
				}

				logger.
					WithField("stack", string(stack[:runtime.Stack(stack, false)])).
					WithField("file", file+":"+strconv.Itoa(line)).
					Errorf("panic '%s'", rvr)

				httpint.WriteErrorResponse(logger, w, r.Header.Get("X-Request-ID"), origin, "Internal server error", http.StatusInternalServerError)
			}()

			next.ServeHTTP(w, r)
		})
	}
}
