package bootstrap

import (
	"encoding/json"
	"github.com/AlexSkr96/Simple-DnD/pkg/errors"
	"github.com/AlexSkr96/Simple-DnD/pkg/logging"
	"github.com/AlexSkr96/Simple-DnD/pkg/middleware"
	"github.com/danielgtaylor/huma/v2"
	"github.com/danielgtaylor/huma/v2/adapters/humachi"
	"github.com/go-chi/chi/v5"
	"net/http"
)

type DnDAPI huma.API

const origin = "dnd-api"

func NewDnDAPIRouter(
	logger logging.Logger,
) DnDAPI {
	r := chi.NewRouter()
	r.Use(
		Custom404Middleware(logger),
		middleware.CORSHandler(),
		middleware.NewPanicRecoverer(logger, origin),
		middleware.AuthenticateEndpoint(logger, origin),
	)

	cfg := huma.DefaultConfig("DnD API", "1.0.0")
	cfg.CreateHooks = nil

	cfg.Transformers = []huma.Transformer{TransformErrorBody()}

	// Переопределяем логику и тело ошибки
	huma.NewError = func(status int, message string, errs ...error) huma.StatusError {
		if status == http.StatusUnprocessableEntity {
			status = http.StatusBadRequest
		}

		details := make([]*huma.ErrorDetail, len(errs))

		for i := range errs {
			if converted, ok := errs[i].(huma.ErrorDetailer); ok {
				details[i] = converted.ErrorDetail()
			} else {
				if errs[i] == nil {
					continue
				}

				details[i] = &huma.ErrorDetail{Message: errs[i].Error()}
			}
		}

		errMsg := &errors.APIError{
			Status: status,
			Title:  http.StatusText(status),
			Detail: message,
			Origin: origin,
			Errors: details,
		}

		return errMsg
	}

	api := humachi.New(r, cfg)

	r.Get("/docs", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		w.Write([]byte(`<!DOCTYPE html>
			<html lang="en">
			<head>
			  <meta charset="utf-8" />
			  <meta name="viewport" content="width=device-width, initial-scale=1" />
			  <meta name="description" content="SwaggerUI" />
			  <title>SwaggerUI</title>
			  <link rel="stylesheet" href="https://unpkg.com/swagger-ui-dist@5.11.0/swagger-ui.css" />
			</head>
			<body>
			<div id="swagger-ui"></div>
			<script src="https://unpkg.com/swagger-ui-dist@5.11.0/swagger-ui-bundle.js" crossorigin></script>
			<script>
			  window.onload = () => {
				window.ui = SwaggerUIBundle({
				  url: '/openapi.json',
				  dom_id: '#swagger-ui',
				});
			  };
			</script>
			</body>
			</html>`,
		))
	})

	return api
}

func TransformErrorBody() huma.Transformer {
	return func(ctx huma.Context, status string, v any) (any, error) {
		if detailedErrors, ok := v.(*errors.APIError); ok {
			detailedErrors.RequestID = ctx.Header("X-Request-Id")

			return detailedErrors, nil
		}

		return v, nil
	}
}

func Custom404Middleware(logger logging.Logger) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			rctx := chi.RouteContext(r.Context())
			tctx := chi.NewRouteContext()

			if !rctx.Routes.Match(tctx, r.Method, r.URL.Path) {
				status := http.StatusNotFound

				errMsg := &errors.APIError{
					Status:    status,
					Title:     http.StatusText(status),
					Detail:    "no matching operation was found",
					Origin:    origin,
					RequestID: r.Header.Get("X-Request-Id"),
				}

				body, err := json.Marshal(errMsg)
				if err != nil {
					logger.Error(err)
					w.WriteHeader(http.StatusNotFound)
					w.Write([]byte("404 page not found")) // nolint: errcheck

					return
				}

				w.Header().Set("Content-Type", "application/json")

				w.WriteHeader(http.StatusNotFound)
				w.Write(body) // nolint: errcheck

				return
			}

			next.ServeHTTP(w, r)
		})
	}
}
