package mdlog_test

import (
	"context"
	"github.com/chaseisabelle/md/mdctx"
	"github.com/chaseisabelle/md/mdlog"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestWithPersistedMetadata(t *testing.T) {
	md := map[string]any{
		"peepee": "poopoo",
	}

	f := func(md map[string]any) {
		val, ok := md["peepee"]

		assert.True(t, ok)
		assert.Equal(t, "poopoo", val)
	}

	ef := func(_ context.Context, _ error, md map[string]any) {
		f(md)
	}

	mf := func(_ context.Context, _ string, md map[string]any) {
		f(md)
	}

	var lgr mdlog.Logger

	lgr = &TestLogger{
		FatalFunc: ef,
		ErrorFunc: ef,
		WarnFunc:  mf,
		InfoFunc:  mf,
		DebugFunc: mf,
	}

	lgr = mdlog.WithPersistedMetadata(lgr, md)

	lgr.Fatal(nil, nil, nil)
	lgr.Error(nil, nil, nil)
	lgr.Warn(nil, "", nil)
	lgr.Info(nil, "", nil)
	lgr.Debug(nil, "", nil)
}

func TestWithRequestID(t *testing.T) {
	exp := uuid.New().String()

	f := func(md map[string]any) {
		act, ok := md["request-id"]

		assert.True(t, ok)
		assert.Equal(t, exp, act)
	}

	ef := func(_ context.Context, _ error, md map[string]any) {
		f(md)
	}

	mf := func(_ context.Context, _ string, md map[string]any) {
		f(md)
	}

	var lgr mdlog.Logger

	lgr = &TestLogger{
		FatalFunc: ef,
		ErrorFunc: ef,
		WarnFunc:  mf,
		InfoFunc:  mf,
		DebugFunc: mf,
	}

	ctx := mdctx.WithRequestID(context.Background(), exp)
	lgr = mdlog.WithRequestID(lgr, "")

	lgr.Fatal(ctx, nil, nil)
	lgr.Error(ctx, nil, nil)
	lgr.Warn(ctx, "", nil)
	lgr.Info(ctx, "", nil)
	lgr.Debug(ctx, "", nil)
}
