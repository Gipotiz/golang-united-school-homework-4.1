package reverse_string

import (
	"fmt"
	"reflect"
	"testing"
)

func TestReverseString(t *testing.T) {
	//v := ReverseString("ÑµÐ²Ð¸Ð\u0080Ñ\u009FÐ")
	v := ReverseString("Hello world!")
	fmt.Println(v)

	if ok := reflect.DeepEqual(v, "!dlrow olleH"); !ok {
		t.Error(v)
	}
}
