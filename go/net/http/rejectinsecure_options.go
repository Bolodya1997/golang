// Copyright 2020 The searKing Author. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Code generated by "go-option -type rejectInsecure"; DO NOT EDIT.

package http

var _default_RejectInsecure_value = func() (val rejectInsecure) { return }()

// A RejectInsecureOption sets options.
type RejectInsecureOption interface {
	apply(*rejectInsecure)
}

// EmptyRejectInsecureOption does not alter the configuration. It can be embedded
// in another structure to build custom options.
//
// This API is EXPERIMENTAL.
type EmptyRejectInsecureOption struct{}

func (EmptyRejectInsecureOption) apply(*rejectInsecure) {}

// RejectInsecureOptionFunc wraps a function that modifies rejectInsecure into an
// implementation of the RejectInsecureOption interface.
type RejectInsecureOptionFunc func(*rejectInsecure)

func (f RejectInsecureOptionFunc) apply(do *rejectInsecure) {
	f(do)
}

// sample code for option, default for nothing to change
func _RejectInsecureOptionWithDefault() RejectInsecureOption {
	return RejectInsecureOptionFunc(func(*rejectInsecure) {
		// TODO nothing to change
	})
}

func (o *rejectInsecure) ApplyOptions(options ...RejectInsecureOption) *rejectInsecure {
	for _, opt := range options {
		if opt == nil {
			continue
		}
		opt.apply(o)
	}
	return o
}