package asciiART

import (
	"fmt"
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
		"white":   "\033[1;37m%s\033[0m",
	}
	if color != "white" {
		if replacestr == "" {
			for _, symbol := range message {
				sym := fmt.Sprintf(Colors[color], string(symbol))
				fmt.Print(sym)
			}
		} else {
			symWhite := ""
			sym := ""
			for _, symbol := range message {
				symWhite = symWhite + fmt.Sprintf(Colors["white"], string(symbol))
			}

			for _, rplsymbol := range replacestr {
				sym = sym + fmt.Sprintf(Colors[color], string(rplsymbol))
			}
			//filastr := (strings.ReplaceAll(sym,sym,symWhite))
			fmt.Print(sym+symWhite)
		}
	} else {
		fmt.Print("\033[1;37m\033[0m", message)
	}
}
