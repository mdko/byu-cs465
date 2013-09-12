package keyexpansion

import (
		"testing"
		"github.com/mdko/cs465/aes/constants"
)

func TestRotWord(t *testing.T) {
		inoutmap := map[[4]byte][4]byte {
				[4]byte { 0x09, 0xcf, 0x4f, 0x3c, } :
				[4]byte { 0xcf, 0x4f, 0x3c, 0x09, } ,

				[4]byte { 0x2a, 0x6c, 0x76, 0x05, } :
				[4]byte { 0x6c, 0x76, 0x05, 0x2a, } ,

				[4]byte { 0x73, 0x59, 0xf6, 0x7f, } :
				[4]byte { 0x59, 0xf6, 0x7f, 0x73, } ,

				[4]byte { 0x6d, 0x7a, 0x88, 0x3b, } :
				[4]byte { 0x7a, 0x88, 0x3b, 0x6d, } ,
		}

		for in, exp := range(inoutmap) {
				if ans := RotWord(in); ans != exp {
						t.Errorf("RotWord(%x) = %x, want %x", in, ans, exp)

				}
		}
}

func TestSubWord(t *testing.T) {
		inoutmap := map[[4]byte][4]byte {
				[4]byte { 0xcf, 0x4f, 0x3c, 0x09, } :
				[4]byte { 0x8a, 0x84, 0xeb, 0x01, } ,

				[4]byte { 0x6c, 0x76, 0x05, 0x2a, } :
				[4]byte { 0x50, 0x38, 0x6b, 0xe5, } ,

				[4]byte { 0x59, 0xf6, 0x7f, 0x73, } :
				[4]byte { 0xcb, 0x42, 0xd2, 0x8f, } ,

				[4]byte { 0x7a, 0x88, 0x3b, 0x6d, } :
				[4]byte { 0xda, 0xc4, 0xe2, 0x3c, } ,
		}

		for in, exp := range(inoutmap) {
				if ans := SubWord(in); ans != exp {
						t.Errorf("SubWord(%x) = %x, want %x", in, ans, exp)

				}
		}
}

func TestMakeByteArray(t *testing.T) {
		inoutmap := map[uint32][4]byte {
				0xcf4f3c09 : [4]byte { 0xcf, 0x4f, 0x3c, 0x09, } ,
				0x8a84eb01 : [4]byte { 0x8a, 0x84, 0xeb, 0x01, } ,
				0x6c76052a : [4]byte { 0x6c, 0x76, 0x05, 0x2a, } ,
		}

		for in, out := range(inoutmap) {
				if ans := MakeByteArray(in); ans != out {
						t.Errorf("MakeByteArray(%x) = %x, want %x", in, ans, out)
				}
		}
}

func TestXORWords(t *testing.T) {
		wordArr0a := [4]byte {
				0x8a, 0x84, 0xeb, 0x01,
		}
		wordArr0b := [4]byte {
				0x01, 0x00, 0x00, 0x00,
		}
		wordArr0c := [4]byte {
				0x8b, 0x84, 0xeb, 0x01,
		}
		if ans0 := XORWords(wordArr0a, wordArr0b); ans0 != wordArr0c {
				t.Errorf("XORWords(%v, %v) = %v, want %v", wordArr0a, wordArr0b, ans0, wordArr0c)
		}

		wordArr1a := [4]byte {
				0xcb, 0x42, 0xd2, 0x8f,
		}
		wordArr1b := [4]byte {
				0x04, 0x00, 0x00, 0x00,
		}
		wordArr1c := [4]byte {
				0xcf, 0x42, 0xd2, 0x8f,
		}
		if ans1 := XORWords(wordArr1a, wordArr1b); ans1 != wordArr1c {
				t.Errorf("XORWords(%v, %v) = %v, want %v", wordArr1a, wordArr1b, ans1, wordArr1c)
		}


}

func TextKeyExpansion(t *testing.T) {
		// key = 0x2b 7e 15 16 28 ae d2 a6 ab f7 15 88 09 cf 4f 3c
		cipherKey0 := [4][4]byte {
				{ 0x2b, 0x7e, 0x15, 0x16, } ,
				{ 0x28, 0xae, 0xd2, 0xa6, } ,
				{ 0xab, 0xf7, 0x15, 0x88, } ,
				{ 0x09, 0xcf, 0x4f, 0x3c, } ,
		}

		var expected [constants.Nb * (constants.Nr + 1)][4]byte

		if result := KeyExpansion(cipherKey0); result == expected {
			t.Errorf("Key Expansion result: %v", result)
		}
//		KeyExpansion([Nk][4])
}
