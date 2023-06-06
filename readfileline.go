package asciiART

import (
	"bufio"
	"fmt"
	"os"
)

func PrintFileLine(str1 string, lineNumber int, filePath string, colorvalue string) {
	file, err := os.Open(filePath)
	// If there is an error, then handle it
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(file)

	// Read line by line
	if str1 == "" {
		line := ""
		lineCount := 0 // A counter used to stop at specific line
		for scanner.Scan() {
			lineCount++
			// save the line and stop the loop
			if lineCount == lineNumber {
				line = scanner.Text()
				break
			}
		}
		Colorize(str1, (colorvalue), line)
	} else {
		line := ""
		lineCount := 0 // A counter used to stop at specific line
		for scanner.Scan() {
			lineCount++
			// save the line and stop the loop
			if lineCount == lineNumber {
				line = scanner.Text()
				break
			}
		}
		Colorize(line, (colorvalue), line)
	}
	
	//fmt.Print(line)
	// Check for any errors during scanning
	if err := scanner.Err(); err != nil {
		fmt.Println("Error scanning file:", err)
		return
	}
}
