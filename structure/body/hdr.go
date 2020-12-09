package body

// Hdr 系统表 page_type = 8
type Hdr struct {
	FileSpaceHeader *FileSpaceHeader // 112 byte
	XDESEntry       *HdrXDESEntry    // 10240 byte 256*XDES Entry
	// 5986 byte empty space
}

// FileSpaceHeader 表头
type FileSpaceHeader struct {
	SpaceID uint32 // 表空间id 4B
	// not used 4B
	Size                             uint32   // 当前表空间占有的页面数 4B
	FreeLimit                        uint32   // 尚未被初始化的最小页号，大于或等于这个页号的区对应的XDES Entry结构都没有被加入FREE链表 4B
	SpaceFlags                       uint32   // 表空间的一些占用存储空间比较小的属性 4B
	FragNUsed                        uint32   // FREE_FRAG链表中已使用的页面数量 4B
	ListBaseNodeForFreeList          [16]byte // FREE链表的基节点 16B
	ListBaseNodeForFreeFragList      [16]byte // FREE_FRAG链表的基节点 16B
	ListBaseNodeForFullFragList      [16]byte // FULL_FRAG链表的基节点 16B
	NextUnusedSegmentID              uint64   // 当前表空间中下一个未使用的 Segment ID
	ListBaseNodeForSegInodesFullList [16]byte // SEG_INODES_FULL链表的基节点 16B
	ListBaseNodeForSegInodesFreeList [16]byte // SEG_INODES_FREE链表的基节点 16B
}

// HdrXDESEntry 分区表
type HdrXDESEntry Xdes
