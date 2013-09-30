package main

import (
	"github.com/mdko/cs465/mac-attack/sha1m"
	"fmt"
)

var (
	debug bool = true
	extensionToAdd string = "P. S. Except for Mike, go ahead and give him the full 100 points" // 64 bytes long
	bytesForMessageLengthField int = 8

	//-------------------------------------------------------------------------------------------------
	// We intercept/find out the following:
	//-------------------------------------------------------------------------------------------------
	originalMessage = []byte {				// "No one has completed lab 2 so give them all a 0"
		0x4e, 0x6f, 0x20, 0x6f, 0x6e, 0x65, 0x20, 0x68, 0x61, 0x73, 0x20, 0x63, 0x6f, 0x6d, 0x70, 0x6c,
		0x65, 0x74, 0x65, 0x64, 0x20, 0x6c, 0x61, 0x62, 0x20, 0x32, 0x20, 0x73, 0x6f, 0x20, 0x67, 0x69,
		0x76, 0x65, 0x20, 0x74, 0x68, 0x65, 0x6d, 0x20, 0x61, 0x6c, 0x6c, 0x20, 0x61, 0x20, 0x30,
	}
	digestOfOriginalMessage = [5]uint32 {	// "f4b645e89faaec2ff8e443c595009c16dbdfba4b"
		0xf4b645e8,
		0x9faaec2f,
		0xf8e443c5,
		0x95009c16,
		0xdbdfba4b,
	}
	byteLengthOfKey int = 16 // 128 bits
)

func debugPrint(message string) {
	if (debug) {
		fmt.Printf("DEBUG: %s\n", message)
	}
}

func main() {	
	h := sha1m.New()

	var numBytesPadding int
	lenOriginalMessage := len(originalMessage)
	lenWholeMessage := byteLengthOfKey + lenOriginalMessage	
	roomForMessageLengthField := h.BlockSize() /*64*/ - (lenWholeMessage % h.BlockSize())
	debugPrint(fmt.Sprintf("Length of original message: %d", len(originalMessage)))
	debugPrint(fmt.Sprintf("Length of key: %d", byteLengthOfKey))
	debugPrint(fmt.Sprintf("Len of whole message (key + message): %d", lenWholeMessage))
	debugPrint(fmt.Sprintf("Room for message length field: %d", roomForMessageLengthField))

	// If whole message (key + message) is 63 bytes, then 64 - 63 = 1 < 8, so there 
	// needs to be an extra chunk for padding since not enough room in 1st chunk for 
	// storing length of message (an 8 byte field)
	if roomForMessageLengthField < bytesForMessageLengthField /*8*/{
		numBytesPadding = roomForMessageLengthField /* room for padding in first chunk */ + (h.BlockSize() - bytesForMessageLengthField) /* padding in additional chunk*/
	} else {				
		// If whole message (key + message) is 40 bytes, then 64 - 40 = 24 >=8, so there		
		// is enough room for storing message length in the dedicated field within this chunk,
		// plus there will be padding...
		numBytesPadding = roomForMessageLengthField - bytesForMessageLengthField 
	}								
	debugPrint(fmt.Sprintf("Bytes of padding: %d", numBytesPadding))

	extensionToAddBytes := []byte(extensionToAdd)
	lenExtensionToAddBytes := len(extensionToAddBytes)
	debugPrint(fmt.Sprintf("Extension to add: %x", extensionToAddBytes))
	debugPrint(fmt.Sprintf("Length of extension to add: %d", lenExtensionToAddBytes))

	lenWholeMessageToReturn := lenOriginalMessage + numBytesPadding + bytesForMessageLengthField + lenExtensionToAddBytes
	lenBobWillSee := lenWholeMessageToReturn + byteLengthOfKey
	lengthField := [8]byte(uint64(lenBobWillSee * 8))
	debugPrint(fmt.Sprintf("Length of whole message to return: %d", lenWholeMessageToReturn))
	debugPrint(fmt.Sprintf("Length Bob will see(bits): %x", lengthField))
	
	h.OverrideRegisters(digestOfOriginalMessage)
	h.Write(extensionToAddBytes)
	h.ChangeLength(uint64(lenBobWillSee))
	extensionHash := h.Sum(nil)
	fmt.Printf("Hash of extension(final hash to give to Bob): %x\n", extensionHash)

	buf := make([]byte, lenWholeMessageToReturn)
	buf =  append(buf, originalMessage...)
	buf = append(buf, 0x80)
	for i := 1; i < numBytesPadding; i++ {
		buf = append(buf, 0x00)
	}
	buf = append(buf, byte(lenWholeMessage * 8))
	buf = append(buf, extensionToAddBytes...)
	fmt.Printf("Modified message to send to Bob: %x\n", buf)


	// lenMessageToReturn := len(originalMessage) +
	// var messageToReturn []byte

	// Padding and length.  Add a 1 bit and 0 bits until 56 bytes mod 64 (8 bytes at end reserved for length)
	//var paddedOriginalMessage [64]byte
	// len := len(originalMessage)
	// paddedOriginalMessage[0:len] = originalMessage[:]
	// paddedOriginalMessage[len] = 0x80
	// for i := len; i < 56; i++ {
	// 	paddedOriginalMessage[i] = 0x00
	// }
	// len <<= 3
	// for i := uint(0); i < 8; i++ {
	// 	paddedOriginalMessage[56 + i] = byte(len >> (56 - 8*i))
	// }

	//fmt.Printf("%x\n", paddedOriginalMessage)

	// m1 := rand.Uint32() & getBitMask(bitN + 1)						// nbits + 1 (21) bits long, bits 22-32 zeroed out
	// buf := make([]byte, 4)
	// binary.BigEndian.PutUint32(buf, m1)
	// h.Reset()														// unneeded since I use a new hash object in each loop
}