package utils

import (
	"reflect"
	"testing"

	"github.com/qingbing/go-helpers/debug"
)

func Test_UnpackSliceAny(t *testing.T) {
	cs := []debug.Case{
		{
			Name: "string",
			Args: map[string]any{
				"data": "string",
			},
			Want: map[string]any{
				"res": []any{"string"},
			},
		},
		{
			Name: "numberic",
			Args: map[string]any{
				"data": 1,
			},
			Want: map[string]any{
				"res": []any{1},
			},
		},
		{
			Name: "boolean",
			Args: map[string]any{
				"data": false,
			},
			Want: map[string]any{
				"res": []any{false},
			},
		},
		{
			Name: "array",
			Args: map[string]any{
				"data": [2]int{11, 22},
			},
			Want: map[string]any{
				"res": []any{11, 22},
			},
		},
		{
			Name: "slice",
			Args: map[string]any{
				"data": []int{11, 22},
			},
			Want: map[string]any{
				"res": []any{11, 22},
			},
		},
		{
			Name: "slice",
			Args: map[string]any{
				"data": map[string]string{"name": "name1", "sex": "sex1"},
			},
			Want: map[string]any{
				"res": []any{"name1", "sex1"},
			},
		},
	}
	for _, c := range cs {
		t.Run(c.Name, func(t *testing.T) {
			res := UnpackSliceAny(c.Args["data"])
			if !reflect.DeepEqual(res, c.Want["res"]) {
				t.Error(c.Errorf(res))
			}
		})
	}
}
