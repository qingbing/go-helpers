package types

import (
	"fmt"
	"testing"
)

func Test_Contains(t *testing.T) {
	mp1 := MapT[string, string]{
		"name": "bing",
		"sex":  "nan",
		"age":  "18",
	}
	mp2 := MapT[string, string]{
		"name": "bing",
		"sex":  "nan",
	}
	fmt.Println(mp1.Contains(mp2))
	fmt.Println(mp2.Contains(mp1))
}

func Test_BelongTo(t *testing.T) {
	mp1 := MapT[string, string]{
		"name": "bing",
		"sex":  "nan",
		"age":  "18",
	}
	mp2 := MapT[string, string]{
		"name": "bing",
		"sex":  "nan",
	}
	fmt.Println(mp1.BelongTo(mp2))
	fmt.Println(mp2.BelongTo(mp1))
}
