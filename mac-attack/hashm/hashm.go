package hashm
// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package hash provides interfaces for hash functions.

import "io"

// Hash is the common interface implemented by all hash functions.
type Hashm interface {
	// Write adds more data to the running hash.
	// It never returns an error.
	io.Writer

	// Sum appends the current hash to b and returns the resulting slice.
	// It does not change the underlying hash state.
	Sum(b []byte) []byte

	// Reset resets the hash to one with zero bytes written.
	Reset()

	// Overrides registers (I added this for mac-attack)
	OverrideRegisters(vals [5]uint32)

	// Change length appended after padding (I added this)
	ChangeLength(length uint64)

	// Size returns the number of bytes Sum will return.
	Size() int

	// BlockSize returns the hash's underlying block size.
	// The Write method must be able to accept any amount
	// of data, but it may operate more efficiently if all writes
	// are a multiple of the block size.
	BlockSize() int
}

// Hash32 is the common interface implemented by all 32-bit hash functions.
type Hash32m interface {
	Hashm
	Sum32() uint32
}

// Hash64 is the common interface implemented by all 64-bit hash functions.
type Hash64m interface {
	Hashm
	Sum64() uint64
}