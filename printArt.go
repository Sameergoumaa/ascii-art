package asciiART

import (
	"fmt"
	"strings"
)

func PrintART(str string) {
	input_strs := strings.Split(str, "\\n")
	for _, word := range input_strs {
		if word == "" {
			fmt.Print("\n")
		} else {

			Print_Each_Rune_Line(word)
		}
	}
}
