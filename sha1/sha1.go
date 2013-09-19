package main

import (
	"crypto/sha1"
	"io"
	"fmt"
)

func main() {
	// input
	// 2^n unique input strings
	// length of n + 1, n = number of bits in hash

	// output
	// truncate to the number of bits that lets me get attack in a few seconds


	h := sha1.New()
	io.WriteString(h, "His money is twice tainted: 'taint yours and 'taint mine.")
    fmt.Printf("% x", h.Sum(nil))
}