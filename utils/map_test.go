package utils

import (
	"testing"

	"github.com/qingbing/go-helpers/debug"
)

func Test_Contains(t *testing.T) {
	cs := []debug.Case{
		{
			Name: "all maps are nil",
			Args: map[string]any{
				"bmap": nil,
				"smap": nil,
			},
			Want: map[string]any{
				"res": false,
			},
		},
		{
			Name: "bigger map is nil",
			Args: map[string]any{
				"bmap": nil,
				"smap": map[string]string{"name": "qingbing", "sex": "nan"},
			},
			Want: map[string]any{
				"res": false,
			},
		},
		{
			Name: "smaller map is nil",
			Args: map[string]any{
				"bmap": map[string]string{"name": "qingbing", "sex": "nan"},
				"smap": nil,
			},
			Want: map[string]any{
				"res": false,
			},
		},
		{
			Name: "all maps's length is zero",
			Args: map[string]any{
				"bmap": map[string]string{},
				"smap": map[string]string{},
			},
			Want: map[string]any{
				"res": false,
			},
		},
		{
			Name: "bigger map's length is zero",
			Args: map[string]any{
				"bmap": map[string]string{},
				"smap": map[string]string{"name": "qingbing", "sex": "nan"},
			},
			Want: map[string]any{
				"res": false,
			},
		},
		{
			Name: "smaller map's length is zero",
			Args: map[string]any{
				"bmap": map[string]string{},
				"smap": map[string]string{"name": "qingbing", "sex": "nan"},
			},
			Want: map[string]any{
				"res": false,
			},
		},
		{
			Name: "bigger map does not contain all",
			Args: map[string]any{
				"bmap": map[string]string{"name": "qingbing"},
				"smap": map[string]string{"name": "qingbing", "sex": "nan"},
			},
			Want: map[string]any{
				"res": false,
			},
		},
		{
			Name: "bigger map is equal to smaller map",
			Args: map[string]any{
				"bmap": map[string]string{"name": "qingbing", "sex": "nan"},
				"smap": map[string]string{"name": "qingbing", "sex": "nan"},
			},
			Want: map[string]any{
				"res": true,
			},
		},
		{
			Name: "bigger map is bigger than smaller map",
			Args: map[string]any{
				"bmap": map[string]string{"name": "qingbing", "sex": "nan"},
				"smap": map[string]string{"name": "qingbing"},
			},
			Want: map[string]any{
				"res": true,
			},
		},
	}
	for _, c := range cs {
		t.Run(c.Name, func(t *testing.T) {
			bmap, ok := c.Args["bmap"].(map[string]string)
			if !ok {
				bmap = nil
			}
			smap, ok := c.Args["smap"].(map[string]string)
			if !ok {
				smap = nil
			}
			res := Contains(bmap, smap)
			if res != c.Want["res"] {
				t.Error(c.Errorf(res))
			}
		})
	}
}
