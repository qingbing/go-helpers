package utils

import (
	"math/rand"
	"time"
)

var rander *rand.Rand

// 启动初始化时设置随机种子
func init() {
	rander = rand.New(rand.NewSource(time.Now().UnixNano()))
}

// 对外暴露 rander
func NewRander() *rand.Rand {
	return rander
}

// 获取随机整数,最大不超过 max
func RandInt(max int) int {
	return rander.Intn(max)
}

// 获取随机整数, 范围在 min 和 max 之间, mix, max 不区分大小, 函数内区分
func Random(min, max int) int {
	if min == max {
		return min
	}
	if max < min {
		min, max = max, min
	}
	return rander.Intn(max-min) + min
}
