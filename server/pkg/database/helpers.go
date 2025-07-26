package database

import "context"

type ctxQueryName struct{}

func WithQueryName(ctx context.Context, name string) context.Context {
	return context.WithValue(ctx, ctxQueryName{}, name)
}

func GetQueryName(ctx context.Context, fallback string) string {
	name := ctx.Value(ctxQueryName{})
	if name == nil {
		return fallback
	}

	return name.(string) // nolint: forcetypeassert
}
