package ffmath

import (
		"github.com/mdko/cs465/aes/constants"
)

func XTime(in byte) (out byte) {
			overflow := (in & 0x80) == 0x80	// The msb is set, so multiplying by two will overflow into a 2nd byte
			out = in << 1
			if bool(overflow) {
					out = out ^ 0x1b
			}
			return out
}

// Note to self: in a state like
// [ ] [ ] [ ] [ ]
// [ ] [ ] [ ] [ ]
// [ ] [ ] [ ] [ ]
// [ ] [ ] [ ] [ ]
//  |	|	|	array[3]
//	|	|	array[2]
//  |   array[1]
//  array[0]
// I.e. in the Go represenation
// state := [4][4]byte {
//			{ 0x49, 0xdb, 0x87, 0x3b }, <--array[0] is a *column* as shown in the spec (each row holds a column, backwards but oh well)
//			{ 0x45, 0x39, 0x53, 0x89 },
//			{ 0x7f, 0x02, 0xd2, 0xf1 },
//			{ 0x77, 0xde, 0x96, 0x1a },
//	}
func MixColumns(in_state[4][constants.Nb]byte) [4][constants.Nb]byte {
		var out_state[4][constants.Nb]byte
		for c := 0; c < constants.Nb; c++ { // for each column (0 to 3)
				out_state[c][0] = FFAdd(
							FFAdd(
								FFAdd(	FFMultiply(0x02, in_state[c][0]),
										FFMultiply(0x03, in_state[c][1])),
								in_state[c][2]),
							in_state[c][3])
				out_state[c][1] = FFAdd(
							FFAdd(
								FFAdd(	in_state[c][0],
										FFMultiply(0x02, in_state[c][1])),
								FFMultiply(0x03, in_state[c][2])),
							in_state[c][3])
				out_state[c][2] = FFAdd(
							FFAdd(
								FFAdd(	in_state[c][0],
										in_state[c][1]),
								FFMultiply(0x02, in_state[c][2])),
							FFMultiply(0x03, in_state[c][3]))
				out_state[c][3] = FFAdd(
							FFAdd(
								FFAdd(	FFMultiply(0x03, in_state[c][0]),
										in_state[c][1]),
								in_state[c][2]),
							FFMultiply(0x02, in_state[c][3]))
		}
		return out_state
}

func InvMixColumns(in_state[4][constants.Nb]byte) [4][constants.Nb]byte {
	var out_state[4][constants.Nb]byte
		for c := 0; c < constants.Nb; c++ { // for each column (0 to 3)
				out_state[c][0] = FFAdd(
							FFAdd(
								FFAdd(	FFMultiply(0x0e, in_state[c][0]),
										FFMultiply(0x0b, in_state[c][1])),
								FFMultiply(0x0d, in_state[c][2])),
							FFMultiply(0x09, in_state[c][3]))
				out_state[c][1] = FFAdd(
							FFAdd(
								FFAdd(	FFMultiply(0x09, in_state[c][0]),
										FFMultiply(0x0e, in_state[c][1])),
								FFMultiply(0x0b, in_state[c][2])),
							FFMultiply(0x0d, in_state[c][3]))
				out_state[c][2] = FFAdd(
							FFAdd(
								FFAdd(	FFMultiply(0x0d, in_state[c][0]),
										FFMultiply(0x09, in_state[c][1])),
								FFMultiply(0x0e, in_state[c][2])),
							FFMultiply(0x0b, in_state[c][3]))
				out_state[c][3] = FFAdd(
							FFAdd(
								FFAdd(	FFMultiply(0x0b, in_state[c][0]),
										FFMultiply(0x0d, in_state[c][1])),
								FFMultiply(0x09, in_state[c][2])),
							FFMultiply(0x0e, in_state[c][3]))
		}
		return out_state
}

func FFAdd(l_byte byte, r_byte byte) byte {
		return l_byte ^ r_byte // XOR
}

func FFMultiply(l_byte byte, r_byte byte) byte {
		// for each 1 bit in the binary representation of the r_byte,
		// do the following (ie, for 0x13 === 0001 0011, ceiling equals
		// 16, 2, and 1.) The following simulates the associative rule
		// described in the slides:
		var result byte
		for bit := 0; bit < 8; bit++ {				// for each bit in r_byte
				bitmask := 0x01 << uint8(bit)
				if r_byte & byte(bitmask) == byte(bitmask) {		// see which ones are high
						ceiling := bitmask					// if high, iteratively do xtime until you can't go any higher for this particular bit
						current_result := l_byte
						i := 0x01
						for i < ceiling {
								current_result = XTime(current_result)
								i = i << 1
						}
						result = FFAdd(result, current_result)
				}
		}
		return result
}
