package main

import (
		"github.com/mdko/cs465/aes/cipher"
)

func main() {
		//in := os.Args[1]
		in := [4][constants.Nb]byte {
			{ 0x32, 0x43, 0xf6, 0xa8, } ,
			{ 0x88, 0x5a, 0x30, 0x8d, } ,
			{ 0x31, 0x31, 0x98, 0xa2, } ,
			{ 0xe0, 0x37, 0x07, 0x34, } ,
		}
		cipherKey := [constants.Nk][4]byte {
			{ 0x2b, 0x7e, 0x15, 0x16, } ,
			{ 0x28, 0xae, 0xd2, 0xa6, } ,
			{ 0xab, 0xf7, 0x15, 0x88, } ,
			{ 0x09, 0xcf, 0x4f, 0x3c, } ,
		}
		roundKey := keyexpansion.KeyExpansion(cipherKey)
		out := cipher.Cipher(in,roundKey)
		fmt.Println("Out: %v\n", out)
}