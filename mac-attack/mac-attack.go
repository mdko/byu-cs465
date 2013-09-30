package main

import (
	"github.com/mdko/cs465/mac-attack/sha1m"
	"fmt"
)

var (
	debug bool = false

	// "No one has completed lab 2 so give them all a 0"
	originalMessage = []byte {
		0x4e, 0x6f, 0x20, 0x6f, 0x6e, 0x65, 0x20, 0x68, 0x61, 0x73, 0x20, 0x63, 0x6f, 0x6d, 0x70, 0x6c,
		0x65, 0x74, 0x65, 0x64, 0x20, 0x6c, 0x61, 0x62, 0x20, 0x32, 0x20, 0x73, 0x6f, 0x20, 0x67, 0x69,
		0x76, 0x65, 0x20, 0x74, 0x68, 0x65, 0x6d, 0x20, 0x61, 0x6c, 0x6c, 0x20, 0x61, 0x20, 0x30,
	}

	// "f4b645e89faaec2ff8e443c595009c16dbdfba4b"
	digestOfOriginalMessage = [5]uint32 {
		0xf4b645e8,
		0x9faaec2f,
		0xf8e443c5,
		0x95009c16,
		0xdbdfba4b,
	}

	lengthOfKey int = 128

	extensionToAdd string = "P. S. Except for Mike, go ahead and give him the full 100 points"
)

func main() {
	extensionToAddBytes := []byte(extensionToAdd)
	//fmt.Printf("%x\n", extensionToAddBytes)
	//fmt.Printf("%d\n", len(originalMessage))
	//fmt.Printf("%d\n", len(extensionToAddBytes))

	h := sha1m.New()
	h.OverrideRegisters(digestOfOriginalMessage)
	h.Write(extensionToAddBytes)
	extensionHash := h.Sum(nil)
	fmt.Printf("%x\n", extensionHash)
	

	// Padding and length.  Add a 1 bit and 0 bits until 56 bytes mod 64 (8 bytes at end reserved for length)
	var paddedOriginalMessage [64]byte
	len := len(originalMessage)
	paddedOriginalMessage[0:len] = originalMessage[:]
	paddedOriginalMessage[len] = 0x80
	for i := len; i < 56; i++ {
		paddedOriginalMessage[i] = 0x00
	}
	len <<= 3
	for i := uint(0); i < 8; i++ {
		paddedOriginalMessage[56 + i] = byte(len >> (56 - 8*i))
	}

	fmt.Printf("%x\n", paddedOriginalMessage)

	// m1 := rand.Uint32() & getBitMask(bitN + 1)						// nbits + 1 (21) bits long, bits 22-32 zeroed out
	// buf := make([]byte, 4)
	// binary.BigEndian.PutUint32(buf, m1)
	// h.Reset()														// unneeded since I use a new hash object in each loop
}