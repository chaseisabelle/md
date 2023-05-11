package mderr

import (
	"errors"
	"fmt"
)

// MDErr the underlying error struct
type MDErr struct {
	message  string
	cause    error
	metadata map[string]any
}

// New creates a new error with the given metadata
func New(msg string, md map[string]any) error {
	return Wrap(nil, msg, md)
}

// Wrap wraps an error with a new error and metadata
func Wrap(err error, msg string, md map[string]any) error {
	if md == nil {
		md = map[string]any{}
	}

	return &MDErr{
		message:  msg,
		cause:    err,
		metadata: md,
	}
}

// Message gets the error message
func (e *MDErr) Message() string {
	return e.message
}

// Cause gets the cause (wrapped error) of the error
func (e *MDErr) Cause() error {
	return e.cause
}

// Metadata gets a map of contextual metadata associated with the error
func (e *MDErr) Metadata() map[string]any {
	return e.metadata
}

// Error is an alias of Error
func (e *MDErr) Error() string {
	return Error(e)
}

// Unwrap is an alias of Cause
func (e *MDErr) Unwrap() error {
	return e.Cause()
}

// Root is an alias of Root
func (e *MDErr) Root() error {
	return Root(e)
}

// Depth is an alias of Depth
func (e *MDErr) Depth() int {
	return Depth(e)
}

// Array is an alias of Array
func (e *MDErr) Array() []error {
	return Array(e)
}

// RArray is an alias of RArray
func (e *MDErr) RArray() []error {
	return RArray(e)
}

// Nest is an alias of Nest
func (e *MDErr) Nest() map[string]any {
	return Nest(e)
}

// Stack is an alias of Stack
func (e *MDErr) Stack() []*struct {
	Message  string         `json:"message"`
	Metadata map[string]any `json:"metadata"`
} {
	return Stack(e)
}

// RStack is an alias of RStack
func (e *MDErr) RStack() []*struct {
	Message  string         `json:"message"`
	Metadata map[string]any `json:"metadata"`
} {
	return RStack(e)
}

// AsIs checks if the error is an MDErr and converts it if so
func AsIs(err error) (*MDErr, bool) {
	if err == nil {
		return nil, false
	}

	mde, ok := err.(*MDErr)

	return mde, ok
}

// Message gets the error message for the error
func Message(err error) string {
	as, is := AsIs(err)

	if is {
		return as.Message()
	}

	return err.Error()
}

// Cause gets the cause of the error
// similar to unwrapping an error
func Cause(err error) error {
	as, is := AsIs(err)

	if is {
		return as.Cause()
	}

	return errors.Unwrap(err)
}

// Metadata gets the metadata for the error
// if the error is not an MDErr, nil is returned
func Metadata(err error) map[string]any {
	as, is := AsIs(err)

	if !is {
		return nil
	}

	return as.Metadata()
}

// Error gets the full unwrapped error message
func Error(err error) string {
	if err == nil {
		return ""
	}

	as, is := AsIs(err)

	if !is {
		return err.Error()
	}

	buf := as.Message()
	unw := as.Cause()

	if unw != nil {
		buf = fmt.Sprintf("%s: %s", buf, Error(unw))
	}

	return buf
}

// Root gets the root cause of an error
// this is the first error that caused
func Root(err error) error {
	if err == nil {
		return nil
	}

	r := Root(errors.Unwrap(err))

	if r == nil {
		return err
	}

	return r
}

// Depth gets the depth of the error stack
func Depth(err error) int {
	if err == nil {
		return 0
	}

	return 1 + Depth(errors.Unwrap(err))
}

// Array gets the error stack
// for contextual logging
func Array(err error) []error {
	arr := make([]error, 0)

	if err == nil {
		return arr
	}

	return append(append(arr, err), Array(errors.Unwrap(err))...)
}

// RArray gets the reverse of Array
func RArray(err error) []error {
	arr := make([]error, 0)

	if err == nil {
		return arr
	}

	return append(RArray(errors.Unwrap(err)), append(arr, err)...)
}

// Nest gets a nested map of the error
// useful for contextual logging
func Nest(err error) map[string]any {
	if err == nil {
		return nil
	}

	return map[string]any{
		"message":  Message(err),
		"cause":    Nest(errors.Unwrap(err)),
		"metadata": Metadata(err),
	}
}

// Stack gets a flattened list of the error stack
// useful for contextual logging
func Stack(err error) []*struct {
	Message  string         `json:"message"`
	Metadata map[string]any `json:"metadata"`
} {
	sta := Array(err)
	arr := make([]*struct {
		Message  string         `json:"message"`
		Metadata map[string]any `json:"metadata"`
	}, len(sta))

	for ind, err := range sta {
		arr[ind] = &struct {
			Message  string         `json:"message"`
			Metadata map[string]any `json:"metadata"`
		}{
			Message:  Message(err),
			Metadata: Metadata(err),
		}
	}

	return arr
}

// RStack gets a flattened list of the reversed error stack
// useful for contextual logging, depending on how you wanna see it
func RStack(err error) []*struct {
	Message  string         `json:"message"`
	Metadata map[string]any `json:"metadata"`
} {
	sta := RArray(err)
	arr := make([]*struct {
		Message  string         `json:"message"`
		Metadata map[string]any `json:"metadata"`
	}, len(sta))

	for ind, err := range sta {
		arr[ind] = &struct {
			Message  string         `json:"message"`
			Metadata map[string]any `json:"metadata"`
		}{
			Message:  Message(err),
			Metadata: Metadata(err),
		}
	}

	return arr
}
