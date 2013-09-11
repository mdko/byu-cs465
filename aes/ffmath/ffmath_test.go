package ffmath

import "testing"

func TestXtime(t *testing.T) {
		inoutmap := map[byte]byte {
			0x57 : 0xae,
			0xae : 0x47,
			0x47 : 0x8e,
			0x8e : 0x07,
		}
		for in, out := range(inoutmap) {
				if x := xtime(in); x != out {
						t.Errorf("xtime(%v) = %v, want %v", in, x, out)
				}
		}
}
