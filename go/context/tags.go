package context

import (
	"context"
)

var (
	// NopTags is a trivial, minimum overhead implementation of Tags for which all operations are no-ops.
	NopTags = &nopTags{}
)

// Tags is the interface used for storing request tags between Context calls.
// The default implementation is *not* thread safe, and should be handled only in the context of the request.
type Tags interface {
	// Set sets the given key in the metadata tags.
	Set(key string, value interface{}) Tags
	// Has checks if the given key exists.
	Has(key string) bool
	// Values returns a map of key to values.
	// Do not modify the underlying map, please use Set instead.
	Values() map[string]interface{}
}

// ExtractTags returns a pre-existing Tags object in the Context.
// If the context wasn't set in a tag interceptor, a no-op Tag storage is returned that will *not* be propagated in context.
func ExtractTags(ctx context.Context, key interface{}) Tags {
	t, ok := ctx.Value(key).(Tags)
	if !ok {
		return NopTags
	}

	return t
}

func WithTags(ctx context.Context, key interface{}, tags Tags) context.Context {
	return context.WithValue(ctx, key, tags)
}

func NewTags() Tags {
	return &mapTags{values: make(map[string]interface{})}
}