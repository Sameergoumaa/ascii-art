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
			asciiART.PrintART(os.Args[1], "standard")
		}
	} else if len(os.Args) >= 3 {
		useColor := flag.Bool("color", false, "display colorized output")
		colorname := flag.String("colors", "White", "coloring")
		flag.Parse()

		if *useColor {

			colorvalue := *colorname
			asciiART.Colorize((colorvalue), "Hello, DigitalOcean!")
			asciiART.PrintART(os.Args[3], "standard")
		} else {
			asciiART.PrintART(os.Args[1], os.Args[2])
		}
	}
}
