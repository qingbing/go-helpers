package types

import (
	"fmt"
	"testing"
)

func Test_List(t *testing.T) {
	lister := List[string]{}
	// Push
	lister.Push("test", "bing").
		Push("qing")
	lister.Print()

	// Pop
	el := lister.Pop()
	fmt.Println(el)
	lister.Print()

	// Shift
	el = lister.Shift()
	fmt.Println(el)
	el = lister.Shift()
	fmt.Println(el)
	el = lister.Pop()
	fmt.Println(el)
	lister.Print()

	// Unshift
	lister.Unshift("test").
		Unshift("bing", "qing")
	lister.Print()
}
