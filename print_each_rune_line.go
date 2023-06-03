package asciiART

import "fmt"

func Print_Each_Rune_Line(str string) {

	// Iterate through eight lines
	for i := 0; i < 8; i++ {
		// Iterate through each character in the input string
		str_len := len(str)
		for idx := 0; idx < str_len; idx++ {
			char := str[idx]
			// Check for the backslash (since this is taken from os arguments you have to read it like this '\\') and then read the letter after it
			if char == '\\' {
				if idx < len(str)-1 {
					// Apply tab if 't' appears
					if str[idx+1] == 't' {
						fmt.Print("\t")
						idx++
					}
				}
			} else {

				PrintFileLine(MapART(byte(char))+i, "../standard.txt")
			}
		}
		fmt.Print("\n") // prints newline to start printing the rest of the art
	}
}
