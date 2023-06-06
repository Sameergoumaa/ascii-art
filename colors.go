package asciiART

import (
	"fmt"
	"strings"
)

func Colorize(replacestr string, color string, message string) {
	var Colors = map[string]string{
		"black":   "\033[1;30m%s\033[0m",
		"red":     "\033[1;31m%s\033[0m",
		"green":   "\033[1;32m%s\033[0m",
		"yellow":  "\033[1;33m%s\033[0m",
		"blue":    "\033[1;34m%s\033[0m",
		"magenta": "\033[1;35m%s\033[0m",
		"cyan":    "\033[1;36m%s\033[0m",
		//"white":   "\033[1;37m%s\033[0m",
		"orange":  "\033[1;38;5;208m%s\033[0m",
	}
	if color != "white" {
		if replacestr == "" {
			for _, symbol := range message {
				fmt.Printf(Colors[color], string(symbol))
			}
		} else {
			for _, symbol := range message {
				fmt.Printf(Colors["white"], string(symbol))
			}
			fmt.Printf(Colors["blue"],strings.ReplaceAll(message, replacestr, message))
		}
	} else {
		fmt.Print((Colors[color]), message)
	}
}
