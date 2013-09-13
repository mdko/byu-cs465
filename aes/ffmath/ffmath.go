package ffmath

import (
		"fmt"
)

const Nb = 4
const DEBUG = false

func XTime(in byte) (out byte) {
			overflow := (in & 0x80) == 0x80	// The msb is set, so multiplying by two will overflow into a 2nd byte
			out = in << 1
			if bool(overflow) {
					out = out ^ 0x1b
			}
			return out
}

func MixColumns(in_state[4][Nb]byte) [4][Nb]byte {
		var out_state[4][Nb]byte
		for c := 0; c < Nb; c++ { // for each column (0 to 3)
				out_state[0][c] = FFAdd(
							FFAdd(
								FFAdd(	FFMultiply(0x02, in_state[0][c]),
										FFMultiply(0x03, in_state[1][c])),
								in_state[2][c]),
							in_state[3][c])
				out_state[1][c] = FFAdd(
							FFAdd(
								FFAdd(	in_state[0][c],
										FFMultiply(0x02, in_state[1][c])),
								FFMultiply(0x03, in_state[2][c])),
							in_state[3][c])
				out_state[2][c] = FFAdd(
							FFAdd(
								FFAdd(	in_state[0][c],
										in_state[1][c]),
								FFMultiply(0x02, in_state[2][c])),
							FFMultiply(0x03, in_state[3][c]))
				out_state[3][c] = FFAdd(
							FFAdd(
								FFAdd(	FFMultiply(0x03, in_state[0][c]),
										in_state[1][c]),
								in_state[2][c]),
							FFMultiply(0x02, in_state[3][c]))
		}
		return out_state
}

// TODO implement
func InvMixColumns(in_state[4][Nb]byte) (out_state[4][Nb]byte) {
	return
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

						if DEBUG {
								fmt.Printf("Current bit %x\n", bitmask)
								fmt.Printf("Current result to add %x\n", current_result)
						}
				}
		}
		return result
}
