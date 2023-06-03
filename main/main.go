package main

import (
	"asciiART"
	"os"
	// "fmt"
)

func main() {
	//asciiART.PrintART("We Gee\n\n\nChapati\t\t\ttuuna")
	asciiART.PrintART(os.Args[1])
}
