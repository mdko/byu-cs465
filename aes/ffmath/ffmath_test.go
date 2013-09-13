package ffmath

import "testing"

func TestXTime(t *testing.T) {
		inoutmap := map[byte]byte {
			0x57 : 0xae,
			0xae : 0x47,
			0x47 : 0x8e,
			0x8e : 0x07,
		}
		for in, out := range(inoutmap) {
				if x := XTime(in); x != out {
						t.Errorf("xtime(%x) = %x, want %x", in, x, out)
				}
		}
}

func TestFFMultiply(t *testing.T) {
		// it is a map {answer : map{operand : operand} }
		inoutmap := map[byte]map[byte]byte {
				0xfe : map[byte]byte {
						0x57 : 0x13,
				},
		}
		for ans, operands  := range(inoutmap) {
				for op1, op2 := range(operands) {
						if out := FFMultiply(op1, op2); ans != out {
								t.Errorf("FFMultiply(%x, %x) = %x, want %x", op1, op2, out, ans)
						}
				}
		}
}

func TestFFAdd(t *testing.T) {
		// it is a map {answer : map{operand : operand} }
		inoutmap := map[byte]map[byte]byte {
				0x05 : map[byte]byte {	
						0x04 : 0x01,
				},
				0x01 : map[byte]byte {
						0x03 : 0x02,
				},
				0x03 : map[byte]byte {
						0x04 : 0x07,
				},
		}
		for ans, operands  := range(inoutmap) {
				for op1, op2 := range(operands) {
						if out := FFAdd(op1, op2); ans != out {
								t.Errorf("FFAdd(%x, %x) = %x, want %x", op1, op2, out, ans)
						}
				}
		}
}

func TestMixColumns(t *testing.T) {
	inState0 := [4][4]byte {
			{ 0xd4, 0xbf, 0x5d, 0x30 },
			{ 0xe0, 0xb4, 0x52, 0xae },
			{ 0xb8, 0x41, 0x11, 0xf1 },
			{ 0x1e, 0x27, 0x98, 0xe5 },
	}

	expectedOutState0 := [4][4]byte {
			{ 0x04, 0x66, 0x81, 0xe5 },
			{ 0xe0, 0xcb, 0x19, 0x9a },
			{ 0x48, 0xf8, 0xd3, 0x7a },
			{ 0x28, 0x06, 0x26, 0x4c },
	}

	if realOutState0 := MixColumns(inState0); realOutState0 != expectedOutState0 {
			t.Errorf("Mix Columns:\ngot  %x\nwant %x", realOutState0, expectedOutState0)
	}

	inState1 := [4][4]byte {
			{ 0x49, 0xdb, 0x87, 0x3b },
			{ 0x45, 0x39, 0x53, 0x89 },
			{ 0x7f, 0x02, 0xd2, 0xf1 },
			{ 0x77, 0xde, 0x96, 0x1a },
	}

	expectedOutState1 := [4][4]byte {
			{ 0x58, 0x4d, 0xca, 0xf1 },
			{ 0x1b, 0x4b, 0x5a, 0xac },
			{ 0xdb, 0xe7, 0xca, 0xa8 },
			{ 0x1b, 0x6b, 0xb0, 0xe5 },
	}

	if realOutState1 := MixColumns(inState1); realOutState1 != expectedOutState1 {
			t.Errorf("Mix Columns:\ngot  %x\nwant %x", realOutState1, expectedOutState1)
	}
}

func TestInvMixColumns(t *testing.T) {
	// bd6e7c3d f2b5779e 0b61216e 8b10b689 (C.1 eic after round 2 is_row)
	inState0 := [4][4]byte {
		{ 0xbd, 0x6e, 0x7c, 0x3d, } ,
		{ 0xf2, 0xb5, 0x77, 0x9e, } ,
		{ 0x0b, 0x61, 0x21, 0x6e, } ,
		{ 0x8b, 0x10, 0xb6, 0x89, } ,
	}
	// 4773b91f f72f3543 61cb018e a1e6cf2c
	expectedState0 := [4][4]byte {
		{ 0x47, 0x73, 0xb9, 0x1f, } ,
		{ 0xf7, 0x2f, 0x35, 0x43, } ,
		{ 0x61, 0xcb, 0x01, 0x8e, } ,
		{ 0xa1, 0xe6, 0xcf, 0x2c, } ,
	}
	if result := InvMixColumns(inState0); result != expectedState0 {
			t.Errorf("InvMixColumns:\ngot  %x\nwant %x", result, expectedState0)
	}

	// c81677bc 9b7ac93b 25027992 b0261996 (C.1 eic after round 5 is_row)
	inState1 := [4][4]byte {
		{ 0xc8, 0x16, 0x77, 0xbc, } ,
		{ 0x9b, 0x7a, 0xc9, 0x3b, } ,
		{ 0x25, 0x02, 0x79, 0x92, } ,
		{ 0xb0, 0x26, 0x19, 0x96, } ,
	}
	// 18f78d77 9a93eef4 f6742967 c47f5ffd
	expectedState1 := [4][4]byte {
		{ 0x18, 0xf7, 0x8d, 0x77, } ,
		{ 0x9a, 0x93, 0xee, 0xf4, } ,
		{ 0xf6, 0x74, 0x29, 0x67, } ,
		{ 0xc4, 0x7f, 0x5f, 0xfd, } ,
	}
	if result := InvMixColumns(inState1); result != expectedState1 {
			t.Errorf("InvMixColumns:\ngot  %x\nwant %x", result, expectedState1)
	}
}
