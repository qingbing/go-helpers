package utils

// 判断一个map中是否完全包含另一个map
func Contains(biggerMap, smallMap map[string]string) bool {
	if len(biggerMap) == 0 || len(smallMap) == 0 {
		return false
	}
	if len(biggerMap) < len(smallMap) {
		return false
	}
	for sk, sv := range smallMap {
		if bv, ok := biggerMap[sk]; !ok || sv != bv {
			return false
		}
	}
	return true
}
