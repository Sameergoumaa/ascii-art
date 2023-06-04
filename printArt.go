package asciiART

import (
	"fmt"
	"strings"
)

func PrintART(str string, fontname string) {
	if str == "\\n" {
		fmt.Println()
		return
	}
	input_strs := strings.Split(str, "\\n")
	for _, word := range input_strs {
		if word == "" {
			fmt.Println()
		} else {
			Print_Each_Rune_Line(word, fontname)
		}
	}
}
