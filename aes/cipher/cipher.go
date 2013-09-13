package cipher

import (
	"github.com/mdko/cs465/aes/constants"
	"github.com/mdko/cs465/aes/ffmath"
	"github.com/mdko/cs465/aes/debug"
	"fmt"
)

func Cipher(inState[4][constants.Nb]byte, w[constants.Nb * (constants.Nr + 1)][4]byte) (outState[4][constants.Nb]byte) {
		
		debug.DebugPrint("Round Number: input")

		var state[4][constants.Nb]byte = inState		
		debug.DebugPrint(fmt.Sprintf("Start of Round: %x", state))

		roundKeyValue := GetCurrentKeyScheduleSubMatrix(w, 0, constants.Nb - 1)
		debug.DebugPrint(fmt.Sprintf("Round Key Value: %x", roundKeyValue))
		
		state = AddRoundKey(state, roundKeyValue)
		for round := 1; round < constants.Nr; round++ {
				debug.DebugPrint(fmt.Sprintf("Round Number: %v", round))
				debug.DebugPrint(fmt.Sprintf("Start of Round: %x", state))
				
				state = SubBytes(state)
				debug.DebugPrint(fmt.Sprintf("After SubBytes: %x", state))

				state = ShiftRows(state)
				debug.DebugPrint(fmt.Sprintf("After ShiftRows: %x", state))

				state = ffmath.MixColumns(state)
				debug.DebugPrint(fmt.Sprintf("After MixColumns: %x", state))

				roundKeyValue = GetCurrentKeyScheduleSubMatrix(w, round*constants.Nb, (round+1)*constants.Nb - 1)
				debug.DebugPrint(fmt.Sprintf("Round Key Value: %x", roundKeyValue))
				state = AddRoundKey(state, roundKeyValue)
				
		}
		
		debug.DebugPrint(fmt.Sprintf("Round Number: %v", constants.Nr))
		debug.DebugPrint(fmt.Sprintf("Start of Round: %x", state))
		
		state = SubBytes(state)
		debug.DebugPrint(fmt.Sprintf("After SubBytes: %x", state))

		state = ShiftRows(state)
		debug.DebugPrint(fmt.Sprintf("After ShiftRows: %x", state))

		roundKeyValue = GetCurrentKeyScheduleSubMatrix(w, constants.Nr*constants.Nb, (constants.Nr+1)*constants.Nb - 1)
		debug.DebugPrint(fmt.Sprintf("Round Key Value: %x", roundKeyValue))
		state = AddRoundKey(state, roundKeyValue)
		
		outState = state
		return
}

func AddRoundKey(inState[4][constants.Nb]byte, keySchedule[4][constants.Nb]byte) (outState[4][constants.Nb]byte) {
		for i := 0; i < 4; i++ {
			for j := 0; j < constants.Nb; j++ {
				outState[i][j] = ffmath.FFAdd(inState[i][j], keySchedule[i][j])
			}
		}
		return
}

func SubBytes(inState[4][constants.Nb]byte) (outState[4][constants.Nb]byte) {
		for x, row := range(inState) {
			for y, byt := range(row) {
				// MsB -> row, LsB -> col, index into Sbox
				msB := (byt & 0xF0) >> 4
				lsB := (byt & 0x0F)
				outState[x][y] = constants.Sbox[msB][lsB]
			}
		}
		return
}

func ShiftRows(inState[4][constants.Nb]byte) (outState[4][constants.Nb]byte) {
		for r := 0; r < 4; r++ {
			for c := 0; c < constants.Nb; c++ {
				outState[c][r] = inState[(c + r) % constants.Nb][r] // I needed to switch the [] order due to how I have set up my matrix (the columns are each row in my matrix)
			}
		}
		return
}

func InvSubBytes(inState[4][constants.Nb]byte) (outState[4][constants.Nb]byte) {
		return
}

func InvShiftRows(inState[4][constants.Nb]byte) (outState[4][constants.Nb]byte) {
		return
}

// start - 1st column
// end - last column (inclusive)
// ie (..., 0, 3) means columns 0, 1, 2, and 3
func GetCurrentKeyScheduleSubMatrix(keySchedule[constants.Nb * (constants.Nr + 1)][4]byte, start int, end int) (roundKeySection[4][constants.Nb]byte) {
		for i := start; i <= end; i++ {
				roundKeySection[i - start] = keySchedule[i]
		}
		return
}

// This function implements the EqInvCipher function
// func InvCipher(inState[4][constants.Nb]byte, dw[constants.Nb][constants.Nr + 1]byte) (outState[4][constants.Nb]byte) {
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

// 		outState = state
// 		return
// }
