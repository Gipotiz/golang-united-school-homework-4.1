package reverse_string

import (
	"fmt"
	"reflect"
	"testing"
)

func TestReverseString(t *testing.T) {
	v := ReverseString("Hello world!")

	if ok := reflect.DeepEqual(v, "!dlrow olleH"); !ok {
		t.Error(v)
	}
	fmt.Println(v)
}
