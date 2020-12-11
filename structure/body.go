package structure

// Body 页中间
type Body interface {
	Unmarshal(data []byte)
}
