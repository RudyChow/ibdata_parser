package body

// DataDictionary 数据字典表，表空间的offset=7页，page_type = 6
type DataDictionary struct {
	Header string // 52B
	// unused 4B
	SegmentHeader string // 10B
	// empty space 16272 B
}

// DataDictionaryHeader 数据字典header
type DataDictionaryHeader struct {
	MaxRowID                    uint64
	MaxTableID                  uint64
	MaxIndexID                  uint64
	MaxSpaceID                  uint32
	MixIDLow                    uint32
	RootOfSysTablesClustIndex   uint32
	RootOfSysTableIDsClustIndex uint32
	RootOfSysColumnsClustIndex  uint32
	RootOfSysIndexesClustIndex  uint32
	RootOfSysFieldsClustIndex   uint32
}
