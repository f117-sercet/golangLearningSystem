package http_server

type ByteView struct {
	b []byte
}

// Len 计算长度
func (v ByteView) Len() int {
	return len(v.b)
}

// ByteSlice 复制并返回切片
func (v ByteView) ByteSlice() []byte {
	return cloneBytes(v.b)
}

func (v ByteView) String() string {
	return string(v.b)
}

func cloneBytes(b []byte) []byte {
	c := make([]byte, len(b))
	copy(c, b)
	return c
}
