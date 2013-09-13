package cipher

import (
	"github.com/mdko/cs465/aes/constants"
	"github.com/mdko/cs465/aes/ffmath"
)

func Cipher(inState[4][constants.Nb]byte, w[constants.Nb * (constants.Nr + 1)][4]byte) (out_state[4][constants.Nb]byte) {
		var state[4][constants.Nb]byte = inState

		state = AddRoundKey(state, GetCurrentKeyScheduleSubMatrix(w, 0, constants.Nb - 1))

		for round := 1; round < constants.Nr; round++ {
				state = SubBytes(state)
				state = ShiftRows(state)
				state = ffmath.MixColumns(state)
				state = AddRoundKey(state, GetCurrentKeyScheduleSubMatrix(w, round*constants.Nb, (round+1)*constants.Nb - 1))
		}

		state = SubBytes(state)
		state = ShiftRows(state)
		state = AddRoundKey(state, GetCurrentKeyScheduleSubMatrix(w, constants.Nr*constants.Nb, (constants.Nr+1)*constants.Nb - 1))

		out_state = state
		return
}

func AddRoundKey(inState[4][constants.Nb]byte, keySchedule[4][constants.Nb]byte) (out_state[4][constants.Nb]byte) {
		for i := 0; i < 4; i++ {
			for j := 0; j < constants.Nb; j++ {
				out_state[i][j] = ffmath.FFAdd(inState[i][j], keySchedule[i][j])
			}
		}
		return
}

func SubBytes(inState[4][constants.Nb]byte) (out_state[4][constants.Nb]byte) {
		for x, row := range(inState) {
			for y, byt := range(row) {
				// MsB -> row, LsB -> col, index into Sbox
				msB := (byt & 0xF0) >> 4
				lsB := (byt & 0x0F)
				out_state[x][y] = constants.Sbox[msB][lsB]
			}
		}
		return
}

func ShiftRows(inState[4][constants.Nb]byte) (out_state[4][constants.Nb]byte) {
		for x, row := range(inState) {
			for y, byt := range(row) {
			}
		}
		return
}

func InvSubBytes(inState[4][constants.Nb]byte) (out_state[4][constants.Nb]byte) {
		return
}

func InvShiftRows(inState[4][constants.Nb]byte) (out_state[4][constants.Nb]byte) {
		return
}

func GetCurrentKeyScheduleSubMatrix(keySchedule[constants.Nb * (constants.Nr + 1)][4]byte, start int, end int) (roundKeySection[4][constants.Nb]byte) {
		return
}

// This function implements the EqInvCipher function
// func InvCipher(inState[4][constants.Nb]byte, dw[constants.Nb][constants.Nr + 1]byte) (out_state[4][constants.Nb]byte) {
// 		var state[4][constants.Nb] byte
// 		state = inState
// 		AddRoundKey(state, dw[constants.Nr*constants.Nb][(constants.Nr+1)*constants.Nb - 1])
// 		for round := constants.Nr - 1; round >= 1; round-- { // TODO review the bounds here
// 				state = InvSubBytes(state)
// 				state = InvShiftRows(state)
// 				state = InvMixColumns(state)
// 				state = AddRoundKey(state, dw[round*constants.Nb][(round+1)*constants.Nb - 1])
// 		}
// 		state = InvSubBytes(state)
// 		state = InvShiftRows(state)
// 		state = AddRoundKey(state, dw[0][constants.Nb-1])

// 		out_state = state
// 		return
// }
