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
	"encoding/json"
	"fmt"
	"ibdata_parser/structure"
	"ibdata_parser/tools"
	"os"

	"github.com/spf13/cobra"
)

var writeFile string

// detailCmd represents the detail command
var detailCmd = &cobra.Command{
	Use:   "detail",
	Short: "A brief description of your command",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		parser := tools.GetParser(ibdataFile)
		page, err := parser.ParsePage(pageOffset, tools.ParserPageHeader|tools.ParserPageBody|tools.ParserPageTrailer)
		if err != nil {
			fmt.Println(err)
			return
		}

		if err != nil {
			fmt.Println(err)
			return
		}
		// 空则输出 stdout 否则写文件
		if writeFile != "" {

			file, err := os.OpenFile(writeFile, os.O_CREATE|os.O_WRONLY, 0666)
			if err != nil {
				fmt.Println(err)
				return
			}
			defer file.Close()
			os.Stdout = file
		}
		outputDetailPage(page)

	},
}

func init() {
	rootCmd.AddCommand(detailCmd)

	detailCmd.Flags().StringVarP(&writeFile, "write", "w", "", "output to file")
	detailCmd.Flags().Uint32VarP(&pageOffset, "page", "p", 0, "innodb page offset")
}

func outputDetailPage(page *structure.Page) {
	outputPageHeader(page)
	outputPageBody(page)
	outputPageTrailer(page)
}

func outputPageBody(page *structure.Page) {
	fmt.Println("File Body:")
	jsonBody, _ := json.MarshalIndent(page.Body, "", "\t")
	fmt.Println(string(jsonBody))
}
