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
	"github.com/spf13/cobra"
	"ssql/utility"
)

var (
	updateExample = `# IMPORTANT: Expressions must be wrapped in quotes.
ssql update some_table set "client_value = 'joe'"`

	updateLong = utility.GetLongDescription("update")
)

// updateCmd represents the select command
var updateCmd = &cobra.Command{
	Use:                   "update",
	DisableFlagsInUseLine: true,
	Short:                 "Update existing records in tables",
	Long:                  updateLong,
	Example:               updateExample,

	Run: func(cmd *cobra.Command, args []string) {
		prevArgCheck := map[string]bool{"placeholder": false}
		utility.CommandOrchestrator("UPDATE", args, prevArgCheck, false)
	},
}

// cobra func
func init() {
	rootCmd.AddCommand(updateCmd)
}
