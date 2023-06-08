package asciiART

import (
	"fmt"
	"strings"
)

func PrintART(repplacestr string, message string, fontname string, colorvalue string) {
	if message == "\\n" {
		fmt.Println()
		return
	}
	if repplacestr == "" {
		input_strs := strings.Split(message, "\\n")
		for _, word := range input_strs {
			if word == "" {
				fmt.Println()
			} else {
				Print_Each_Rune_Line(repplacestr, word, fontname, colorvalue)
			}
		}
	} else {
		final := ""
		final_rpl := ""
		input_strs := strings.Split(message, "\\n")
		for _, word := range input_strs {
			if word == "" {
				fmt.Println()
			} else {
				final = final+word
				//Print_Each_Rune_Line("", word, fontname, colorvalue)
			}
		}
		repl_str := strings.Split(repplacestr, "\\n")
		for _, word := range repl_str {
			final_rpl = final_rpl + word	
		}
		Print_Each_Rune_Line(final_rpl, final, fontname, colorvalue)
	}
}
