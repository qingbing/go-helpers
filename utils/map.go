package utils

import "reflect"

// 判断一个map中是否完全包含另一个map
func Contains[KT comparable, VT any](biggerMap, smallMap map[KT]VT) bool {
	if len(biggerMap) == 0 || len(smallMap) == 0 {
		return false
	}
	for sk, sv := range smallMap {
		if bv, ok := biggerMap[sk]; !ok || !reflect.DeepEqual(sv, bv) {
			return false
		}
	}
	return true
}
