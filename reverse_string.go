package reverse_string

func ReverseString(input string) (output string) {
	// solution goes here

	arr := []rune(input)
	var rev_arr []rune

	for i := len(input) - 1; i >= 0; i-- {
		rev_arr = append(rev_arr, arr[i])
	}
	output = string(rev_arr)

	return
}
