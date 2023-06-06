package asciiART

import "fmt"

func Colorize(color string, message string) {
	var Colors = map[string]string{
		"Black"  : "\u001b[30m",
		"Red"    : "\u001b[31m",
		"Green"  : "\u001b[32m",
		"Yellow" : "\u001b[33m",
		"Blue"   : "\u001b[34m",
		"Reset"  : "\u001b[0m",
	}
	fmt.Println(string(Colors[color]))
}
