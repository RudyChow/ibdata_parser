package body

// Xdes 每个区的第一个page page_type = 9
type Xdes struct {
	// 112 none
	XDESEntries [256]*XdesExtry // 10240 byte 256*XDES Entry
	// 5986 byte empty space
}

// XdesExtry XdesExtry
type XdesExtry struct {
	SegmentID       uint64
	ListNode        *ListNode
	State           uint32
	PageStateBitmap [16]byte
}
