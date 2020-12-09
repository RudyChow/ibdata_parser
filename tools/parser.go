package tools

import (
	"errors"
	"ibdata_parser/structure"
	"log"
	"os"
)

type ibdataParser struct {
	file     *os.File
	fileName string
}

// 关闭文件
func (parser *ibdataParser) Close() {
	parser.file.Close()
}

// 解析单页
func (parser *ibdataParser) ParseSinglePage(pageNo uint32) {

}

// 检查大小
func (parser *ibdataParser) CheckSize() error {
	fileInfo, _ := os.Stat(parser.fileName)

	if fileInfo.Size()%structure.PAGE_SIZE > 0 {
		return errors.New("wrong size")
	}

	return nil
}

// 获取解析器
func GetParser(fileName string) *ibdataParser {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	return &ibdataParser{
		file:     file,
		fileName: fileName,
	}
}
