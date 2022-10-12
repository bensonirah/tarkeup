/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

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
	"log"
	"net/http"
	"strconv"
	s "strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/gosuri/uitable"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List existing repository",
	Long:  `List existing repository.`,
	Run: func(cmd *cobra.Command, args []string) {
		var username string = "b.randrianarison"
		var password string = "mnus2HHP4EVgNmY"

		client := &http.Client{}
		req, err := http.NewRequest("GET", "https://tools.arkeup.com/svn/PHP/", nil)

		req.SetBasicAuth(username, password)

		resp, err := client.Do(req)

		doc, err := goquery.NewDocumentFromReader(resp.Body)

		if err != nil {
			log.Fatal(err)
		}
		defer resp.Body.Close()

		t := uitable.New()
		t.AddRow("ID", "NAME", "LINK")

		doc.Find("ul li").Each(func(i int, qs *goquery.Selection) {
			url := fmt.Sprintf("https://tools.arkeup.com/svn/PHP/%strunk", qs.Text())
			name := s.Replace(qs.Text(), "/", "", 1)
			t.AddRow(strconv.Itoa(i), name, url)
		})

		fmt.Println(t)

	},
}

func init() {
	repoCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
