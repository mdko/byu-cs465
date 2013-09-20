package main

import (
	"crypto/sha1"
	"math/rand"
	"encoding/binary"
	"fmt"
)

func main() {
	var nbits int = 20
	// input
	// 2^n unique input strings
	// length of n + 1, n = number of bits in hash
	// so if n = 20, 2^20 = 1,048,576 values of length 21 bits

	// output
	// truncate to the number of bits that lets me get attack in a few seconds

	//------------------------------------------------------------------------
	//------------------------------------------------------------------------
	// Collision attack:
	// Find two separate input messages that both hash to same digest
	// Expected Time: 2^(n/2)
	//------------------------------------------------------------------------
	generatedMessages := make(map[uint32]uint32, 2^nbits) // Hold 32 bit integers, but the last 11 bits will be zeroed
	var numTries int = 0
	for {	
		numTries++
		
		// Generate new message m1 that is n + 1 bits long
		m1 := rand.Uint32() & 0x001FFFFF // nbits + 1 (21) bits long, bits 22-32 zeroed out
		buf := make([]byte, 4)
		fmt.Printf("Message m1: %x\n", m1)
		
		// Put it in a byte array for use in the hash function (cast to uint64 because that's what function accepts)
		binary.PutUvarint(buf, uint64(m1))
		
		// Generate hash
		h := sha1.New()
		h.Write(buf)
		hashArr := h.Sum(nil) 						// a []byte
		hash64, _ := binary.Uvarint(hashArr) 		// a uint64
		hash := uint32(hash64) & 0x0000FFFFF		// nbits (20) long
		fmt.Printf("Hash: %x\n\n", hash)
		//h.Reset()

		// Check to see if we've already discovered this hash by searching the map
		// If there is a match in the map, m2 = value associated with that hash in map
		// Else, add h to the map, with h being the key and m1 being the value
		m2, present := generatedMessages[hash]
		if present {
			if (m2 != m1) {
				fmt.Println("Found two messages that hash to same digest but aren't the same")
				fmt.Printf("%x\n%x\n", m1, m2)
				fmt.Printf("in %v tries\n", numTries)
				break
			}
		} else {
			generatedMessages[hash] = m1
		}
	}
	//------------------------------------------------------------------------
	//------------------------------------------------------------------------
	// Preimage attack:
	// Given a hash value, find the source message that hashes to that value
	// Expected Time: 2^n
	//------------------------------------------------------------------------
}