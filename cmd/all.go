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

// allCmd represents the all command
var allCmd = &cobra.Command{
	Use:   "all",
	Short: "have a quick shot",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		parser := tools.GetParser(ibdataFile)
		pages := parser.FastParseEveryPage()

		sumResult := make(map[uint16]uint32)
		for _, page := range pages {
			fmt.Printf("page offset %08d, page type <%s>\n", page.FileHeader.Offset, page.FileHeader.GetTypeName())
			sumResult[page.FileHeader.Ptype] = sumResult[page.FileHeader.Ptype] + 1
		}
		fmt.Printf("\nTotal number of page: %d\n", len(pages))
		for ptype, sum := range sumResult {
			fmt.Printf("%s: %d\n", structure.PageTypeNameDict[ptype], sum)
		}
	},
}

func init() {
	rootCmd.AddCommand(allCmd)
}
