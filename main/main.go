package main

import (
	"asciiART"
	"fmt"
	"os"
	// "fmt"
)

func main() {
	//asciiART.PrintART("We Gee\n\n\nChapati\t\t\ttuuna")
	if len(os.Args) < 2 {
		fmt.Print()
	} else if len(os.Args) >= 2 && len(os.Args) < 3 {
		if os.Args[1] == "" {
			fmt.Print()
		} else {
			asciiART.PrintART(os.Args[1], "standard")
		}
	} else if len(os.Args) >= 3 {
		asciiART.PrintART(os.Args[1], os.Args[2])
	}
}
