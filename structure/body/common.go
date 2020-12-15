package body

import "encoding/binary"

// ListNode 通用链表基节点 12B
type ListNode struct {
	PrevNodePageNumber uint32
	PrevNodeOffset     uint16
	NextNodePageNumber uint32
	NextNOdeOffset     uint16
}

// InitListNode 通过byte初始化一个ListNode
func InitListNode(data []byte) *ListNode {
	return &ListNode{
		PrevNodePageNumber: binary.BigEndian.Uint32(data[0:4]),
		PrevNodeOffset:     binary.BigEndian.Uint16(data[4:6]),
		NextNodePageNumber: binary.BigEndian.Uint32(data[6:10]),
		NextNOdeOffset:     binary.BigEndian.Uint16(data[10:12]),
	}
}

// ListBaseNode 基节点 16B
type ListBaseNode struct {
	ListLength          uint32
	FirstNodePageNumber uint32
	FirstNodeOffset     uint16
	LastNodePageNumber  uint32
	LastNodeOffset      uint16
}

// InitListBaseNode 通过byte初始化一个basenode
func InitListBaseNode(data []byte) *ListBaseNode {
	return &ListBaseNode{
		ListLength:          binary.BigEndian.Uint32(data[0:4]),
		FirstNodePageNumber: binary.BigEndian.Uint32(data[4:8]),
		FirstNodeOffset:     binary.BigEndian.Uint16(data[8:10]),
		LastNodePageNumber:  binary.BigEndian.Uint32(data[10:14]),
		LastNodeOffset:      binary.BigEndian.Uint16(data[14:16]),
	}
}

// SegmentHeader 段头部信息 记录本页面所在段对应的INODE Entry位置信息 10B
type SegmentHeader struct {
	SpaceIDOfTheInodeEntry    uint32
	PageNumberOfTheInodeEntry uint32
	ByteOffsetOfTheInodeEntry uint16
}

// InitSegmentHeader 通过byte初始化一个SegmentHeader
func InitSegmentHeader(data []byte) *SegmentHeader {
	return &SegmentHeader{
		SpaceIDOfTheInodeEntry:    binary.BigEndian.Uint32(data[0:4]),
		PageNumberOfTheInodeEntry: binary.BigEndian.Uint32(data[4:8]),
		ByteOffsetOfTheInodeEntry: binary.BigEndian.Uint16(data[8:10]),
	}
}
