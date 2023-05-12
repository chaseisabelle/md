package mdlog

import (
	"context"
	"github.com/aws/aws-xray-sdk-go/xray"
	"github.com/chaseisabelle/md/mdctx"
	"github.com/chaseisabelle/md/mderr"
)

// WithPersistedMetadata applies the given metadata to
// all log entries
// useful for adding things like "env" or "app-name" that would be
// in all log entries
func WithPersistedMetadata(lgr Logger, pmd map[string]any) Logger {
	if pmd == nil || len(pmd) == 0 {
		return lgr
	}

	mod := func(md map[string]any) map[string]any {
		if md == nil {
			return pmd
		}

		for key, val := range pmd {
			md[key] = val
		}

		return md
	}

	return WithMods(lgr, func(ctx context.Context, err error, md map[string]any, f ErrFunc) {
		f(ctx, err, mod(md))
	}, func(ctx context.Context, msg string, md map[string]any, f MsgFunc) {
		f(ctx, msg, mod(md))
	})
}

// WithErrorTrace applies error trace logger middleware
// the returned Logger will inject the error trace into the
// metadata payload from the context with the key
// if key == "" then "error-trace" is used
func WithErrorTrace(lgr Logger, key string) Logger {
	if key == "" {
		key = "error-trace"
	}

	return WithErrMod(lgr, func(ctx context.Context, err error, md map[string]any, f ErrFunc) {
		if err != nil {
			if md == nil {
				md = map[string]any{}
			}

			md[key] = mderr.Stack(err)
		}

		f(ctx, err, md)
	})
}

// WithRequestID applies request id logger middleware
// the returned Logger will inject the request id into the
// metadata payload from the context with the key
// if key == "" then "request-id" is used
func WithRequestID(lgr Logger, key string) Logger {
	if key == "" {
		key = "request-id"
	}

	mod := func(ctx context.Context, md map[string]any) map[string]any {
		if ctx == nil {
			return md
		}

		rid := mdctx.RequestID(ctx)

		if rid == "" {
			return md
		}

		if md == nil {
			md = map[string]any{}
		}

		md[key] = rid

		return md
	}

	return WithMods(lgr, func(ctx context.Context, err error, md map[string]any, f ErrFunc) {
		f(ctx, err, mod(ctx, md))
	}, func(ctx context.Context, msg string, md map[string]any, f MsgFunc) {
		f(ctx, msg, mod(ctx, md))
	})
}

// WithAWSXRayTraceID applies aws xray logger middleware
// the returned Logger will inject aws xray's trace id into the
// metadata payload with the key
// if key == "" then "trace-id" is used
func WithAWSXRayTraceID(lgr Logger, key string) Logger {
	if key == "" {
		key = "trace-id"
	}

	mod := func(ctx context.Context, md map[string]any) map[string]any {
		if ctx == nil {
			return md
		}

		tid := xray.TraceID(ctx)

		if tid == "" {
			return md
		}

		if md == nil {
			md = map[string]any{}
		}

		md[key] = tid

		return md
	}

	return WithMods(lgr, func(ctx context.Context, err error, md map[string]any, f ErrFunc) {
		f(ctx, err, mod(ctx, md))
	}, func(ctx context.Context, msg string, md map[string]any, f MsgFunc) {
		f(ctx, msg, mod(ctx, md))
	})
}
