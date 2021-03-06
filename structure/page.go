package structure

import (
	"encoding/binary"
	"errors"
	"ibdata_parser/structure/body"
)

// page info
const (
	PageSize = 16 * 1024

	PageHeaderSize  = 38
	PageBodySize    = PageSize - PageHeaderSize - PageTrailerSize
	PageTrailerSize = 8

	PageHeaderOffset  = 0
	PageBodyOffset    = PageHeaderOffset + PageHeaderSize
	PageTrailerOffset = PageBodyOffset + PageBodySize

	PageSysDataDictOffset = 7
)

// Page innodb页 page = header(38) + body + trailer(8)
type Page struct {
	FileHeader  *FileHeader
	Body        Body
	FileTrailer FileTrailer
}

// Unmarshal 解析整体
func (page *Page) Unmarshal(data []byte) {
	// 头部
	page.UnmarshalHeader(data[PageHeaderOffset:PageHeaderSize])
	// body
	page.UnmarshalTrailer(data[PageBodyOffset : PageBodyOffset+PageBodySize])
	// 尾部
	page.UnmarshalTrailer(data[PageTrailerOffset : PageTrailerOffset+PageTrailerSize])
}

// UnmarshalHeader 解析头部
func (page *Page) UnmarshalHeader(data []byte) {
	page.FileHeader.unmarshal(data)
}

// UnmarshalBody 解析body
func (page *Page) UnmarshalBody(data []byte) error {
	switch page.FileHeader.Ptype {
	case FilPageTypeFspHdr: // FSP_HDR
		page.Body = body.NewHdr()
	case FilPageTypeXdes: // XDES
		page.Body = body.NewXdes()
	case FilPageInode: // INODE
		page.Body = body.NewInode()
	case FilPageTypeSys: // SYS
		// data dict
		if page.FileHeader.Offset == PageSysDataDictOffset {
			page.Body = body.NewDataDictionary()
		} else {
			return errors.New("unfinish feature")
		}
	case FilPageIbufBitmap: // IBUFBITMAP
		page.Body = body.NewIbufBitmap()
	case FilPageTypeIndex: // INDEX
		page.Body = body.NewIndex()
	default:
		return errors.New("unfinish feature")

	}

	page.Body.Unmarshal(data)
	return nil
}

// UnmarshalTrailer 解析尾部
func (page *Page) UnmarshalTrailer(data []byte) {
	page.FileTrailer = (FileTrailer)(binary.BigEndian.Uint64(data[0:8]))
}

// NewPage 创建一个页
func NewPage() *Page {
	return &Page{
		FileHeader:  &FileHeader{},
		Body:        nil,
		FileTrailer: 0,
	}
}
