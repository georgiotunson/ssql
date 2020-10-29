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
	deleteExample = `# IMPORTANT NOTES:
 - Expressions must be wrapped in quotes.
 - If you need to reference a specific name within an
   expression,it should be wrapped in single quotes
   within the double quotes of the expression or vice versa.
    
# Delete item from table
ssql delete from zoo where "name='cow'"
ssql delete from table_name where "column_name = 'some_column'"`

	deleteLong = utility.GetLongDescription("delete")
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:                   "delete",
	DisableFlagsInUseLine: true,
	Short:                 "Delete existing records in a table",
	Long:                  deleteLong,
	Example:               deleteExample,

	Run: func(cmd *cobra.Command, args []string) {
		var (
			// valid args for show tables command
			validArgsDatabase = map[string]bool{
				"database": true,
				"DATABASE": true,
			}
		)

		prevArgCheck := map[string]bool{"placeholder": false}
		utility.CommandOrchestrator("DELETE", args, prevArgCheck, true, validArgsDatabase)
	},
}

// cobra func
func init() {
	rootCmd.AddCommand(deleteCmd)
}
