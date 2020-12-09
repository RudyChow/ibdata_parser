/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"ibdata_parser/structure"
	"ibdata_parser/tools"

	"github.com/spf13/cobra"
)

var pageOffset uint32

// simpleCmd represents the simple command
var simpleCmd = &cobra.Command{
	Use:   "simple",
	Short: "get a simple page (only show file header and file trailer)",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		parser := tools.GetParser(ibdataFile)
		page, err := parser.ParsePage(pageOffset, tools.ParserPageHeader|tools.ParserPageTrailer)
		if err != nil {
			fmt.Println(err)
			return
		}

		outputSimplePage(page)

	},
}

func init() {
	rootCmd.AddCommand(simpleCmd)

	simpleCmd.Flags().Uint32VarP(&pageOffset, "page", "p", 0, "innodb page offset")
}

func outputSimplePage(page *structure.Page) {
	fmt.Println("File Header:")
	fmt.Printf("space or checksum -> %d\n", page.FileHeader.SpaceOrChksum)
	fmt.Printf("offset -> %d\n", page.FileHeader.Offset)
	fmt.Printf("perv page -> %d\n", page.FileHeader.Perv)
	fmt.Printf("next page -> %d\n", page.FileHeader.Next)
	fmt.Printf("lsn -> %d\n", page.FileHeader.Lsn)
	fmt.Printf("type -> %d (%s)\n", page.FileHeader.Ptype, page.FileHeader.GetTypeName())
	fmt.Printf("file flush lsn -> %d\n", page.FileHeader.FileFlushLsn)
	fmt.Printf("arch log no or space id -> %d\n", page.FileHeader.ArchLogNoOrSpaceID)

	fmt.Println("")

	fmt.Println("File Trailer:")
	fmt.Printf("all -> %d\n", page.FileTrailer)
	fmt.Printf("checksum -> %d\n", page.FileTrailer.GetChksum())
	fmt.Printf("lsn -> %d\n", page.FileTrailer.GetPageLsn())
}
