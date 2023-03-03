package utils

import "encoding/json"

// 任意数据转换成成json字符串
func ToJson(data any) (string, error) {
	bs, err := json.Marshal(data)
	if err != nil {
		return "", err
	} else {
		return string(bs), nil
	}
}

// json 字符串转换成结构体，数组，map等,
// 对于传递的参数 objType, 必须传递指针进来
func JsonToObj(data string, objType any) error {
	return json.Unmarshal([]byte(data), &objType)
}
