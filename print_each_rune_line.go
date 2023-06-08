package asciiART

import "fmt"

func Print_Each_Rune_Line(replacestr string, message string, fontname string, colorvalue string) {
	// Iterate through eight lines.

	if replacestr == "" {
		for i := 0; i < 8; i++ {
			// Iterate through each character in the input string
			str2_len := len(message)
			for idx := 0; idx < str2_len; idx++ {
				char := message[idx]
				// Check for the backslash (since this is taken from os arguments you have to read it like this '\\') and then read the letter after it
				if char == '\\' {
					if idx < len(message)-1 {
						// Apply tab if 't' appears
						if message[idx+1] == 't' {
							fmt.Print("\t")
							idx++
						} else {
							PrintFileLine(0, MapART(rune(char))+i, MapFont(fontname), colorvalue)
						}
					}
				} else {
					PrintFileLine(0, MapART(rune(char))+i, MapFont(fontname), colorvalue)
				}
			}
			fmt.Print("\n") // prints newline to start printing the rest of the art
		}

	} else {
		var idx int
		var idx2 int
		str1_len := len(message)
		str2_len := len(replacestr)
		for i := 0; i < 8; i++ {
			// Iterate through each character in the input string
			for idx = 0; idx < str1_len; idx++ {
				char := message[idx]
				for idx2 = 0; idx2 <= str2_len -1; idx2++ {
					char2 := replacestr[idx2]
					//PrintFileLine(MapART(rune(char2))+i, 0, MapFont(fontname), colorvalue)
					if char == '\\' {
						if idx < len(message)-1 {
							// Apply tab if 't' appears
							if message[idx+1] == 't' {
								fmt.Print("\t")
								idx++
							} else {
								PrintFileLine(MapART(rune(char2))+i, MapART(rune(char))+i, MapFont(fontname), colorvalue)
							}
						}
					} else {
						PrintFileLine(MapART(rune(char2))+i, MapART(rune(char))+i, MapFont(fontname), colorvalue)
					}
				}
			}
			fmt.Print("\n") // prints newline to start printing the rest of the art
		}
	}
}
