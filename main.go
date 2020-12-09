package main

import (
	"ibdata_parser/tools"
	"log"
)

func main() {
	ibdataParser := tools.GetParser("ibdata1")

	// 检查文件大小
	err := ibdataParser.CheckSize()
	if err != nil {
		log.Fatal(err)
	}

	// 先定义缓冲区
	// byteHeader := make([]byte, 38)
	// bytesRead, _ := f.Read(byteHeader)
	// log.Printf("Number of bytes read: %d\n", bytesRead)

	// header := &structure.FileHeader{}
	// header.Unmarshal(byteHeader)
	// log.Println(header)

	// log.Println(
	// 	structure.FIL_PAGE_TYPE_ALLOCATED,
	// 	structure.FIL_PAGE_UNDO_LOG,
	// 	structure.FIL_PAGE_TYPE_BLOB,
	// 	structure.FIL_PAGE_TYPE_INDEX)

	// if err := f.Close(); err != nil {
	// 	log.Fatal(err)
	// }
}
