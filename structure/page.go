package structure

import "encoding/binary"

// page info
const (
	PAGE_SIZE = 64 * 1024

	PAGE_HEADER_SIZE  = 38
	PAGE_TRAILER_SIZE = 8
)

// page type
const (
	FIL_PAGE_TYPE_ALLOCATED = iota
	_
	FIL_PAGE_UNDO_LOG
	FIL_PAGE_INODE
	FIL_PAGE_IBUF_FREE_LIST
	FIL_PAGE_IBUF_BITMAP
	FIL_PAGE_TYPE_SYS
	FIL_PAGE_TYPE_TRX_SYS
	FIL_PAGE_TYPE_FSP_HDR
	FIL_PAGE_TYPE_XDES
	FIL_PAGE_TYPE_BLOB
	FIL_PAGE_TYPE_INDEX = 0x45BF
)

type Page struct {
	FileHeader  *FileHeader
	Body        interface{}
	FileTrailer FileTrailer
}

type FileHeader struct {
	space_or_chksum         uint32
	offset                  uint32
	perv                    uint32
	next                    uint32
	lsn                     uint64
	ptype                   uint16
	file_flush_lsn          uint64
	arch_log_no_or_space_id uint32
}

func (fileHeader *FileHeader) Unmarshal(data []byte) {

	fileHeader.space_or_chksum = binary.BigEndian.Uint32(data[0:4])
	fileHeader.offset = binary.BigEndian.Uint32(data[4:8])
	fileHeader.perv = binary.BigEndian.Uint32(data[8:12])
	fileHeader.next = binary.BigEndian.Uint32(data[12:16])
	fileHeader.lsn = binary.BigEndian.Uint64(data[16:24])
	fileHeader.ptype = binary.BigEndian.Uint16(data[24:26])
	fileHeader.file_flush_lsn = binary.BigEndian.Uint64(data[26:34])
	fileHeader.arch_log_no_or_space_id = binary.BigEndian.Uint32(data[34:38])
}

type FileTrailer uint64

func NewPage() *Page {
	return &Page{
		FileHeader:  &FileHeader{},
		Body:        nil,
		FileTrailer: 0,
	}
}
