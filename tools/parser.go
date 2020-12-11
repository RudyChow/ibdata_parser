package tools

import (
	"errors"
	"ibdata_parser/structure"
	"log"
	"os"
)

// func ParsePage's param
const (
	ParserPageHeader = 1 << iota
	ParserPageBody
	ParserPageTrailer
	ParserPageAll = ParserPageHeader | ParserPageBody | ParserPageTrailer
)

// IbdataParser ibdata解析器
type IbdataParser struct {
	file     *os.File
	fileInfo os.FileInfo
}

// Close 关闭文件
func (parser *IbdataParser) Close() {
	parser.file.Close()
}

// ParsePage 解析单页
func (parser *IbdataParser) ParsePage(pageNo uint32, parserType uint8) (*structure.Page, error) {

	// readat
	buf := make([]byte, structure.PageSize)
	_, err := parser.file.ReadAt(buf, int64(pageNo*structure.PageSize))
	if err != nil {
		return nil, err
	}

	page := structure.NewPage()
	// 解析头部(如果解析body,必须要解析header)
	if parserType&ParserPageHeader == ParserPageHeader || parserType&ParserPageBody == ParserPageBody {
		page.UnmarshalHeader(buf[structure.PageHeaderOffset:structure.PageHeaderSize])
	}
	// 解析body
	if parserType&ParserPageBody == ParserPageBody {
		page.UnmarshalBody(buf[structure.PageBodyOffset : structure.PageBodyOffset+structure.PageBodySize])
	}
	// 解析尾部
	if parserType&ParserPageTrailer == ParserPageTrailer {
		page.UnmarshalTrailer(buf[structure.PageTrailerOffset : structure.PageTrailerOffset+structure.PageTrailerSize])
	}

	return page, nil
}

// FastParseEveryPage 快速解析每个页
func (parser *IbdataParser) FastParseEveryPage() []*structure.Page {
	pageCount := (uint32)(parser.fileInfo.Size() / structure.PageSize)

	pageResult := make([]*structure.Page, pageCount)
	var i uint32 = 0
	for ; i < pageCount; i++ {
		page, _ := parser.ParsePage(i, ParserPageHeader)
		pageResult[i] = page
	}

	return pageResult
}

// CheckSize 检查大小
func (parser *IbdataParser) CheckSize() error {

	if parser.fileInfo.Size()%structure.PageSize > 0 {
		return errors.New("wrong size")
	}

	return nil
}

// GetParser 获取解析器
func GetParser(fileName string) *IbdataParser {
	// 打开文件
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	// 查看文件信息
	fileInfo, err := os.Stat(fileName)
	if err != nil {
		log.Fatal(err)
	}
	return &IbdataParser{
		file:     file,
		fileInfo: fileInfo,
	}
}
