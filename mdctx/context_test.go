package mdctx_test

import (
	"context"
	"github.com/chaseisabelle/md/mdctx"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRequestID(t *testing.T) {
	rid := uuid.New().String()
	ctx := context.Background()

	ctx = mdctx.WithRequestID(ctx, rid)

	assert.Equal(t, rid, mdctx.RequestID(ctx))
}
