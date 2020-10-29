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
	showExample = `# Show all databases for current host.
ssql show databases
ssql show schemas

# Show tables for current database(must use 'ssql use some_database' first).
ssql use some_database
ssql show tables`

	showLong = utility.GetLongDescription("show")
)

// showCmd represents the show command
var showCmd = &cobra.Command{
	Use:                   "show",
	DisableFlagsInUseLine: true,
	Short:                 "Show databases or show tables",
	Long:                  showLong,
	Example:               showExample,

	Run: func(cmd *cobra.Command, args []string) {
		var (
			// valid args for show tables command
			validArgsTables = map[string]bool{
				"tables": true,
				"TABLES": true,
			}

			// valid args for show tables command
			validArgsColumn = map[string]bool{
				"COLUMNS": true,
				"columns": true,
			}

			// valid args for like command
			validArgsLike = map[string]bool{
				"LIKE": true,
				"like": true,
			}
		)

		utility.CommandOrchestrator("SHOW", args, validArgsLike, true, validArgsTables, validArgsColumn)
	},
}

// cobra func
func init() {
	rootCmd.AddCommand(showCmd)
}
