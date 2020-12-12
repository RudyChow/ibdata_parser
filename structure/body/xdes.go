package body

import "encoding/binary"

// XDESEntriesSize entries总大小
const (
	XDESEntriedLen  = 256
	XDESEntrySize   = 40
	XDESEntriesSize = XDESEntriedLen * XDESEntrySize
)

// Xdes 每个区的第一个page page_type = 9
type Xdes struct {
	// 112 none
	XDESEntries [XDESEntriedLen]*XdesEntry // 10240 byte 256*XDES Entry
	// 5986 byte empty space
}

// Unmarshal 解析
func (xdes *Xdes) Unmarshal(data []byte) {
	// 解析entries
	XDESEntriesBytes := data[FileSpaceHeaderSize : FileSpaceHeaderSize+XDESEntriesSize]
	for i := 0; i < XDESEntriedLen; i++ {
		xdes.XDESEntries[i] = InitXdesEntry(XDESEntriesBytes[i*XDESEntrySize : i*XDESEntrySize+XDESEntrySize])
	}
}

// XdesEntry XdesEntry 40B
type XdesEntry struct {
	SegmentID       uint64    // 8B
	ListNode        *ListNode // 12B
	State           uint32    // 4B FREE空闲的区、FREE_FRAG有剩余空间的碎片区、FULL_FRAG没有剩余空间的碎片区和FSEG附属于某个段的区
	PageStateBitmap [16]byte  // 16B 一共128个bit，一个区有64页，所以每两个bit表示一个页的信息，第一个bit表示页是否空闲，第二个bit则暂无用处
}

// InitXdesEntry 通过byte初始化一个basenode
func InitXdesEntry(data []byte) *XdesEntry {
	var bitmap [16]byte
	copy(bitmap[:], data[24:40])
	return &XdesEntry{
		SegmentID:       binary.BigEndian.Uint64(data[0:8]),
		ListNode:        InitListNode(data[8:20]),
		State:           binary.BigEndian.Uint32(data[20:24]),
		PageStateBitmap: bitmap,
	}
}

// NewXdes new一个xdes
func NewXdes() *Xdes {
	return &Xdes{
		XDESEntries: [XDESEntriedLen]*XdesEntry{},
	}
}
