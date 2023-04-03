package errors

import (
	"fmt"
	"time"
)

// 定制化错误消息
type Error struct {
	Code string
	Msg  string
	Time time.Time
}

// 实现 builtin.Error 接口
func (err Error) Error() string {
	return fmt.Sprintf("Code:%s;Message:%s;Time:%s", err.Code, err.Msg, err.Time.Format("2006-01-02 15:04:05"))
}

// 实例化一个定制化错误
func New(code string, msg string) Error {
	return Error{
		Code: code,
		Msg:  msg,
		Time: time.Now(),
	}
}
