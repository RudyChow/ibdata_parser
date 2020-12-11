package body

import (
	"encoding/binary"
)

const (
	FileSpaceHeaderSize = 112
)

// Hdr 系统表 page_type = 8
// FSP_HDR 和 XDES 唯一的区别就是多了一个 FSP Header
type Hdr struct {
	FileSpaceHeader *FileSpaceHeader // 112 byte
	Xdes            *Xdes            // 10240 byte 256*XDES Entry
	// 5986 byte empty space
}

// Unmarshal 解析
func (hdr *Hdr) Unmarshal(data []byte) {
	// 解析头
	hdr.FileSpaceHeader.unmarshal(data[0:FileSpaceHeaderSize])
	// 解析xdes
	hdr.Xdes.Unmarshal(data)
}

// FileSpaceHeader 表头
type FileSpaceHeader struct {
	SpaceID uint32 // 表空间id 4B
	// not used 4B
	Size                uint32        // 当前表空间占有的页面数 4B
	FreeLimit           uint32        // 尚未被初始化的最小页号，大于或等于这个页号的区对应的XDES Entry结构都没有被加入FREE链表 4B
	SpaceFlags          uint32        // 表空间的一些占用存储空间比较小的属性 4B
	FragNUsed           uint32        // FREE_FRAG链表中已使用的页面数量 4B
	FreeList            *ListBaseNode // FREE链表的基节点 16B
	FreeFragList        *ListBaseNode // FREE_FRAG链表的基节点 16B
	FullFragList        *ListBaseNode // FULL_FRAG链表的基节点 16B
	NextUnusedSegmentID uint64        // 当前表空间中下一个未使用的 Segment ID
	SegInodesFullList   *ListBaseNode // SEG_INODES_FULL链表的基节点 16B
	SegInodesFreeList   *ListBaseNode // SEG_INODES_FREE链表的基节点 16B
}

// unmarshal 解析 FSP Header
func (fileSpaceHeader *FileSpaceHeader) unmarshal(data []byte) {
	fileSpaceHeader.SpaceID = binary.BigEndian.Uint32(data[0:4]) // 4B
	// not used 4B 4:8
	fileSpaceHeader.Size = binary.BigEndian.Uint32(data[8:12])                 // 4B
	fileSpaceHeader.FreeLimit = binary.BigEndian.Uint32(data[12:16])           // 4B
	fileSpaceHeader.SpaceFlags = binary.BigEndian.Uint32(data[16:20])          // 4B
	fileSpaceHeader.FragNUsed = binary.BigEndian.Uint32(data[20:24])           // 4B
	fileSpaceHeader.FreeList = InitListBaseNode(data[24:40])                   // 16B
	fileSpaceHeader.FreeFragList = InitListBaseNode(data[40:56])               // 16B
	fileSpaceHeader.FullFragList = InitListBaseNode(data[56:72])               // 16B
	fileSpaceHeader.NextUnusedSegmentID = binary.BigEndian.Uint64(data[72:80]) // 8B
	fileSpaceHeader.SegInodesFullList = InitListBaseNode(data[80:96])          // 16B
	fileSpaceHeader.SegInodesFreeList = InitListBaseNode(data[96:112])         // 16B

}

// NewHdr new一个hdr
func NewHdr() *Hdr {
	return &Hdr{
		FileSpaceHeader: &FileSpaceHeader{},
		Xdes:            NewXdes(),
	}
}
