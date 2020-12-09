package structure

import (
	"encoding/binary"
	"fmt"
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
)

// page type
const (
	FilPageTypeAllocated = iota // 0
	_
	FilPageUndoLog                             // 2
	FilPageInode                               // 3
	FilPageIbufFreeList                        // 4
	FilPageIbufBitmap                          // 5
	FilPageTypeSys                             // 6
	FilPageTypeTrxSys                          // 7
	FilPageTypeFspHdr                          // 8
	FilPageTypeXdes                            // 9
	FilPageTypeBlob                            // 10
	FilPageTypeZBlob                           // 11
	FilPageTypeZBlob2                          // 12
	FilPageTypeUnknown                         // 13
	FilPageTypeCompressed                      // 14
	FilPageTypeEncrypted                       // 15
	FilPageTypeCompressedAndEncrypted          // 16
	FilPageTypeEncryptedRtree                  // 17
	FilPageTypeSdiBlob                         // 18
	FilPageTypeSdiZBlob                        // 19
	FilPageTypeLegacyDblwr                     // 20
	FilPageTypeResgArray                       // 21
	FilPageTypeLobIndex                        // 22
	FilPageTypeLobData                         // 23
	FilPageTypeLobFirst                        // 24
	FilPageTypeZBlobFirst                      // 25
	FilPageTypeZBlobData                       // 26
	FilPageTypeZBlobIndex                      // 27
	FilPageTypeZBlobFrag                       // 28
	FilPageTypeZBlobFragExtry                  // 29
	FilPageTypeSDI                    = 0x45BD // 17853
	FilPageTypeRTREE                  = 0x45BE // 17854
	FilPageTypeIndex                  = 0x45BF // 17855
)

// PageTypeNameDict 类型映射
var PageTypeNameDict = map[uint16]string{
	FilPageTypeAllocated:              "Freshly allocated page",
	FilPageUndoLog:                    "Undo log page",
	FilPageInode:                      "Index node",
	FilPageIbufFreeList:               "Insert buffer free list",
	FilPageIbufBitmap:                 "Insert buffer bitmap",
	FilPageTypeSys:                    "System page",
	FilPageTypeTrxSys:                 "Transaction system data",
	FilPageTypeFspHdr:                 "File space header",
	FilPageTypeXdes:                   "Extent descriptor page",
	FilPageTypeBlob:                   "Uncompressed BLOB page",
	FilPageTypeZBlob:                  "First compressed BLOB page",
	FilPageTypeZBlob2:                 "Subsequent compressed BLOB page",
	FilPageTypeUnknown:                "garbage",
	FilPageTypeCompressed:             "Compressed page",
	FilPageTypeEncrypted:              "Encrypted page",
	FilPageTypeCompressedAndEncrypted: "Compressed and Encrypted page",
	FilPageTypeEncryptedRtree:         "Encrypted R-tree page",
	FilPageTypeSdiBlob:                "Uncompressed SDI BLOB page",
	FilPageTypeSdiZBlob:               "Commpressed SDI BLOB page",
	FilPageTypeLegacyDblwr:            "Legacy doublewrite buffer page",
	FilPageTypeResgArray:              "Rollback Segment Array page",
	FilPageTypeLobIndex:               "Index pages of uncompressed LOB",
	FilPageTypeLobData:                "Data pages of uncompressed LOB",
	FilPageTypeLobFirst:               "The first page of an uncompressed LOB",
	FilPageTypeZBlobFirst:             "The first page of a compressed LOB",
	FilPageTypeZBlobData:              "Data pages of compressed LOB",
	FilPageTypeZBlobIndex:             "Index pages of compressed LOB",
	FilPageTypeZBlobFrag:              "Fragment pages of compressed LOB",
	FilPageTypeZBlobFragExtry:         "Index pages of fragment pages (compressed LOB)",
	FilPageTypeSDI:                    "Tablespace SDI Index page",
	FilPageTypeRTREE:                  "R-tree node",
	FilPageTypeIndex:                  "B-tree node",
}

// Page innodb页
type Page struct {
	FileHeader  *FileHeader
	Body        interface{}
	FileTrailer FileTrailer
}

// Unmarshal 解析整体
func (page *Page) Unmarshal(data []byte) {
	// 头部
	page.UnmarshalHeader(data[PageHeaderOffset:PageHeaderSize])
	// body @todo

	// 尾部
	page.UnmarshalTrailer(data[PageTrailerOffset : PageTrailerOffset+PageTrailerSize])
}

// UnmarshalHeader 解析头部
func (page *Page) UnmarshalHeader(data []byte) {
	page.FileHeader.unmarshal(data)
}

// UnmarshalBody 解析body
func (page *Page) UnmarshalBody(data []byte) {
}

// UnmarshalTrailer 解析尾部
func (page *Page) UnmarshalTrailer(data []byte) {
	page.FileTrailer = (FileTrailer)(binary.BigEndian.Uint64(data[0:8]))
}

// FileHeader 页头
type FileHeader struct {
	SpaceOrChksum      uint32 // 校验码
	Offset             uint32 // 页偏移量
	Perv               uint32 // 上一页
	Next               uint32 // 下一页
	Lsn                uint64 // 页面被最后修改时的lsn
	Ptype              uint16 // 页面类型
	FileFlushLsn       uint64 // 文件更新的lsn
	ArchLogNoOrSpaceID uint32 // 所属表空间id
}

// unmarshal 解析
func (fileHeader *FileHeader) unmarshal(data []byte) {

	fileHeader.SpaceOrChksum = binary.BigEndian.Uint32(data[0:4])
	fileHeader.Offset = binary.BigEndian.Uint32(data[4:8])
	fileHeader.Perv = binary.BigEndian.Uint32(data[8:12])
	fileHeader.Next = binary.BigEndian.Uint32(data[12:16])
	fileHeader.Lsn = binary.BigEndian.Uint64(data[16:24])
	fileHeader.Ptype = binary.BigEndian.Uint16(data[24:26])
	fileHeader.FileFlushLsn = binary.BigEndian.Uint64(data[26:34])
	fileHeader.ArchLogNoOrSpaceID = binary.BigEndian.Uint32(data[34:38])
}

// GetTypeName 获取类型名称
func (fileHeader *FileHeader) GetTypeName() string {
	name, ok := PageTypeNameDict[fileHeader.Ptype]
	if ok {
		return name
	}
	return fmt.Sprintf("Unknown Type %d", fileHeader.Ptype)

}

// FileTrailer 页尾
type FileTrailer uint64

// GetChksum 获取校验和
func (fileTrailer FileTrailer) GetChksum() uint32 {
	return (uint32)(fileTrailer >> 32)
}

// GetPageLsn 获取lsn
func (fileTrailer FileTrailer) GetPageLsn() uint32 {
	return (uint32)(fileTrailer)
}

// NewPage 创建一个页
func NewPage() *Page {
	return &Page{
		FileHeader:  &FileHeader{},
		Body:        nil,
		FileTrailer: 0,
	}
}
