/*
Copyright © 2020 Georgio Tunson

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
	"os"

	"github.com/spf13/cobra"
)

var (
	noArgsResponse = `Please provide an additional command. Options include 'list', 'set', & 'current.'

  Examples:
    ssql host list [flags]
    ssql host set [some_host]
    ssql host current`
)

// hostCmd represents the host command
var hostCmd = &cobra.Command{
	Use:   "host",
	Short: "Manage host configuration",
	Long:  `The host command is used to set, list, and show the current host.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println(noArgsResponse)
			os.Exit(1)
		}
	},
}

// cobra func
func init() {
	rootCmd.AddCommand(hostCmd)
}
