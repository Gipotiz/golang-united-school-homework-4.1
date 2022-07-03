package reverse_string

func ReverseString(input string) (output string) {
	// solution goes here

	arr := []rune(input)
	var revArr []rune

	for i := len(arr) - 1; i >= 0; i-- {
		revArr = append(revArr, arr[i])
	}
	output = string(revArr)

	return
}
