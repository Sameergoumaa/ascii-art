package asciiART

func MapART(char byte) int {

	return 9*(int(char)-40+8) + 2
}
