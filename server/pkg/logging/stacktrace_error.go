package logging

import "github.com/pkg/errors"

type StacktraceError interface {
	StackTrace() errors.StackTrace
}

func UnwrapStacktrace(err error) StacktraceError {
	uerr, ok := err.(StacktraceError) //nolint: errorlint
	if ok {
		return uerr
	}

	err = errors.Unwrap(err)
	if err == nil {
		return nil
	}

	return UnwrapStacktrace(err)
}
