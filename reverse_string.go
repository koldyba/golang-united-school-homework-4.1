package reverse_string

func ReverseString(input string) (output string) {
	inputInRunes := []rune(input)
	outputInRunes := []rune{}
	for i := len(inputInRunes) - 1; i >= 0; i-- {
		outputInRunes = append(outputInRunes, inputInRunes[i])
	}
	return string(outputInRunes)
}
