package debug

import "fmt"

// 统一测试用例结构
type Case struct {
	Name string
	Args map[string]any
	Want map[string]any
}

// 统一测试错误输出
func (c *Case) Errorf(res any) string {
	return fmt.Sprintf("Fail Case: %s, args: %+v, Want: %+v, Result: %v\n", c.Name, c.Args, c.Want, res)
}
