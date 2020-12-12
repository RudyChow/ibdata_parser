package body

import "encoding/binary"

const (
	InodeEntriesLen    = 85
	FragmentEntrySize  = 192
	FragmentEntriesLen = 32
)

// Inode page_type = 3 段的索引信息
type Inode struct {
	ListNode     *ListNode                    // 12B 通用链表节点
	InodeEntries [InodeEntriesLen]*InodeEntry // 16320B 段描述信息 85个INODE Entry
	// 6B 尚未使用空间
}

// Unmarshal 解析
func (inode *Inode) Unmarshal(data []byte) {
	inode.ListNode = InitListNode(data[0:12])
	for i, begin := 0, 12; i < InodeEntriesLen; i++ {
		inode.InodeEntries[i] = InitInodeEntry(data[begin+i*FragmentEntrySize : begin+(i+1)*FragmentEntrySize])
	}
}

// InodeEntry InodeEntry  192B
type InodeEntry struct {
	SegmentID                  uint64
	NotFullNUsed               uint32
	ListBaseNodeForFreeList    *ListBaseNode
	ListBaseNodeForNotFullList *ListBaseNode
	ListBaseNodeForFullList    *ListBaseNode
	MagincNumber               uint32                     // 97937874 表示innode entry已经初始化
	FragmentEntries            [FragmentEntriesLen]uint32 // 零散的页面号
}

// InitInodeEntry 初始化一个nodeEntry
func InitInodeEntry(data []byte) *InodeEntry {
	entry := &InodeEntry{
		SegmentID:                  binary.BigEndian.Uint64(data[0:8]),
		NotFullNUsed:               binary.BigEndian.Uint32(data[8:12]),
		ListBaseNodeForFreeList:    InitListBaseNode(data[12:28]),
		ListBaseNodeForNotFullList: InitListBaseNode(data[28:44]),
		ListBaseNodeForFullList:    InitListBaseNode(data[44:60]),
		MagincNumber:               binary.BigEndian.Uint32(data[60:64]),
	}

	for i, fragmentEntriesBegin := 0, 64; i < FragmentEntriesLen; i++ {
		entry.FragmentEntries[i] = binary.BigEndian.Uint32(data[fragmentEntriesBegin+i*4 : fragmentEntriesBegin+(i+1)*4])
	}

	return entry
}

// NewInode new一个Inode
func NewInode() *Inode {
	return &Inode{}
}
