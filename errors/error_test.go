package errors

import (
	"fmt"
	"testing"
)

func TestXxx(t *testing.T) {
	err := New("500", "错误")
	fmt.Println(err)
}
