package main

import (
	"crypto/sha1"
	"math/rand"
	"encoding/binary"
	"time"
	"fmt"
)

var debug bool = false

// Input strings of length of n + 1, n = number of bits in hash
var collisionNbits int = 30
var preimageNbits int = 18

var numTrials int = 100


func debugPrint(message string) {
	if (debug) {
		fmt.Printf("DEBUG: %s\n", message)
	}
}

func getBitMask(numHighBits int) (mask uint32) {
	mask = 1
	mask = (mask << uint16(numHighBits)) - 1
	debugPrint(fmt.Sprintf("Bit mask for nbit %v: %b\n", numHighBits, mask))
	return
}

func main() {
	runCollisionAttack(numTrials)
	run2ndPreimageAttack(numTrials)
}

func runCollisionAttack(numTrials int) {
	//-------------------------------------------------------------------------------------------------------
	//-------------------------------------------------------------------------------------------------------
	// Collision attack:
	// Find two separate input messages that both hash to same digest
	// Expected Time: 2^(n/2)
	// Eg. for n = 20, 2^10 = 1024 ~ approximately number of tries needed
	//-------------------------------------------------------------------------------------------------------
	fmt.Println("Running Collision Attack")
	rand.Seed(time.Now().UTC().UnixNano())

	generatedMessages := make(map[uint32]uint32, 2^collisionNbits) 	// Hold 32 bit integers, but the last 11 bits will be zeroed
	var numTries int = 0
	for {	
		numTries++
		
		// Generate new message m1 that is n + 1 bits long
		m1 := rand.Uint32() & getBitMask(collisionNbits + 1)			// nbits + 1 (21) bits long, bits 22-32 zeroed out
		buf := make([]byte, 4)
		debugPrint(fmt.Sprintf("Message m1: %x\n", m1))
		
		// Put it in a byte array for use in the hash function (cast to uint64 because that's what function accepts)
		binary.LittleEndian.PutUint32(buf, m1)
		debugPrint(fmt.Sprintf("Buf: %x\n", buf))
		
		// Generate hash
		h := sha1.New()
		h.Write(buf)
		hashArr := h.Sum(nil) 											// a []byte
		hash32 := binary.LittleEndian.Uint32(hashArr) 					// a uint32
		hash := hash32 & getBitMask(collisionNbits)					// nbits (20) long
		debugPrint(fmt.Sprintf("Hash: %x\n\n", hash))
		//h.Reset()														// unneeded since I use a new hash object in each loop

		// Check to see if we've already discovered this hash by searching the map
		// If there is a match in the map, m2 = value associated with that hash in map
		// Else, add h to the map, with h being the key and m1 being the value
		m2, present := generatedMessages[hash]
		if present {
			if (m2 != m1) {
				debugPrint(fmt.Sprint("Found two messages that hash to same digest but aren't the same"))
				debugPrint(fmt.Sprintf("%x\n%x\n", m1, m2))
				break
			}
		} else {
			generatedMessages[hash] = m1
		}
	}
	fmt.Printf("Collision Attack: For a message of %v bits and hash of %v bits, took %v tries\n", 
		collisionNbits+1, collisionNbits, numTries)
}

func run2ndPreimageAttack(numTrials int) {
	//-------------------------------------------------------------------------------------------------------
	//-------------------------------------------------------------------------------------------------------
	// 2nd Preimage attack:
	// Given a message m1, find another message m2 s. that m1 != m2 and hash(m1) = hash(m2)
	// Expected Time: 2^n
	// Eg. for n = 18, 2^n = 262,144 tries
	//-------------------------------------------------------------------------------------------------------
	fmt.Println("Running 2nd Pre-image Attack")
	rand.Seed(time.Now().UTC().UnixNano())

	// Generate new message m1 that is n + 1 bits long
	m1 := rand.Uint32() & getBitMask(preimageNbits + 1)
	buf := make([]byte, 4)
	debugPrint(fmt.Sprintf("Message m1: %x\n", m1))
	
	// Put it in a byte array for use in the hash function (cast to uint64 because that's what function accepts)
	binary.LittleEndian.PutUint32(buf, m1)
	debugPrint(fmt.Sprintf("Buf: %x\n", buf))
	
	// Generate hash
	h := sha1.New()
	h.Write(buf)
	hashArr := h.Sum(nil) 									// a []byte
	hash32 := binary.LittleEndian.Uint32(hashArr) 			// a uint32
	m1hash := hash32 & getBitMask(preimageNbits)			// nbits (20) long
	h.Reset()
	debugPrint(fmt.Sprintf("Hash: %x\n\n", m1hash))

	var numTries int = 0
	for {	
		numTries++
		
		// Generate new message m2 that is n + 1 bits long
		m2 := rand.Uint32() & getBitMask(preimageNbits + 1)
		buf = make([]byte, 4)
		debugPrint(fmt.Sprintf("Message m2: %x\n", m2))
		
		// Put it in a byte array for use in the hash function (cast to uint64 because that's what function accepts)
		binary.LittleEndian.PutUint32(buf, m2)
		debugPrint(fmt.Sprintf("Buf: %x\n", buf))
		
		// Generate hash
		h = sha1.New()
		h.Write(buf)
		hashArr = h.Sum(nil) 								// a []byte
		hash32 = binary.LittleEndian.Uint32(hashArr) 		// a uint32
		m2hash := hash32 & getBitMask(preimageNbits)		// nbits (20) long
		debugPrint(fmt.Sprintf("Hash: %x\n\n", m2hash))
		//h.Reset()											// unneeded since I use a new hash object in each loop

		if (m2hash == m1hash) && (m1 != m2) {
			debugPrint(fmt.Sprint("Found two messages that hash to same digest but aren't the same"))
			debugPrint(fmt.Sprintf("%x\n%x\n", m1, m2))
			break
		}
	}
	fmt.Printf("Pre-image attack: For a message of %v bits and hash of %v bits, took %v tries\n", 
		preimageNbits+1, preimageNbits, numTries)
}