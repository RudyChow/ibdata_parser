package body

// ListNode 通用链表基节点
type ListNode struct {
	PrevNodePageNumber uint32
	PrevNodeOffset     uint16
	NextNodePageNumber uint32
	NextNOdeOffset     uint16
}

// ListBaseNode 基节点
type ListBaseNode struct {
	ListLength          uint32
	FirstNodePageNumber uint32
	FirstNOdeOffset     uint16
	LastNodePageNumber  uint32
	LastNodeOffset      uint16
}

// SegmentHeader 段头部
type SegmentHeader struct {
	SpaceIDOfTheInodeEntry    uint32
	PageNumberOfTheInodeEntry uint32
	ByteOffsetOfTheInodeExtry uint32
}
