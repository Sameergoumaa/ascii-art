package main

import (
	"asciiART"
	"flag"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Print()
	} else if len(os.Args) >= 2 && len(os.Args) < 3 {
		if os.Args[1] == "" {
			fmt.Print()
		} else {
			asciiART.PrintART("",os.Args[1], "standard", "white")
		}
	} else if len(os.Args) >= 3 {
		//useColor := flag.Bool("color", false, "display colorized output")
		colorname := flag.String("color", "", "coloring")
		flag.Parse()
		if *colorname != "" {
			colorvalue := *colorname
			if len(os.Args) <= 3 {
				asciiART.PrintART("",os.Args[2], "standard", colorvalue)
			} else {
				asciiART.PrintART(os.Args[2],os.Args[3], "standard", colorvalue)
			}
		} else {
			asciiART.PrintART("",os.Args[1], os.Args[2], "white")
		}
	}
}
