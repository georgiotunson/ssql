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
	insertExample = `# IMPORTANT NOTES: 
	- Expressions must be wrapped in quotes.
	- If you need to input a specific name within an expression,
		it should be wrapped in single quotes within the double quotes of the expression
		or vice versa.

# Insert into a table 
ssql insert into "table(name)" "values('Fred')"`

	insertLong = utility.GetLongDescription("insert")
)

// insertCmd represents the insert command
var insertCmd = &cobra.Command{
	Use:                   "insert",
	DisableFlagsInUseLine: true,
	Short:                 "Insert rows into an existing table",
	Long:                  insertLong,
	Example:               insertExample,

	Run: func(cmd *cobra.Command, args []string) {
		prevArgCheck := map[string]bool{"placeholder": false}
		utility.CommandOrchestrator("INSERT", args, prevArgCheck, false)
	},
}

// cobra func
func init() {
	rootCmd.AddCommand(insertCmd)
}
