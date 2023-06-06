package asciiART

func MapART(char rune) int {
	return 9*(int(char)-32) + 2
}
