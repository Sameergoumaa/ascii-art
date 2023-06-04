package asciiART

func MapART(char rune) int {
	return 9*(int(char)-32) + 2
	//return 9*(int(char)-40+8) + 2
}
