package wrapper

type SortWrapper struct {
	Items  []any               `json:"items"`
	SortBy func(p, q any) bool `json:"-"`
}

// 重新 sort.Len() 方法
func (w SortWrapper) Len() int { return len(w.Items) }

// 重新 sort.Swap() 方法
func (w SortWrapper) Swap(i, j int) { w.Items[i], w.Items[j] = w.Items[j], w.Items[i] }

// 重新 sort.Less() 方法
func (w SortWrapper) Less(i, j int) bool { return w.SortBy(w.Items[i], w.Items[j]) }
