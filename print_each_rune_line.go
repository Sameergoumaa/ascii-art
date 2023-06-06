package asciiART

import "fmt"

func Print_Each_Rune_Line(str1 string, str2 string, fontname string, colorvalue string) {
	// Iterate through eight lines.
	if str1 == "" {
		for i := 0; i < 8; i++ {
			// Iterate through each character in the input string
			str2_len := len(str2)
			for idx := 0; idx < str2_len; idx++ {
				char := str2[idx]
				// Check for the backslash (since this is taken from os arguments you have to read it like this '\\') and then read the letter after it
				if char == '\\' {
					if idx < len(str2)-1 {
						// Apply tab if 't' appears
						if str2[idx+1] == 't' {
							fmt.Print("\t")
							idx++
						} else {
							PrintFileLine(str1,MapART(rune(char))+i, MapFont(fontname), colorvalue)
						}
					}
				} else {
					PrintFileLine(str1,MapART(rune(char))+i, MapFont(fontname), colorvalue)
				}
			}
			fmt.Print("\n") // prints newline to start printing the rest of the art
		}
	} else {
		
		for i := 0; i < 8; i++ {
			// Iterate through each character in the input string
			str2_len := len(str2)
			for idx := 0; idx < str2_len; idx++ {
				char := str2[idx]
				// Check for the backslash (since this is taken from os arguments you have to read it like this '\\') and then read the letter after it
				if char == '\\' {
					if idx < len(str2)-1 {
						// Apply tab if 't' appears
						if str2[idx+1] == 't' {
							fmt.Print("\t")
							idx++
						} else {
							PrintFileLine(str1,MapART(rune(char))+i, MapFont(fontname), colorvalue)
						}
					}
				} else {
					PrintFileLine(str1,MapART(rune(char))+i, MapFont(fontname), colorvalue)
				}
			}
			fmt.Print("\n") // prints newline to start printing the rest of the art

			str1_len := len(str1)
			for idx := 0; idx < str1_len; idx++ {
				char := str1[idx]
				if idx < len(str1)-1 {
					PrintFileLine(str1,MapART(rune(char))+i, MapFont(fontname), colorvalue)
				} else {
					PrintFileLine(str1,MapART(rune(char))+i, MapFont(fontname), colorvalue)
				}
			}
		}
	}
}
