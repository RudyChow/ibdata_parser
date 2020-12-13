package body

import (
	"encoding/binary"
)

// FSP_HDR 相关常量
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
	Size                uint32        // 当前表空间总的PAGE个数，扩展文件时需要更新该值 4B
	FreeLimit           uint32        // 当前尚未初始化的最小Page No。从该Page往后的都尚未加入到表空间的FREE LIST上 4B
	SpaceFlags          uint32        // 当前表空间的FLAG信息 4B
	FragNUsed           uint32        // FSP_FREE_FRAG链表上已被使用的Page数，用于快速计算该链表上可用空闲Page数 4B
	FreeList            *ListBaseNode // （直属于表空间）FREE链表的基节点，当一个Extent中所有page都未被使用时，放到该链表上，可以用于随后的分配 16B
	FreeFragList        *ListBaseNode // （直属于表空间）FREE_FRAG链表的基节点，通常这样的Extent中的Page可能归属于不同的segment，用于segment frag array page的分配 16B
	FullFragList        *ListBaseNode // （直属于表空间）FULL_FRAG链表的基节点。Extent中所有的page都被使用掉时，会放到该链表上，当有Page从该Extent释放时，则移回FREE_FRAG链表 16B
	NextUnusedSegmentID uint64        // 当前文件中最大Segment ID + 1，用于段分配时的seg id计数器 ID
	SegInodesFullList   *ListBaseNode // SEG_INODES_FULL链表的基节点，已被完全用满的Inode Page链表 16B
	SegInodesFreeList   *ListBaseNode // SEG_INODES_FREE链表的基节点，至少存在一个空闲Inode Entry的Inode Page被放到该链表上 16B
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
