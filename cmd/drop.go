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
	dropExample = `# Drop a database 
ssql drop databases some_database

# Drop a table
ssql drop table some_table`

	dropLong = utility.GetLongDescription("drop")
)

// dropCmd represents the drop command
var dropCmd = &cobra.Command{
	Use:                   "drop",
	DisableFlagsInUseLine: true,
	Short:                 "Drop a database or drop a table",
	Long:                  dropLong,
	Example:               dropExample,

	Run: func(cmd *cobra.Command, args []string) {
		var (
			// valid args for create table command
			validArgsTable = map[string]bool{
				"table": true,
				"TABLE": true,
			}
		)
		prevArgCheck := map[string]bool{"placeholder": false}
		utility.CommandOrchestrator("DROP", args, prevArgCheck, true, validArgsTable)
	},
}

// cobra func
func init() {
	rootCmd.AddCommand(dropCmd)
}
