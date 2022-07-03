package reverse_string

func ReverseString(input string) (output string) {
	// solution goes here

	for i := len(input) - 1; i >= 0; i-- {
		output += string(input[i])
	}

	return
}
