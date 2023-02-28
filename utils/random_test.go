package utils

import (
	"testing"

	"github.com/qingbing/go-helpers/debug"
)

func Test_RandInt(t *testing.T) {
	cs := []debug.Case{
		{
			Name: "随机10以内",
			Args: map[string]any{
				"max": 10,
			},
			Want: map[string]any{
				"min": 0,
				"max": 10,
			},
		},
		{
			Name: "随机100以内",
			Args: map[string]any{
				"max": 100,
			},
			Want: map[string]any{
				"min": 0,
				"max": 100,
			},
		},
	}
	for _, c := range cs {
		t.Run(c.Name, func(t *testing.T) {
			res := RandInt(c.Args["max"].(int))
			min := c.Want["min"].(int)
			max := c.Want["max"].(int)
			if res < min || res >= max {
				t.Error(c.Errorf(res))
			}
		})
	}
}

func Test_Random(t *testing.T) {
	cs := []debug.Case{
		{
			Name: "随机2,5之间的整数",
			Args: map[string]any{
				"min": 2,
				"max": 5,
			},
			Want: map[string]any{
				"min": 2,
				"max": 5,
			},
		},
		{
			Name: "随机2,5之间的整数",
			Args: map[string]any{
				"min": 5,
				"max": 2,
			},
			Want: map[string]any{
				"min": 2,
				"max": 5,
			},
		},
	}
	for _, c := range cs {
		t.Run(c.Name, func(t *testing.T) {
			res := Random(c.Args["min"].(int), c.Args["max"].(int))
			min := c.Want["min"].(int)
			max := c.Want["max"].(int)
			if res < min || res >= max {
				t.Error(c.Errorf(res))
			}
		})
	}
}
