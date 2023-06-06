package asciiART

import (
	"fmt"
	"strings"
)

func PrintART(str1 string, str2 string, fontname string, colorvalue string) {
	if str2 == "\\n" {
		fmt.Println()
		return
	}
	input_strs := strings.Split(str2, "\\n")
	for _, word := range input_strs {
		if word == "" {
			fmt.Println()
		} else {
			Print_Each_Rune_Line(str1, word, fontname, colorvalue)
		}
	}
	
}
