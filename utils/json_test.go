package utils

import (
	"fmt"
	"log"
	"reflect"
	"testing"

	"github.com/qingbing/go-helpers/debug"
)

type tp struct {
	Name string `json:name`
	Age  int    `json:age`
	Say  func() `json:"-"` // 对于 json 不支持序列化的类型，必须要使用该标签进行序列化屏蔽，否则json化会失败
}

func Test_ToJson(t *testing.T) {
	cs := []debug.Case{
		{
			Name: "number-json",
			Args: map[string]any{
				"data": 10,
			},
			Want: map[string]any{
				"res": "10",
				"err": nil,
			},
		},
		{
			Name: "string-json",
			Args: map[string]any{
				"data": "hello",
			},
			Want: map[string]any{
				"res": "\"hello\"",
				"err": nil,
			},
		},
		{
			Name: "bolean-json",
			Args: map[string]any{
				"data": true,
			},
			Want: map[string]any{
				"res": "true",
				"err": nil,
			},
		},
	}
	for _, c := range cs {
		t.Run(c.Name, func(t *testing.T) {
			str, err := ToJson(c.Args["data"])
			res := map[string]any{
				"res": str,
				"err": err,
			}
			if !reflect.DeepEqual(res, c.Want) {
				fmt.Println(res)
				fmt.Println(c.Want)
				t.Error(c.Errorf(res))
			}
		})
	}
}

func Test_JsonToObj(t *testing.T) {
	tp1 := tp{
		Name: "qing",
		Age:  11,
		Say:  func() {},
	}
	str, err := ToJson(tp1)
	if err != nil {
		log.Fatalf("ToJson failed: %v", err)
	}
	fmt.Println(str)
	tp2 := tp{}
	err = JsonToObj(str, &tp2)
	if err != nil {
		log.Fatalf("ToJson failed: %v", err)
	}
	tp2.Age = 33
	fmt.Println(tp2)
}
