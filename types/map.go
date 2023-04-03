package types

import "reflect"

// 泛型集合
type MapT[KT comparable, VT any] map[KT]VT

// 判断 mp 是否包含 sm 的所有元数
func (mp MapT[KT, KV]) Contains(sm MapT[KT, KV]) bool {
	if len(mp) == 0 || len(sm) == 0 {
		return false
	}
	for sk, sv := range sm {
		if bv, ok := mp[sk]; !ok || !reflect.DeepEqual(sv, bv) {
			return false
		}
	}
	return true
}

// 判断 mp 的元素是否都在 bm 中
func (mp MapT[KT, KV]) BelongTo(bm MapT[KT, KV]) bool {
	if len(mp) == 0 || len(bm) == 0 {
		return false
	}
	for sk, sv := range mp {
		if bv, ok := bm[sk]; !ok || !reflect.DeepEqual(sv, bv) {
			return false
		}
	}
	return true
}
