package types

import "fmt"

// 泛型切片
type List[T any] []T

// 列表长度
func (l *List[T]) Len() int {
	return len(*l)
}

// 列表末尾压入元素
func (l *List[T]) Push(els ...T) *List[T] {
	for _, el := range els {
		*l = append(*l, el)
	}
	return l
}

// 列表末尾弹出一个元素
func (l *List[T]) Pop() (el T) {
	len := l.Len()
	if len == 0 {
		return
	}
	el = (*l)[len-1]
	*l = (*l)[:(len - 1)]
	return
}

// 列表头部插入元素
func (l *List[T]) Unshift(els ...T) *List[T] {
	*l = append(els, (*l)...)
	return l
}

// 列表弹出一个元素
func (l *List[T]) Shift() (el T) {
	len := l.Len()
	if len == 0 {
		return
	}
	el = (*l)[0]
	*l = (*l)[1:]
	return
}

// 打印列表信息
func (l *List[T]) Print() {
	fmt.Println("Length:", (*l).Len())
	for i, v := range *l {
		fmt.Printf("Index: %d, Element: %#v\n", i, v)
	}
}
