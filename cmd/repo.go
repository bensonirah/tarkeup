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
	"os"
	"strconv"
	s "strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/aquasecurity/table"
	"github.com/spf13/cobra"
)

// repoCmd represents the repo command
var repoCmd = &cobra.Command{
	Use:   "repo",
	Short: "Manage arkeup repository",
	Long:  `Manage an arkeup repository`,
	Run: func(cmd *cobra.Command, args []string) {

		var username string = "test"
		var password string = "test"

		client := &http.Client{}
		req, err := http.NewRequest("GET", "https://tools.arkeup.com/svn/PHP/", nil)

		req.SetBasicAuth(username, password)

		resp, err := client.Do(req)

		doc, err := goquery.NewDocumentFromReader(resp.Body)

		if err != nil {
			log.Fatal(err)
		}
		defer resp.Body.Close()

		t := table.New(os.Stdout)
		t.SetHeaders("ID", "Name", "Url")
		t.SetBorders(false)

		doc.Find("ul li").Each(func(i int, qs *goquery.Selection) {
			url := fmt.Sprintf("https://tools.arkeup.com/svn/PHP/%strunk", qs.Text())
			name := s.Replace(qs.Text(), "/", "", 1)
			t.AddRow(strconv.Itoa(i), name, url)
		})
		t.Render()
	},
}

func init() {
	rootCmd.AddCommand(repoCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// repoCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// repoCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
