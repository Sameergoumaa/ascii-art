package asciiART

import (
	"bufio"
	"fmt"
	"os"
)

func PrintFileLine(lineNumber2 int, lineNumber int, filePath string, colorvalue string) {

	file, err := os.Open(filePath)
	// If there is an error, then handle it
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Create a scanner to read the file messageline by messageline
	scanner1 := bufio.NewScanner(file)
	scanner2 := bufio.NewScanner(file)

	// Read messageline by messageline
	if lineNumber2 == 0 {
		messageline := ""
		lineCount := 0 // A counter used to stop at specific messageline
		for scanner1.Scan() {
			lineCount++
			// save the messageline and stop the loop
			if lineCount == lineNumber {
				messageline = scanner1.Text()
				break
			}
		}
		Colorize("", (colorvalue), messageline)
	} else {
		messageline := ""
		rplline := ""
		lineCount1 := 0 // A counter used to stop at specific messageline
		lineCount2 := 0 // A counter used to stop at specific messageline
		for scanner1.Scan() {
			lineCount1++
			// save the messageline and stop the loop
			if lineCount1 == lineNumber {
				messageline = scanner1.Text()
				break
			}
		}

		for scanner2.Scan() {
			lineCount2++
			// save the messageline and stop the loop
			if lineCount2 == lineNumber2 {
				rplline = scanner2.Text()
				break
			}
		}
		Colorize(rplline, (colorvalue), messageline)
	}

	// Check for any errors during scanning
	if err := scanner1.Err(); err != nil {
		fmt.Println("Error scanning file:", err)
		return
	}
}
