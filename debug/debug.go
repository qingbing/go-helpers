package debug

import (
	"fmt"
	"runtime"
)

// 获取或返回调用处的调用栈
func CallerStack(isPrint bool) string {
	rStr := ""
	for i := 1; ; i++ {
		pc, file, line, ok := runtime.Caller(i)
		if !ok {
			break
		}
		rStr += fmt.Sprintf("%#v:%#v(%#v)\n", file, line, pc)
	}
	if isPrint {
		fmt.Println("\n===== Caller stack start*****")
		fmt.Println(rStr)
		fmt.Println("***** Caller stack end *****")
		return ""
	}
	return rStr
}
