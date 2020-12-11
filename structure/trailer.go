package structure

import "encoding/json"

// FileTrailer 页尾
type FileTrailer uint64

// GetChksum 获取校验和
func (fileTrailer FileTrailer) GetChksum() uint32 {
	return (uint32)(fileTrailer >> 32)
}

// GetPageLsn 获取lsn
func (fileTrailer FileTrailer) GetPageLsn() uint32 {
	return (uint32)(fileTrailer)
}

// MarshalJSON 序列化成json
func (fileTrailer FileTrailer) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Trailer uint64
		Chksum  uint32
		Lsn     uint32
	}{
		uint64(fileTrailer),
		fileTrailer.GetChksum(),
		fileTrailer.GetPageLsn(),
	})
}
