package asciiART

import "fmt"

func Print_Each_Rune_Line(str string) {

	// Iterate through eight lines
	for i := 0; i < 8; i++ {
		// Iterate through each character in the input string
		for _, char := range str {
			PrintFileLine(MapART(byte(char))+i, "../standard.txt")
			if char == '\t' {

				fmt.Print("\t")
			}
		}
		fmt.Print("\n") // prints newline to start printing the rest of the art
	}
}
