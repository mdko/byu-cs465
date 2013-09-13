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

	//a := make([][]int, 4)	// slices
	////for i := range a {
	////		a[i] = make([]int, 4)
	////}
	//aAll := make([]int, 4*4)
	//for i := range a {
	//		a[i], aAll = aAll[:4], aAll[4:]
	//}
	//ax := len(a)
	//ay := len(a[0])
	//fmt.Println(ax, ay)

	//var b [4][5]int			// arrays
	//x := len(b)				// number in 1st dimension
	//y := len(b[0])			// number of elements in each element (2nd dim)
	//fmt.Println(x, y)
}

func TestMixColumns(t *testing.T) {
	// inState0 := [4][4]byte {
	// 		{ 0xd4, 0xe0, 0xb8, 0x1e },
	// 		{ 0xbf, 0xb4, 0x41, 0x27 },
	// 		{ 0x5d, 0x52, 0x11, 0x98 },
	// 		{ 0x30, 0xae, 0xf1, 0xe5 },
	// }

	// expectedOutState0 := [4][4]byte {
	// 		{ 0x04, 0xe0, 0x48, 0x28 },
	// 		{ 0x66, 0xcb, 0xf8, 0x06 },
	// 		{ 0x81, 0x19, 0xd3, 0x26 },
	// 		{ 0xe5, 0x9a, 0x7a, 0x4c },
	// }
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
			t.Errorf("FFMultiply:\ngot  %x\nwant %x", realOutState0, expectedOutState0)
	}

	// inState1 := [4][4]byte {
	// 		{ 0x49, 0x45, 0x7f, 0x77 },
	// 		{ 0xdb, 0x39, 0x02, 0xde },
	// 		{ 0x87, 0x53, 0xd2, 0x96 },
	// 		{ 0x3b, 0x89, 0xf1, 0x1a },
	// }

	// expectedOutState1 := [4][4]byte {
	// 		{ 0x58, 0x1b, 0xdb, 0x1b },
	// 		{ 0x4d, 0x4b, 0xe7, 0x6b },
	// 		{ 0xca, 0x5a, 0xca, 0xb0 },
	// 		{ 0xf1, 0xac, 0xa8, 0xe5 },
	// }
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
			t.Errorf("FFMultiply:\ngot  %x\nwant %x", realOutState1, expectedOutState1)
	}
}
