package keyexpansion

import (
		"testing"
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

func TextKeyExpansion(t * testing.T) {

}
