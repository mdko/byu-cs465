package cipher

import (
		"testing"
		"github.com/mdko/cs465/aes/constants"
)

func TestAddRoundKey(t *testing.T) {
	state0 := [4][4]byte {
		{ 0x32, 0x43, 0xf6, 0xa8, } ,
		{ 0x88, 0x5a, 0x30, 0x8d, } ,
		{ 0x31, 0x31, 0x98, 0xa2, } ,
		{ 0xe0, 0x37, 0x07, 0x34, } ,
	}
	roundKey0 := [4][4]byte {
		{ 0x2b, 0x7e, 0x15, 0x16, } ,
		{ 0x28, 0xae, 0xd2, 0xa6, } ,
		{ 0xab, 0xf7, 0x15, 0x88, } ,
		{ 0x09, 0xcf, 0x4f, 0x3c, } ,
	}
	expectedState0 := [4][4]byte {
		{ 0x19, 0x3d, 0xe3, 0xbe, } ,
		{ 0xa0, 0xf4, 0xe2, 0x2b, } ,
		{ 0x9a, 0xc6, 0x8d, 0x2a, } ,
		{ 0xe9, 0xf8, 0x48, 0x08, } ,
	}
	if newState := AddRoundKey(state0, roundKey0); newState != expectedState0 {
		t.Errorf("Test 0:\nExpected: %v\nGiven: %v\n", expectedState0, newState)
	}

	state1 := [4][4]byte {
		{ 0x04, 0x66, 0x81, 0xe5, } ,
		{ 0xe0, 0xcb, 0x19, 0x9a, } ,
		{ 0x48, 0xf8, 0xd3, 0x7a, } ,
		{ 0x28, 0x06, 0x26, 0x4c, } ,
	}
	roundKey1 := [4][4]byte {
		{ 0xa0, 0xfa, 0xfe, 0x17, } ,
		{ 0x88, 0x54, 0x2c, 0xb1, } ,
		{ 0x23, 0xa3, 0x39, 0x39, } ,
		{ 0x2a, 0x6c, 0x76, 0x05, } ,
	}
	expectedState1 := [4][4]byte {
		{ 0xa4, 0x9c, 0x7f, 0xf2, } ,
		{ 0x68, 0x9f, 0x35, 0x2b, } ,
		{ 0x6b, 0x5b, 0xea, 0x43, } ,
		{ 0x02, 0x6a, 0x50, 0x49, } ,
	}
	if newState := AddRoundKey(state1, roundKey1); newState != expectedState1 {
		t.Errorf("Test 1:\nExpected: %v\nGiven: %v\n", expectedState1, newState)
	}
}

func TestSubBytes(t *testing.T) {
	state0 := [4][4]byte {
		{ 0x19, 0x3d, 0xe3, 0xbe, } ,
		{ 0xa0, 0xf4, 0xe2, 0x2b, } ,
		{ 0x9a, 0xc6, 0x8d, 0x2a, } ,
		{ 0xe9, 0xf8, 0x48, 0x08, } ,
	}
	expectedState0 := [4][4]byte {
		{ 0xd4, 0x27, 0x11, 0xae, } ,
		{ 0xe0, 0xbf, 0x98, 0xf1, } ,
		{ 0xb8, 0xb4, 0x5d, 0xe5, } ,
		{ 0x1e, 0x41, 0x52, 0x30, } ,
	}
	if newState := SubBytes(state0); newState != expectedState0 {
		t.Errorf("Test 0:\nExpected: %v\nGiven: %v\n", expectedState0, newState)
	}

	state1 := [4][4]byte {
		{ 0xa4, 0x9c, 0x7f, 0xf2, } ,
		{ 0x68, 0x9f, 0x35, 0x2b, } ,
		{ 0x6b, 0x5b, 0xea, 0x43, } ,
		{ 0x02, 0x6a, 0x50, 0x49, } ,
	}
	expectedState1 := [4][4]byte {
		{ 0x49, 0xde, 0xd2, 0x89, } ,
		{ 0x45, 0xdb, 0x96, 0xf1, } ,
		{ 0x7f, 0x39, 0x87, 0x1a, } ,
		{ 0x77, 0x02, 0x53, 0x3b, } ,
	}
	if newState := SubBytes(state1); newState != expectedState1 {
		t.Errorf("Test 1:\nExpected: %v\nGiven: %v\n", expectedState1, newState)
	}

}

func TestShiftRows(t *testing.T) {
	state0 := [4][4]byte {
		{ 0xd4, 0x27, 0x11, 0xae, } ,
		{ 0xe0, 0xbf, 0x98, 0xf1, } ,
		{ 0xb8, 0xb4, 0x5d, 0xe5, } ,
		{ 0x1e, 0x41, 0x52, 0x30, } ,
	}
	expectedState0 := [4][4]byte {
		{ 0xd4, 0xbf, 0x5d, 0x30, } ,
		{ 0xe0, 0xb4, 0x52, 0xae, } ,
		{ 0xb8, 0x41, 0x11, 0xf1, } ,
		{ 0x1e, 0x27, 0x98, 0xe5, } ,
	}
	if newState := ShiftRows(state0); newState != expectedState0 {
		t.Errorf("Test 0:\nExpected: %v\nGiven: %v\n", expectedState0, newState)
	}

	state1 := [4][4]byte {
		{ 0x49, 0xde, 0xd2, 0x89, } ,
		{ 0x45, 0xdb, 0x96, 0xf1, } ,
		{ 0x7f, 0x39, 0x87, 0x1a, } ,
		{ 0x77, 0x02, 0x53, 0x3b, } ,
	}
	expectedState1 := [4][4]byte {
		{ 0x49, 0xdb, 0x87, 0x3b, } ,
		{ 0x45, 0x39, 0x53, 0x89, } ,
		{ 0x7f, 0x02, 0xd2, 0xf1, } ,
		{ 0x77, 0xde, 0x96, 0x1a, } ,
	}
	if newState := ShiftRows(state1); newState != expectedState1 {
		t.Errorf("Test 1:\nExpected: %v\nGiven: %v\n", expectedState1, newState)
	}
}

func TestGetCurrentKeyScheduleSubMatrix(t *testing.T) {
	keySchedule := [constants.Nb * (constants.Nr + 1)][4]byte {
					{ 0x2b, 0x7e, 0x15, 0x16, } ,	// w0
					{ 0x28, 0xae, 0xd2, 0xa6, } ,	// w1
					{ 0xab, 0xf7, 0x15, 0x88, } ,	// w2
					{ 0x09, 0xcf, 0x4f, 0x3c, } ,	// w3
					{ 0xa0, 0xfa, 0xfe, 0x17, } ,	// ...
					{ 0x88, 0x54, 0x2c, 0xb1, } ,
					{ 0x23, 0xa3, 0x39, 0x39, } , 
					{ 0x2a, 0x6c, 0x76, 0x05, } ,
					{ 0xf2, 0xc2, 0x95, 0xf2, } ,
					{ 0x7a, 0x96, 0xb9, 0x43, } ,
					{ 0x59, 0x35, 0x80, 0x7a, } ,
					{ 0x73, 0x59, 0xf6, 0x7f, } ,
					{ 0x3d, 0x80, 0x47, 0x7d, } ,
					{ 0x47, 0x16, 0xfe, 0x3e, } ,
					{ 0x1e, 0x23, 0x7e, 0x44, } ,
					{ 0x6d, 0x7a, 0x88, 0x3b, } ,
					{ 0xef, 0x44, 0xa5, 0x41, } ,
					{ 0xa8, 0x52, 0x5b, 0x7f, } ,
					{ 0xb6, 0x71, 0x25, 0x3b, } ,
					{ 0xdb, 0x0b, 0xad, 0x00, } ,
					{ 0xd4, 0xd1, 0xc6, 0xf8, } ,
					{ 0x7c, 0x83, 0x9d, 0x87, } ,
					{ 0xca, 0xf2, 0xb8, 0xbc, } ,
					{ 0x11, 0xf9, 0x15, 0xbc, } ,
					{ 0x6d, 0x88, 0xa3, 0x7a, } ,
					{ 0x11, 0x0b, 0x3e, 0xfd, } ,
					{ 0xdb, 0xf9, 0x86, 0x41, } ,
					{ 0xca, 0x00, 0x93, 0xfd, } ,
					{ 0x4e, 0x54, 0xf7, 0x0e, } ,
					{ 0x5f, 0x5f, 0xc9, 0xf3, } ,
					{ 0x84, 0xa6, 0x4f, 0xb2, } ,
					{ 0x4e, 0xa6, 0xdc, 0x4f, } ,
					{ 0xea, 0xd2, 0x73, 0x21, } ,
					{ 0xb5, 0x8d, 0xba, 0xd2, } ,
					{ 0x31, 0x2b, 0xf5, 0x60, } ,
					{ 0x7f, 0x8d, 0x29, 0x2f, } ,
					{ 0xac, 0x77, 0x66, 0xf3, } ,
					{ 0x19, 0xfa, 0xdc, 0x21, } ,
					{ 0x28, 0xd1, 0x29, 0x41, } ,
					{ 0x57, 0x5c, 0x00, 0x6e, } ,
					{ 0xd0, 0x14, 0xf9, 0xa8, } ,
					{ 0xc9, 0xee, 0x25, 0x89, } ,
					{ 0xe1, 0x3f, 0x0c, 0xc8, } ,
					{ 0xb6, 0x63, 0x0c, 0xa6, } , // w[constants.Nb * (constants.Nr + 1) - 1]
		}

	// 0 - 3
	expected0 := [constants.Nb][4]byte {
					{ 0x2b, 0x7e, 0x15, 0x16, } ,
					{ 0x28, 0xae, 0xd2, 0xa6, } ,
					{ 0xab, 0xf7, 0x15, 0x88, } ,
					{ 0x09, 0xcf, 0x4f, 0x3c, } ,
	}

	// 7 - 10
	expected1 := [constants.Nb][4]byte {
					{ 0x2a, 0x6c, 0x76, 0x05, } ,
					{ 0xf2, 0xc2, 0x95, 0xf2, } ,
					{ 0x7a, 0x96, 0xb9, 0x43, } ,
					{ 0x59, 0x35, 0x80, 0x7a, } ,
	}

	if result := GetCurrentKeyScheduleSubMatrix(keySchedule, 0, 3); result != expected0 {
		t.Errorf("GetCurrentKeyScheduleSubMatrix:\ngot  %x\nwant %x\n", result, expected0)
	}

	if result := GetCurrentKeyScheduleSubMatrix(keySchedule, 7, 10); result != expected1 {
		t.Errorf("GetCurrentKeyScheduleSubMatrix:\ngot  %x\nwant %x\n", result, expected1)
	}
}

func TestInvSubBytes(t *testing.T) {
}

func TestInvShiftRows(t *testing.T) {
}

func TestCipher(t *testing.T) {
}

func TestInvCipher(t *testing.T) {
}
