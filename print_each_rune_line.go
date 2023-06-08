package asciiART

import "fmt"

func Print_Each_Rune_Line(replacestr string, message string, fontname string, colorvalue string) {
	// Iterate through eight lines.

	if replacestr == "" {
		for i := 0; i < 8; i++ {
			// Iterate through each character in the input string
			orgstr_len := len(message)
			for orgstridx := 0; orgstridx < orgstr_len; orgstridx++ {
				char := message[orgstridx]
				// Check for the backslash (since this is taken from os arguments you have to read it like this '\\') and then read the letter after it
				if char == '\\' {
					if orgstridx < len(message)-1 {
						// Apply tab if 't' appears
						if replacestr[orgstridx+1] == 't' {
							fmt.Print("\t")
							orgstridx++
						} else {
							PrintFileLine1(replacestr, MapART(rune(char))+i, MapFont(fontname), colorvalue)
						}
					}
				} else {
					PrintFileLine1(replacestr, MapART(rune(char))+i, MapFont(fontname), colorvalue)
				}
			}
			fmt.Print("\n") // prints newline to start printing the rest of the art
		}
	} else {
		for i := 0; i < 8; i++ {
			// Iterate through each character in the input string
			orgstr_len := len(message)
			for orgstridx := 0; orgstridx < orgstr_len; orgstridx++ {
				char := message[orgstridx]
				// Check for the backslash (since this is taken from os arguments you have to read it like this '\\') and then read the letter after it
				if char == '\\' {
					if orgstridx < len(message)-1 {
						// Apply tab if 't' appears
						if replacestr[orgstridx+1] == 't' {
							fmt.Print("\t")
							orgstridx++
						} else {
							PrintFileLine1(replacestr, MapART(rune(char))+i, MapFont(fontname), colorvalue)
						}
					}
				} else {
					PrintFileLine1(replacestr, MapART(rune(char))+i, MapFont(fontname), colorvalue)
				}
			}
			rplstr_len := len(replacestr)
			for rplstridx := 0; rplstridx < rplstr_len; rplstridx++ {
				char2 := replacestr[rplstridx]
				// Check for the backslash (since this is taken from os arguments you have to read it like this '\\') and then read the letter after it
				PrintFileLine2(MapART(rune(char2))+i, MapFont(fontname), colorvalue)
			}
			fmt.Print("\n") // prints newline to start printing the rest of the art
		}
	}
}
