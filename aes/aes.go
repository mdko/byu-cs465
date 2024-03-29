package main

import (
		"github.com/mdko/cs465/aes/cipher"
		"github.com/mdko/cs465/aes/constants"
		"github.com/mdko/cs465/aes/keyexpansion"
		"fmt"
)

func main() {
		//in := os.Args[1]
		var in [4][constants.Nb]byte
		var cipherKey [constants.Nk][4]byte
		var out [4][constants.Nb]byte
		var roundKey [constants.Nb * (constants.Nr + 1)][4]byte
		var decryption [4][constants.Nb]byte




		// -----------------------------------------------------------------------
		// -----------------------------------------------------------------------
		// Appendix B
		// 128-bit example from spec
		// Note: need to change Nk and Nr in constants.go to 4 and 4, respectively
		// -----------------------------------------------------------------------
		
		//in = [4][constants.Nb]byte {
		//	{ 0x32, 0x43, 0xf6, 0xa8, } ,
		//	{ 0x88, 0x5a, 0x30, 0x8d, } ,
		//	{ 0x31, 0x31, 0x98, 0xa2, } ,
		//	{ 0xe0, 0x37, 0x07, 0x34, } ,
		//}
		//cipherKey = [constants.Nk][4]byte {
		//	{ 0x2b, 0x7e, 0x15, 0x16, } ,
		//	{ 0x28, 0xae, 0xd2, 0xa6, } ,
		//	{ 0xab, 0xf7, 0x15, 0x88, } ,
		//	{ 0x09, 0xcf, 0x4f, 0x3c, } ,
		//}
		//roundKey = keyexpansion.KeyExpansion(cipherKey)
		//out = cipher.Cipher(in,roundKey)
		//fmt.Printf("In:  %x\nKey: %x\nOut: %x\n", in, cipherKey, out)




		// -----------------------------------------------------------------------
		// -----------------------------------------------------------------------
		// Appendix C
		// 128-bit example from spec
		// Note: need to change Nk and Nr in constants.go to 4 and 4, respectively
		// -----------------------------------------------------------------------

		//in = [4][constants.Nb]byte {
		//	{ 0x00, 0x11, 0x22, 0x33, } ,
		//	{ 0x44, 0x55, 0x66, 0x77, } ,
		//	{ 0x88, 0x99, 0xaa, 0xbb, } ,
		//	{ 0xcc, 0xdd, 0xee, 0xff, } ,
		//}
		//cipherKey = [constants.Nk][4]byte {
		//	{ 0x00, 0x01, 0x02, 0x03, } ,
		//	{ 0x04, 0x05, 0x06, 0x07, } ,
		//	{ 0x08, 0x09, 0x0a, 0x0b, } ,
		//	{ 0x0c, 0x0d, 0x0e, 0x0f, } ,
		//}
		//roundKey = keyexpansion.KeyExpansion(cipherKey)
		//out = cipher.Cipher(in, roundKey)
		//fmt.Printf("Encryption:\nIn:  %x\nKey: %x\nOut: %x\n", in, cipherKey, out)
		//decryption = cipher.InvCipher(out, roundKey)
		//fmt.Printf("Decryption:\nIn:  %x\nKey: %x\nOut: %x\n", out, cipherKey, decryption)




		// -----------------------------------------------------------------------
		// ------------------------------------------------------------------------
		// Appendix C
		// 192-bit example from spec
		// Note: need to change Nk and Nr in constants.go to 6 and 12, respectively
		// -----------------------------------------------------------------------
		
		// in = [4][constants.Nb]byte {
		// 	{ 0x00, 0x11, 0x22, 0x33, } ,
		// 	{ 0x44, 0x55, 0x66, 0x77, } ,
		// 	{ 0x88, 0x99, 0xaa, 0xbb, } ,
		// 	{ 0xcc, 0xdd, 0xee, 0xff, } ,
		// }
		// cipherKey = [constants.Nk][4]byte {
		// 	{ 0x00, 0x01, 0x02, 0x03, } ,
		// 	{ 0x04, 0x05, 0x06, 0x07, } ,
		// 	{ 0x08, 0x09, 0x0a, 0x0b, } ,
		// 	{ 0x0c, 0x0d, 0x0e, 0x0f, } ,
		// 	{ 0x10, 0x11, 0x12, 0x13, } ,
		// 	{ 0x14, 0x15, 0x16, 0x17, } ,
		// }
		// roundKey = keyexpansion.KeyExpansion(cipherKey)
		// out = cipher.Cipher(in, roundKey)
		// fmt.Printf("Encryption:\nIn:  %x\nKey: %x\nOut: %x\n", in, cipherKey, out)
		// decryption = cipher.InvCipher(out, roundKey)
		// fmt.Printf("Decryption:\nIn:  %x\nKey: %x\nOut: %x\n", out, cipherKey, decryption)




		// -----------------------------------------------------------------------
		// ------------------------------------------------------------------------`
		// Appendix C
		// 256-bit example from spec
		// Note: need to change Nk and Nr in constants.go to 8 and 14, respectively
		// -----------------------------------------------------------------------
		
		 in = [4][constants.Nb]byte {
		 	{ 0x00, 0x11, 0x22, 0x33, } ,
		 	{ 0x44, 0x55, 0x66, 0x77, } ,
		 	{ 0x88, 0x99, 0xaa, 0xbb, } ,
		 	{ 0xcc, 0xdd, 0xee, 0xff, } ,
		 }
		 cipherKey = [constants.Nk][4]byte {
		 	{ 0x00, 0x01, 0x02, 0x03, } ,
		 	{ 0x04, 0x05, 0x06, 0x07, } ,
		 	{ 0x08, 0x09, 0x0a, 0x0b, } ,
		 	{ 0x0c, 0x0d, 0x0e, 0x0f, } ,
		 	{ 0x10, 0x11, 0x12, 0x13, } ,
		 	{ 0x14, 0x15, 0x16, 0x17, } ,
		 	{ 0x18, 0x19, 0x1a, 0x1b, } ,
		 	{ 0x1c, 0x1d, 0x1e, 0x1f, } ,
		 }
		 roundKey = keyexpansion.KeyExpansion(cipherKey)
		 out = cipher.Cipher(in, roundKey)
		 fmt.Printf("Encryption:\nIn:  %x\nKey: %x\nOut: %x\n", in, cipherKey, out)
		 decryption = cipher.InvCipher(out, roundKey)
		 fmt.Printf("Decryption:\nIn:  %x\nKey: %x\nOut: %x\n", out, cipherKey, decryption)
}
