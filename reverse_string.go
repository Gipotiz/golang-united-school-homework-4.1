package reverse_string

func ReverseString(input string) (output string) {
	// solution goes here

	slice := make([]rune, len(input))

	for i, j := len(input)-1, 0; i >= 0; i, j = i-1, j+1 {
		slice[j] = rune(input[i])
	}
	output = string(slice)
	return
}
