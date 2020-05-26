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
	"net/url"
	"os"
	"path"

	"github.com/go-resty/resty/v2"
	"github.com/spf13/cobra"
)

// aliveCmd represents the alive command
var aliveCmd = &cobra.Command{
	Use:   "alive",
	Short: "Check if the site is alive",
	Long: `Sends a GET HTTP request to a site in order to check its liveness.
	If the site returns a status code in the range 200-299, it will be considered alive.
	If the site returns any other status code, the check will fail.
	
	When the check succedes, it will be produce an exit code of 0. Any failure will produce a difference exit code. Additionaly, there will always be a message in STDOUT when the check fails.`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		baseUrl := args[0]
		fullUrl, err := url.Parse(baseUrl)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Not a valid URL: %s\n", fullUrl)
			os.Exit(1)
		}
		fullUrl.Path = path.Join(fullUrl.Path, pathFlag)
		if fullUrl.Scheme == "" {
			fullUrl.Scheme = "https"
		}
		resp, err := resty.New().
			SetDebug(verboseFlag).
			NewRequest().
			Get(fullUrl.String())
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error checking liveness: %s\n", err.Error())
			os.Exit(2)
		}
		if !resp.IsSuccess() {
			fmt.Fprintf(os.Stderr, "The site is not alive. Response status %s\n", resp.Status())
			os.Exit(2)
		}
		fmt.Printf("The site %s is alive.\n", baseUrl)
	},
}

var pathFlag string
var verboseFlag bool

func init() {
	rootCmd.AddCommand(aliveCmd)
	aliveCmd.Flags().StringVar(&pathFlag, "path", "/_meta/alive", "The path for the alive endpoint")
	aliveCmd.Flags().BoolVarP(&verboseFlag, "verbose", "v", false, "Enable verbose mode")
}
