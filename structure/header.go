package structure

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
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

type jsonFileHeader FileHeader

// MarshalJSON 序列化成json
func (fileHeader FileHeader) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		jsonFileHeader
		TypeName string
	}{
		jsonFileHeader: jsonFileHeader(fileHeader),
		TypeName:       fileHeader.GetTypeName()})
}

// GetTypeName 获取类型名称
func (fileHeader *FileHeader) GetTypeName() string {
	name, ok := PageTypeNameDict[fileHeader.Ptype]
	if ok {
		return name
	}
	return fmt.Sprintf("Unknown Type %d", fileHeader.Ptype)

}
