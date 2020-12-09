package body

// Inode page_type = 3
type Inode struct {
	ListNode     *ListNode   // 12B 通用链表节点
	InodeExtries [16320]byte // 16320B 段描述信息 85个INODE Entry
	// 6B 尚未使用空间
}

// InodeExtry InodeExtry
type InodeExtry struct {
	SegmentID                  uint64
	NOtFullNUsed               uint32
	ListBaseNodeForFreeList    *ListBaseNode
	ListBaseNodeForNotFullList *ListBaseNode
	ListBaseNodeForFullList    *ListBaseNode
	MagincNumber               uint32     // 97937874 表示innode entry已经初始化
	FragmentEntries            [32]uint32 // 零散的页面号
}
