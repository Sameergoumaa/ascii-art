package asciiART

import (
	"bufio"
	"fmt"
	"os"
)

var messageline string
var messageline2 string
var rplline string

func PrintFileLine1(replacestr string, lineNumber int, filePath string, colorvalue string) {

	file, err := os.Open(filePath)
	// If there is an error, then handle it
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()
	// Create a scanner to read the file messageline by messageline
	scanner := bufio.NewScanner(file)
	
	if replacestr == "" {
		// Read messageline by messageline
		lineCount := 0 // A counter used to stop at specific messageline
		for scanner.Scan() {
			lineCount++
			// save the messageline and stop the loop
			if lineCount == lineNumber {
				messageline = scanner.Text()
				break
			}
		}
		Colorize("", (colorvalue), messageline)
	} else {
		lineCount := 0 // A counter used to stop at specific messageline
		for scanner.Scan() {
			lineCount++
			// save the messageline and stop the loop
			if lineCount == lineNumber {
				messageline2 = scanner.Text()
				break
			}
		}
		messageline = messageline2
	}
}

func PrintFileLine2(lineNumber int, filePath string, colorvalue string) {
	file2, err := os.Open(filePath)
	// If there is an error, then handle it
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file2.Close()

	// Create a scanner to read the file messageline by messageline
	scanner1 := bufio.NewScanner(file2)
	lineCount1 := 0 // A counter used to stop at specific messageline

	for scanner1.Scan() {
		lineCount1++
		// save the messageline and stop the loop
		if lineCount1 == lineNumber {
			rplline = scanner1.Text()
			break
		}
	}
	fmt.Print(messageline2)
	return
	Colorize(rplline, messageline2, colorvalue)
	// Check for any errors during scanning
	if err := scanner1.Err(); err != nil {
		fmt.Println("Error scanning file:", err)
		return
	}
}
