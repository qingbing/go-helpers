package wrapper

import (
	"fmt"
	"sort"
	"testing"

	"github.com/qingbing/go-helpers/utils"
)

type sortItem struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type sortItemWrapper struct {
	SortWrapper
}

func Test_SortWrapper(t *testing.T) {
	jStr := ""
	w := &sortItemWrapper{
		SortWrapper: SortWrapper{
			SortBy: func(p, q any) bool {
				pi := p.(sortItem)
				qi := q.(sortItem)
				return pi.Name < qi.Name
			},
		},
	}
	w.Items = append(w.Items, sortItem{Name: "qingbing", Age: 19})
	w.Items = append(w.Items, sortItem{Name: "nai", Age: 19})
	w.Items = append(w.Items, sortItem{Name: "yi", Age: 19})
	w.Items = append(w.Items, sortItem{Name: "cheng", Age: 19})

	jStr, _ = utils.ToJson(w)
	fmt.Println(jStr)

	sort.Sort(w) // 排序

	jStr, _ = utils.ToJson(w)
	fmt.Println(jStr)

	x := sort.Reverse(w) // 倒叙

	jStr, _ = utils.ToJson(x)
	fmt.Println(jStr)
}
