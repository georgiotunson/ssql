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
	selectExample = `# IMPORTANT: Expressions must be wrapped in quotes.

ssql select id, name from some_table where "id > 90000" and "id < 95000"
ssql select "*" from some_table where "id < 10"

# IMPORTANT: If you need to include a specific name within an expression,
it should be wrapped in single quotes within the double quotes of the expression
or vice versa.

ssql select "*" from some_table where "some_column = 'some_value'"

ssql select id FROM some_table ORDER BY id DESC
ssql select id, name from some_table left join age using "(id)"`

	selectLong = utility.GetLongDescription("select")
)

// selectCmd represents the select command
var selectCmd = &cobra.Command{
	Use:                   "select",
	DisableFlagsInUseLine: true,
	Short:                 "Select data from tables",
	Long:                  selectLong,
	Example:               selectExample,

	Run: func(cmd *cobra.Command, args []string) {
		prevArgCheck := map[string]bool{"placeholder": false}
		utility.CommandOrchestrator("SELECT", args, prevArgCheck, false)
	},
}

// cobra func
func init() {
	rootCmd.AddCommand(selectCmd)
}
