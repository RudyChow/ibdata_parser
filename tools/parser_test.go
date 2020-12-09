package tools_test

import (
	"ibdata_parser/tools"
	"testing"
)

// go test -v ./tools/parser_test.go -test.run TestParseSinglePage
func TestParseSinglePage(t *testing.T) {

	ibdataParser := tools.GetParser("../ibdata1")

	ibdataParser.CheckSize()
	page, err := ibdataParser.ParsePage(0, tools.ParserPageAll)

	if err != nil {
		t.Error(err)
		t.Fail()
	}
	t.Logf("%+v", page.FileHeader)

}

// go test -v ./tools/parser_test.go -test.run TestFastParseEveryPage
func TestFastParseEveryPage(t *testing.T) {

	ibdataParser := tools.GetParser("../ibdata1")

	ibdataParser.CheckSize()
	ibdataParser.FastParseEveryPage()

	// if err != nil {
	// 	t.Error(err)
	// 	t.Fail()
	// }

}
