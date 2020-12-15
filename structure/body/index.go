package body

import (
	"encoding/binary"
	"encoding/json"
)

// 记录类型
const (
	RecordTypeNormal = iota
	RecordTypeNonLeaf
	RecordTypeMin
	RecordTypeMax

	RecordFlagDeleteMask uint8  = 1 << 5
	RecordFlagMinRecMask uint8  = 1 << 4
	RecordFlagNOwned     uint8  = 0xff
	RecordFlagHeapNo     uint16 = 0xfff8
	RecordFlagType       uint16 = 0x7
)

// RecordTypeDict 数据类型字典
var RecordTypeDict = map[uint8]string{
	RecordTypeNormal:  "normal",
	RecordTypeNonLeaf: "NonLeaf",
	RecordTypeMin:     "MinRecord",
	RecordTypeMax:     "MaxRecord",
}

// Index 相关常量
const (
	IndexHeaderSize   = 56
	DefaultRecordSize = 13
)

// 页方向
const (
	_ = iota
	PageDirectionLeft
	PageDirectionRight
	PageDirectionSameRec
	PageDirectionSamePage
	PageDirectionNoDirection
)

// PageDictionDict 方向字典
var PageDictionDict = map[uint8]string{
	PageDirectionLeft:        "left",
	PageDirectionRight:       "right",
	PageDirectionSameRec:     "same rec",
	PageDirectionSamePage:    "same page",
	PageDirectionNoDirection: "no direction",
}

// Index 索引页 同时也是数据页
type Index struct {
	Header   *IndexHeader         // 56B
	Infumum  *DefaultMinMaxRecord // 13B infimum
	Supremum *DefaultMinMaxRecord // 13B supremum
}

// Unmarshal 解析
func (index *Index) Unmarshal(data []byte) {
	index.Header.Unmarshal(data[0:IndexHeaderSize])
	index.Infumum.Unmarshal(data[56:69])
	index.Supremum.Unmarshal(data[69:82])
}

// IndexHeader 索引页头 56B
type IndexHeader struct {
	NDirSlots  uint16         // 2B 在页目录中的槽数量
	HeapTop    uint16         // 2B 还未使用的空间最小地址，也就是说从该地址之后就是Free Space
	NHeap      uint16         // 2B 本页中的记录的数量（包括最小和最大记录以及标记为删除的记录）
	Free       uint16         // 2B 第一个已经标记为删除的记录地址（各个已删除的记录通过next_record也会组成一个单链表，这个单链表中的记录可以被重新利用）
	Garbage    uint16         // 2B 已删除记录占用的字节数
	LastInsert uint16         // 2B 最后插入记录的位置
	Dircetion  uint16         // 2B 记录插入的方向
	NDircetion uint16         // 2B 一个方向连续插入的记录数量
	NRecs      uint16         // 2B 该页中记录的数量（不包括最小和最大记录以及被标记为删除的记录）
	MaxTrxID   uint64         // 8B 修改当前页的最大事务ID，该值仅在二级索引中定义
	Level      uint16         // 2B 当前页在B+树中所处的层级
	IndexID    uint64         // 8B 索引ID，表示当前页属于哪个索引
	BtrSegLeaf *SegmentHeader // 10B B+树叶子段的头部信息，仅在B+树的Root页定义
	BtrSegTop  *SegmentHeader // 10B B+树非叶子段的头部信息，仅在B+树的Root页定义
}

// Unmarshal 解析
func (indexHeader *IndexHeader) Unmarshal(data []byte) {
	indexHeader.NDirSlots = binary.BigEndian.Uint16(data[0:2])
	indexHeader.HeapTop = binary.BigEndian.Uint16(data[2:4])
	indexHeader.NHeap = binary.BigEndian.Uint16(data[4:6])
	indexHeader.Free = binary.BigEndian.Uint16(data[6:8])
	indexHeader.Garbage = binary.BigEndian.Uint16(data[8:10])
	indexHeader.LastInsert = binary.BigEndian.Uint16(data[10:12])
	indexHeader.Dircetion = binary.BigEndian.Uint16(data[12:14])
	indexHeader.NDircetion = binary.BigEndian.Uint16(data[14:16])
	indexHeader.NRecs = binary.BigEndian.Uint16(data[16:18])
	indexHeader.MaxTrxID = binary.BigEndian.Uint64(data[18:26])
	indexHeader.Level = binary.BigEndian.Uint16(data[26:28])
	indexHeader.IndexID = binary.BigEndian.Uint64(data[28:36])
	indexHeader.BtrSegLeaf = InitSegmentHeader(data[36:46])
	indexHeader.BtrSegTop = InitSegmentHeader(data[46:56])
}

// RecordInfo 每个记录头一定会有的信息
type RecordInfo struct { // 5B
	FirstFlags  [1]byte // 1B -- 1bit 预留位1  1bit 预留位2  1bit 标记该记录是否被删除  1bit B+树的每层非叶子节点中的最小记录都会添加该标记  4bits 表示当前记录拥有的记录数
	SecondFlags uint16  // 2B -- 13bit 表示当前记录在记录堆的位置信息   3bit 表示当前记录的类型，0表示普通记录，1表示B+树非叶节点记录，2表示最小记录，3表示最大记录
	NextRecord  uint16  // 2B 表示下一条记录的相对位置
}

// Unmarshal 解析
func (recordInfo *RecordInfo) Unmarshal(data []byte) {
	first := [1]byte{}
	copy(first[:], data[0:1])
	recordInfo.FirstFlags = first
	recordInfo.SecondFlags = binary.BigEndian.Uint16(data[1:3])
	recordInfo.NextRecord = binary.BigEndian.Uint16(data[3:5])
}

// MarshalJSON 序列化成json
func (recordInfo *RecordInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(
		struct {
			DeleteMask bool
			MinRecMask bool
			NOwned     uint8 // 页内会进行分组，然后进行二分查找，这个相当于组内共有几条记录
			HeapNo     uint16
			RecordType string
			NextRecord uint16 // 相当于指针，组成链表
		}{
			DeleteMask: (recordInfo.FirstFlags[0] | RecordFlagDeleteMask) == RecordFlagDeleteMask,
			MinRecMask: (recordInfo.FirstFlags[0] | RecordFlagMinRecMask) == RecordFlagMinRecMask,
			NOwned:     recordInfo.FirstFlags[0] & RecordFlagNOwned,
			HeapNo:     (recordInfo.SecondFlags & RecordFlagHeapNo) >> 3,
			RecordType: RecordTypeDict[uint8(recordInfo.SecondFlags&RecordFlagType)],
			NextRecord: recordInfo.NextRecord,
		})
}

// DefaultMinMaxRecord 默认会有的最小和最大的记录 13B
type DefaultMinMaxRecord struct {
	RecordInfo *RecordInfo // 5B 记录的信息
	DType      string      `json:"Type"` //infimum 或者 supremum
}

// Unmarshal 解析
func (defaultMinMaxRecord *DefaultMinMaxRecord) Unmarshal(data []byte) {
	defaultMinMaxRecord.RecordInfo.Unmarshal(data[0:5])
	defaultMinMaxRecord.DType = string(data[5:13])
}

// NewIndex 创建一个索引页
func NewIndex() *Index {
	return &Index{
		Header: &IndexHeader{},
		Infumum: &DefaultMinMaxRecord{
			&RecordInfo{},
			"",
		},
		Supremum: &DefaultMinMaxRecord{
			&RecordInfo{},
			"",
		},
	}
}
