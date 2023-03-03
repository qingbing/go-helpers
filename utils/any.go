package utils

import (
	"reflect"
)

// 将任意数据转换成 []any 类型
func UnpackSliceAny(data any) []any {
	rt := reflect.TypeOf(data)
	switch rt.Kind() {
	case reflect.Slice:
	case reflect.Array:
	case reflect.Map:
		rv := reflect.ValueOf(data)
		rs := make([]any, rv.Len())
		mks := rv.MapKeys()
		for i := 0; i < rv.Len(); i++ {
			rs[i] = rv.MapIndex(mks[i]).Interface()
		}
		return rs
	default:
		return []any{data}
	}
	rv := reflect.ValueOf(data)
	rs := make([]any, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		rs[i] = rv.Index(i).Interface()
	}
	return rs
}
