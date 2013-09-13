package debug

import (
		"fmt"
		"github.com/mdko/cs465/aes/constants"
)

func DebugPrint(message string) {
	if constants.Debug == true {
			fmt.Printf("DEBUG: %s\n", message)
	}
}
