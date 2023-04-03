package utils

import (
	"fmt"
	"testing"
)

func Test_Contains(t *testing.T) {
	mp1 := map[string]string{
		"name": "bing",
		"sex":  "nan",
		"age":  "18",
	}
	mp2 := map[string]string{
		"name": "bing",
		"sex":  "nan",
	}

	fmt.Println(Contains[string, string](mp1, mp2))
	fmt.Println(Contains(mp1, mp2))

	mp11 := map[string]int{
		"name": 10,
		"sex":  11,
		"age":  11,
	}
	mp12 := map[string]int{
		"name": 10,
		"sex":  11,
	}
	fmt.Println(Contains[string, int](mp11, mp12))
	fmt.Println(Contains(mp11, mp12))
}
