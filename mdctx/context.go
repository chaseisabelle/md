package mdctx

import (
	"context"
)

type Key int

const (
	RequestIDKey Key = iota
)

// SetRequestID sets a request id in the context
// if request id is empty, nothing is set and original context is returned
func SetRequestID(ctx context.Context, rid string) context.Context {
	if rid == "" {
		return ctx
	}

	return context.WithValue(ctx, RequestIDKey, rid)
}

// GetRequestID gets the request id from a context
// empty string if no request id set
func GetRequestID(ctx context.Context) string {
	if ctx == nil {
		return ""
	}

	val := ctx.Value(RequestIDKey)

	if val == nil {
		return ""
	}

	rid, ok := val.(string)

	if !ok {
		return ""
	}

	return rid
}
