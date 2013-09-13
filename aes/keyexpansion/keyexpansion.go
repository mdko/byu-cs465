package keyexpansion

import (
		"github.com/mdko/cs465/aes/constants"
		"fmt"
)

const (
		DEBUG = false
)

func KeyExpansion(key[constants.Nk][4]byte) (w[constants.Nb * (constants.Nr + 1)][4]byte) {
		var temp[4]byte
		for i := 0; i < constants.Nk; i++ {
				for j := 0; j < 4; j++ {
						w[i][j] = key[i][j]
				}
		}

		for i := constants.Nk; i < constants.Nb * (constants.Nr + 1); i++ {
			temp = w[i - 1]
				if i % constants.Nk == 0 {
						temp = XORWords(SubWord(RotWord(temp)), MakeByteArray(constants.Rcon[i/constants.Nk]))
				} else if (constants.Nk > 6) && ((i % constants.Nk) == 4) {
						temp = SubWord(temp)
				}
				w[i] = XORWords(w[i - constants.Nk], temp)
				if DEBUG {
						fmt.Printf("%x = %x XOR %x\n", w[i], temp, w[i - constants.Nk])
				}
		}
		return
}

func SubWord(inWord[4] byte) (outWord[4] byte) {
		for loc, byt := range(inWord) {
				row := byt & 0x0F
				col := (byt & 0xF0) / 16
				outWord[loc] = constants.Sbox[col][row] // TODO review this indexing, it works this way but I am confused a little as to how the matrix is set up (index into column, then row of that column)
		}
		return
}

func XORWords(inWordL[4] byte, inWordR[4] byte) (outWord[4] byte) {
		for i := 0; i < len(outWord); i++ {
			outWord[i] = inWordL[i] ^ inWordR[i]
		}
		return
}

func MakeByteArray(word uint32) (outWord[4]byte) {
		outWord[0] = byte((word & 0xFF000000) >> 24)
		outWord[1] = byte((word & 0x00FF0000) >> 16)
		outWord[2] = byte((word & 0x0000FF00) >> 8)
		outWord[3] = byte(word & 0x000000FF)
		return
}

func RotWord(inWord[4] byte) (outWord[4] byte) {
		// The following looks weird, but given a 32-bit word like
		// 0x09cf4f3c, it is stored in in array like {0x09, 0xcf, 0xf4, 0x3c},
		// so byte 0 in the array is the most significant byte in the real word
		outWord[0] = inWord[1]
		outWord[1] = inWord[2]
		outWord[2] = inWord[3]
		outWord[3] = inWord[0]
		return
}