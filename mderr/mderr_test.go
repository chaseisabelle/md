package mderr_test

import (
	"fmt"
	"github.com/chaseisabelle/mderr"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMDErr(t *testing.T) {
	err0 := mderr.New("error 0", map[string]any{
		"foo": "bar",
	})

	as0, is0 := mderr.AsIs(err0)

	assert.NotNil(t, err0)
	assert.True(t, mderr.Is(err0))
	assert.NotNil(t, mderr.As(err0))
	assert.True(t, is0)
	assert.NotNil(t, as0)
	assert.Equal(t, "error 0", as0.Error())
	assert.Equal(t, "error 0", mderr.Error(err0))
	assert.Equal(t, "error 0", as0.Message())
	assert.Equal(t, "error 0", mderr.Message(err0))
	assert.Nil(t, as0.Cause())
	assert.Nil(t, mderr.Cause(err0))
	assert.Equal(t, map[string]any{"foo": "bar"}, as0.Metadata())
	assert.Equal(t, map[string]any{"foo": "bar"}, mderr.Metadata(err0))
	assert.Equal(t, err0, as0.Root())
	assert.Equal(t, err0, mderr.Root(err0))
	assert.Equal(t, 1, as0.Depth())
	assert.Equal(t, 1, mderr.Depth(err0))
	assert.Equal(t, []error{err0}, as0.Array())
	assert.Equal(t, []error{err0}, mderr.Array(err0))

	err1 := mderr.Wrap(err0, "error 1", map[string]any{
		"pee": "poo",
	})

	as1, is1 := mderr.AsIs(err1)

	assert.NotNil(t, err1)
	assert.True(t, mderr.Is(err1))
	assert.NotNil(t, mderr.As(err1))
	assert.True(t, is1)
	assert.NotNil(t, as1)
	assert.Equal(t, "error 1: error 0", as1.Error())
	assert.Equal(t, "error 1: error 0", mderr.Error(err1))
	assert.Equal(t, "error 1", as1.Message())
	assert.Equal(t, "error 1", mderr.Message(err1))
	assert.Equal(t, err0, as1.Cause())
	assert.Equal(t, err0, mderr.Cause(err1))
	assert.Equal(t, map[string]any{"pee": "poo"}, as1.Metadata())
	assert.Equal(t, map[string]any{"pee": "poo"}, mderr.Metadata(err1))
	assert.Equal(t, err0, as1.Root())
	assert.Equal(t, err0, mderr.Root(err1))
	assert.Equal(t, 2, as1.Depth())
	assert.Equal(t, 2, mderr.Depth(err1))
	assert.Equal(t, []error{err1, err0}, as1.Array())
	assert.Equal(t, []error{err1, err0}, mderr.Array(err1))

	err2 := mderr.Wrap(err1, "error 2", nil)

	as2, is2 := mderr.AsIs(err2)

	assert.NotNil(t, err2)
	assert.True(t, mderr.Is(err2))
	assert.NotNil(t, mderr.As(err2))
	assert.True(t, is2)
	assert.NotNil(t, as2)
	assert.Equal(t, "error 2: error 1: error 0", as2.Error())
	assert.Equal(t, "error 2: error 1: error 0", mderr.Error(err2))
	assert.Equal(t, "error 2", as2.Message())
	assert.Equal(t, "error 2", mderr.Message(err2))
	assert.Equal(t, err1, as2.Cause())
	assert.Equal(t, err1, mderr.Cause(err2))
	assert.Equal(t, map[string]any{}, as2.Metadata())
	assert.Equal(t, map[string]any{}, mderr.Metadata(err2))
	assert.Equal(t, err0, as2.Root())
	assert.Equal(t, err0, mderr.Root(err2))
	assert.Equal(t, 3, as2.Depth())
	assert.Equal(t, 3, mderr.Depth(err2))
	assert.Equal(t, []error{err2, err1, err0}, as2.Array())
	assert.Equal(t, []error{err2, err1, err0}, mderr.Array(err2))
}

func TestError(t *testing.T) {
	err0 := fmt.Errorf("error 0")

	as0, is0 := mderr.AsIs(err0)

	assert.NotNil(t, err0)
	assert.False(t, mderr.Is(err0))
	assert.Nil(t, mderr.As(err0))
	assert.False(t, is0)
	assert.Nil(t, as0)
	assert.Equal(t, "error 0", mderr.Error(err0))
	assert.Equal(t, "error 0", mderr.Message(err0))
	assert.Nil(t, mderr.Cause(err0))
	assert.Nil(t, mderr.Metadata(err0))
	assert.Equal(t, err0, mderr.Root(err0))
	assert.Equal(t, 1, mderr.Depth(err0))
	assert.Equal(t, []error{err0}, mderr.Array(err0))

	err1 := fmt.Errorf("error 1: %w", err0)

	as1, is1 := mderr.AsIs(err1)

	assert.NotNil(t, err1)
	assert.False(t, mderr.Is(err1))
	assert.Nil(t, mderr.As(err1))
	assert.False(t, is1)
	assert.Nil(t, as1)
	assert.Equal(t, "error 1: error 0", mderr.Error(err1))
	assert.Equal(t, "error 1: error 0", mderr.Message(err1))
	assert.Equal(t, err0, mderr.Cause(err1))
	assert.Nil(t, mderr.Metadata(err1))
	assert.Equal(t, err0, mderr.Root(err1))
	assert.Equal(t, 2, mderr.Depth(err1))
	assert.Equal(t, []error{err1, err0}, mderr.Array(err1))

	err2 := fmt.Errorf("error 2: %w", err1)

	as2, is2 := mderr.AsIs(err2)

	assert.NotNil(t, err2)
	assert.False(t, mderr.Is(err2))
	assert.Nil(t, mderr.As(err2))
	assert.False(t, is2)
	assert.Nil(t, as2)
	assert.Equal(t, "error 2: error 1: error 0", mderr.Error(err2))
	assert.Equal(t, "error 2: error 1: error 0", mderr.Message(err2))
	assert.Equal(t, err1, mderr.Cause(err2))
	assert.Nil(t, mderr.Metadata(err2))
	assert.Equal(t, err0, mderr.Root(err2))
	assert.Equal(t, 3, mderr.Depth(err2))
	assert.Equal(t, []error{err2, err1, err0}, mderr.Array(err2))
}

func TestMix(t *testing.T) {
	err0 := fmt.Errorf("error 0")

	as0, is0 := mderr.AsIs(err0)

	assert.NotNil(t, err0)
	assert.False(t, mderr.Is(err0))
	assert.Nil(t, mderr.As(err0))
	assert.False(t, is0)
	assert.Nil(t, as0)
	assert.Equal(t, "error 0", mderr.Error(err0))
	assert.Equal(t, "error 0", mderr.Message(err0))
	assert.Nil(t, mderr.Cause(err0))
	assert.Nil(t, mderr.Metadata(err0))
	assert.Equal(t, err0, mderr.Root(err0))
	assert.Equal(t, 1, mderr.Depth(err0))
	assert.Equal(t, []error{err0}, mderr.Array(err0))

	err1 := mderr.Wrap(err0, "error 1", map[string]any{
		"pee": "poo",
	})

	as1, is1 := mderr.AsIs(err1)

	assert.NotNil(t, err1)
	assert.True(t, mderr.Is(err1))
	assert.NotNil(t, mderr.As(err1))
	assert.True(t, is1)
	assert.NotNil(t, as1)
	assert.Equal(t, "error 1: error 0", as1.Error())
	assert.Equal(t, "error 1: error 0", mderr.Error(err1))
	assert.Equal(t, "error 1", as1.Message())
	assert.Equal(t, "error 1", mderr.Message(err1))
	assert.Equal(t, err0, as1.Cause())
	assert.Equal(t, err0, mderr.Cause(err1))
	assert.Equal(t, map[string]any{"pee": "poo"}, as1.Metadata())
	assert.Equal(t, map[string]any{"pee": "poo"}, mderr.Metadata(err1))
	assert.Equal(t, err0, as1.Root())
	assert.Equal(t, err0, mderr.Root(err1))
	assert.Equal(t, 2, as1.Depth())
	assert.Equal(t, 2, mderr.Depth(err1))
	assert.Equal(t, []error{err1, err0}, as1.Array())
	assert.Equal(t, []error{err1, err0}, mderr.Array(err1))

	err2 := fmt.Errorf("error 2: %w", err1)

	as2, is2 := mderr.AsIs(err2)

	assert.NotNil(t, err2)
	assert.False(t, mderr.Is(err2))
	assert.Nil(t, mderr.As(err2))
	assert.False(t, is2)
	assert.Nil(t, as2)
	assert.Equal(t, "error 2: error 1: error 0", mderr.Error(err2))
	assert.Equal(t, "error 2: error 1: error 0", mderr.Message(err2))
	assert.Equal(t, err1, mderr.Cause(err2))
	assert.Nil(t, mderr.Metadata(err2))
	assert.Equal(t, err0, mderr.Root(err2))
	assert.Equal(t, 3, mderr.Depth(err2))
	assert.Equal(t, []error{err2, err1, err0}, mderr.Array(err2))

	err3 := mderr.Wrap(err2, "error 3", map[string]any{
		"foo": "bar",
	})

	as3, is3 := mderr.AsIs(err3)

	assert.NotNil(t, err3)
	assert.True(t, mderr.Is(err3))
	assert.NotNil(t, mderr.As(err3))
	assert.True(t, is3)
	assert.NotNil(t, as3)
	assert.Equal(t, "error 3: error 2: error 1: error 0", as3.Error())
	assert.Equal(t, "error 3: error 2: error 1: error 0", mderr.Error(err3))
	assert.Equal(t, "error 3", as3.Message())
	assert.Equal(t, "error 3", mderr.Message(err3))
	assert.Equal(t, err2, as3.Cause())
	assert.Equal(t, err2, mderr.Cause(err3))
	assert.Equal(t, map[string]any{"foo": "bar"}, as3.Metadata())
	assert.Equal(t, map[string]any{"foo": "bar"}, mderr.Metadata(err3))
	assert.Equal(t, err0, as3.Root())
	assert.Equal(t, err0, mderr.Root(err3))
	assert.Equal(t, 4, as3.Depth())
	assert.Equal(t, 4, mderr.Depth(err3))
	assert.Equal(t, []error{err3, err2, err1, err0}, as3.Array())
	assert.Equal(t, []error{err3, err2, err1, err0}, mderr.Array(err3))
}
