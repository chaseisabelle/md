package md

import "github.com/chaseisabelle/md/mderr"

// MD is a shorthand way to create metadata map[string]any
// instead of calling md.E(nil, "this is an error", map[string]any{"foo":"bar"})
// you can call md.E(nil, "this is an error", md.MD{"foo":"bar"})
type MD map[string]any

// E is a shorthand way to create an error with metadata
func E(msg string, md MD) error {
	return mderr.Wrap(nil, msg, md)
}

// W is a shorthand way to wrap and error with metadata
func W(err error, msg string, md MD) error {
	return mderr.Wrap(err, msg, md)
}
