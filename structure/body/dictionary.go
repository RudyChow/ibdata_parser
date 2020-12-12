package body

import "encoding/binary"

// DataDictionary 数据字典表，存储元数据的地方 表空间的offset=7页，page_type = 6
type DataDictionary struct {
	Header *DataDictionaryHeader // 52B
	// unused 4B
	SegmentHeader *SegmentHeader // 10B
	// empty space 16272 B
}

// Unmarshal 解析
func (dataDictionary *DataDictionary) Unmarshal(data []byte) {
	dataDictionary.Header.Unmarshal(data[0:52])
	dataDictionary.SegmentHeader = InitSegmentHeader(data[56:66])
}

// DataDictionaryHeader 数据字典header 52B
type DataDictionaryHeader struct {
	MaxRowID                    uint64 // 全局row_id
	MaxTableID                  uint64 // InnoDB存储引擎中的所有的表都对应一个唯一的ID，每次新建一个表时，就会把本字段的值作为该表的ID，然后自增本字段的值
	MaxIndexID                  uint64 // InnoDB存储引擎中的所有的索引都对应一个唯一的ID，每次新建一个索引时，就会把本字段的值作为该索引的ID，然后自增本字段的值
	MaxSpaceID                  uint32 // InnoDB存储引擎中的所有的表空间都对应一个唯一的ID，每次新建一个表空间时，就会把本字段的值作为该表空间的ID，然后自增本字段的值
	MixIDLow                    uint32 // 这个字段没啥用，跳过
	RootOfSysTablesClustIndex   uint32 // 本字段代表SYS_TABLES表聚簇索引的根页面的页号
	RootOfSysTableIDsClustIndex uint32 // 本字段代表SYS_TABLES表为ID列建立的二级索引的根页面的页号
	RootOfSysColumnsClustIndex  uint32 // 本字段代表SYS_COLUMNS表聚簇索引的根页面的页号
	RootOfSysIndexesClustIndex  uint32 // 本字段代表SYS_INDEXES表聚簇索引的根页面的页号
	RootOfSysFieldsClustIndex   uint32 // 本字段代表SYS_FIELDS表聚簇索引的根页面的页号
}

// Unmarshal 解析
func (dataDictionaryHeader *DataDictionaryHeader) Unmarshal(data []byte) {
	dataDictionaryHeader.MaxRowID = binary.BigEndian.Uint64(data[0:8])
	dataDictionaryHeader.MaxTableID = binary.BigEndian.Uint64(data[8:16])
	dataDictionaryHeader.MaxIndexID = binary.BigEndian.Uint64(data[16:24])
	dataDictionaryHeader.MaxSpaceID = binary.BigEndian.Uint32(data[24:28])
	dataDictionaryHeader.MixIDLow = binary.BigEndian.Uint32(data[28:32])
	dataDictionaryHeader.RootOfSysTablesClustIndex = binary.BigEndian.Uint32(data[32:36])
	dataDictionaryHeader.RootOfSysTableIDsClustIndex = binary.BigEndian.Uint32(data[36:40])
	dataDictionaryHeader.RootOfSysColumnsClustIndex = binary.BigEndian.Uint32(data[40:44])
	dataDictionaryHeader.RootOfSysIndexesClustIndex = binary.BigEndian.Uint32(data[44:48])
	dataDictionaryHeader.RootOfSysFieldsClustIndex = binary.BigEndian.Uint32(data[48:52])
}

// NewDataDictionary NewDataDictionary
func NewDataDictionary() *DataDictionary {
	return &DataDictionary{
		Header:        &DataDictionaryHeader{},
		SegmentHeader: nil,
	}
}
