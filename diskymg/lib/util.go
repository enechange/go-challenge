package lib

// 受け取った値をポインタ型にして返す
func Ptr[T any](v T) *T {
	return &v
}
