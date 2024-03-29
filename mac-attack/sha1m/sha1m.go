package sha1m
// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package sha1 implements the SHA1 hash algorithm as defined in RFC 3174.

import (
	"github.com/mdko/cs465/mac-attack/cryptom"
	"github.com/mdko/cs465/mac-attack/hashm"
)

func init() {
	cryptom.RegisterHash(cryptom.SHA1, New)
}

// The size of a SHA1 checksum in bytes.
const Size = 20

// The blocksize of SHA1 in bytes.
const BlockSize = 64

const (
	chunk = 64			//512 bits in each chunk
	init0 = 0x67452301
	init1 = 0xEFCDAB89
	init2 = 0x98BADCFE
	init3 = 0x10325476
	init4 = 0xC3D2E1F0
)

// digest represents the partial evaluation of a checksum.
type digest struct {
	h   [5]uint32
	x   [chunk]byte
	nx  int
	len uint64
}

func (d *digest) Reset() {
	d.h[0] = init0
	d.h[1] = init1
	d.h[2] = init2
	d.h[3] = init3
	d.h[4] = init4
	d.nx = 0
	d.len = 0
}

func (d *digest) OverrideRegisters(vals [5]uint32) {
	d.h[0] = vals[0]
	d.h[1] = vals[1]
	d.h[2] = vals[2]
	d.h[3] = vals[3]
	d.h[4] = vals[4]
	d.nx = 0
	d.len = 0
}

func (d *digest) ChangeLength(length uint64) {
	d.len = length
}

// New returns a new hash.Hash computing the SHA1 checksum.
func New() hashm.Hashm {
	d := new(digest)
	d.Reset()
	return d
}

func (d *digest) Size() int { return Size }

func (d *digest) BlockSize() int { return BlockSize }

func (d *digest) Write(p []byte) (nn int, err error) {
	nn = len(p)
	d.len += uint64(nn)
	if d.nx > 0 {
		n := len(p)
		if n > chunk-d.nx {
			n = chunk - d.nx
		}
		for i := 0; i < n; i++ {
			d.x[d.nx+i] = p[i]
		}
		d.nx += n
		if d.nx == chunk {
			Block(d, d.x[0:])
			d.nx = 0
		}
		p = p[n:]
	}
	if len(p) >= chunk {
		n := len(p) &^ (chunk - 1)
		Block(d, p[:n])
		p = p[n:]
	}
	if len(p) > 0 {
		d.nx = copy(d.x[:], p)
	}
	return
}

func (d0 *digest) Sum(in []byte) []byte {
	// Make a copy of d0 so that caller can keep writing and summing.
	d := *d0

	// Padding.  Add a 1 bit and 0 bits until 56 bytes mod 64.
	len := d.len
	var tmp [64]byte
	tmp[0] = 0x80
	if len%64 < 56 {
		d.Write(tmp[0 : 56-len%64])
	} else {
		d.Write(tmp[0 : 64+56-len%64])
	}

	// Length in bits.
	len <<= 3
	for i := uint(0); i < 8; i++ {
		tmp[i] = byte(len >> (56 - 8*i))
	}
	d.Write(tmp[0:8])

	if d.nx != 0 {
		panic("d.nx != 0")
	}

	var digest [Size]byte
	for i, s := range d.h {
		digest[i*4] = byte(s >> 24)
		digest[i*4+1] = byte(s >> 16)
		digest[i*4+2] = byte(s >> 8)
		digest[i*4+3] = byte(s)
	}

	return append(in, digest[:]...)
}