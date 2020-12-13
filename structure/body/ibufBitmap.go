package body

// IBUFBITMAP 相关常量
const (
	PerPageBits     = 4
	IbufBitmapPages = 16384 // 256M / 16K = 16384 刚好记录了一个组的页的记录
	IbufBitmapSize  = IbufBitmapPages * PerPageBits / 8
)

// IbufBitmap insert buffer bitmap
type IbufBitmap struct {
	Bitmap [IbufBitmapSize]byte // 8192B
	// empty space 8146B
}

// Unmarshal 解析
func (ibufBitmap *IbufBitmap) Unmarshal(data []byte) {
	var bitmap [IbufBitmapSize]byte
	copy(bitmap[:], data[0:IbufBitmapSize])
	ibufBitmap.Bitmap = bitmap
}

// NewIbufBitmap 创建一个bitmap
func NewIbufBitmap() *IbufBitmap {
	return &IbufBitmap{}
}
