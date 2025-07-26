package http

import (
	"encoding/json"
	"github.com/AlexSkr96/Simple-DnD/pkg/logging"
	"net/http"

	"github.com/danielgtaylor/huma/v2"
	"github.com/pkg/errors"

	errpkg "github.com/AlexSkr96/Simple-DnD/pkg/errors"
)

func WriteErrorResponse(logger logging.Logger, writer http.ResponseWriter, reqID, origin, msg string, statusCode int) {
	errResp := errpkg.APIError{
		Status:    statusCode,
		Title:     http.StatusText(statusCode),
		Detail:    msg,
		RequestID: reqID,
		Origin:    origin,
	}

	message, err := json.Marshal(errResp)
	if err != nil {
		logger.Error(errors.WithStack(err))
		http.Error(writer, err.Error(), statusCode)

		return
	}

	writer.Header().Set("Content-Type", "application/problem+json; charset=utf-8")
	writer.WriteHeader(statusCode)

	_, err = writer.Write(message)
	if err != nil {
		logger.Error(errors.WithStack(err))
	}
}

func WriteErrorResponseHuma(ctx huma.Context, logger logging.Logger, reqID, origin, msg string, statusCode int) {
	errResp := errpkg.APIError{
		Status:    statusCode,
		Title:     http.StatusText(statusCode),
		Detail:    msg,
		RequestID: reqID,
		Origin:    origin,
	}

	message, err := json.Marshal(errResp)
	if err != nil {
		logger.Error(errors.WithStack(err))

		ctx.SetStatus(statusCode)

		if _, err := ctx.BodyWriter().Write(message); err != nil {
			logger.Error(errors.WithStack(err))
		}

		return
	}

	ctx.SetHeader("Content-Type", "application/problem+json; charset=utf-8")
	ctx.SetStatus(statusCode)

	if _, err := ctx.BodyWriter().Write(message); err != nil {
		logger.Error(errors.WithStack(err))
	}
}
