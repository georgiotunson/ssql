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
	createExample = `# Create a new table for a database. 
ssql create table cow "(name TEXT, age INTEGER, PRIMARY KEY(age))"

# Create new database
ssql create database database_name`

	createLong = utility.GetLongDescription("create")
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:                   "create",
	DisableFlagsInUseLine: true,
	Short:                 "Create new database or table",
	Long:                  createLong,
	Example:               createExample,

	Run: func(cmd *cobra.Command, args []string) {
		var (
			// valid args for create table command
			validArgsTable = map[string]bool{
				"table": true,
				"TABLE": true,
			}
		)
		prevArgCheck := map[string]bool{"placeholder": false}
		utility.CommandOrchestrator("CREATE", args, prevArgCheck, true, validArgsTable)
	},
}

// cobra func
func init() {
	rootCmd.AddCommand(createCmd)
}
