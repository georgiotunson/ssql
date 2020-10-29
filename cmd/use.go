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
	"log"
	"ssql/utility"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	homedir "github.com/mitchellh/go-homedir"
)

var (
	useExample = `# Use a specific database. 
ssql use some_database`
)

// useCmd represents the use command
var useCmd = &cobra.Command{
	Use:                   "use",
	DisableFlagsInUseLine: true,
	Short:                 "Select a db to use for your currently set host",
	Long: `Set the db that will be stored to the state of your current session.
Any sql queries requiring a db selection will use the currently set db. Set
db reverts back to "" after using the 'ssql host set some_host' command.`,
	Example: useExample,

	Run: func(cmd *cobra.Command, args []string) {
		// validate args
		if len(args) > 1 {
			fmt.Println("Only one databases can be set at a time.")
		} else if len(args) < 1 {
			fmt.Println("Please provide a database name for the use command.")
		}

		desiredDb := args[0]

		// get user's $HOME dir
		home, err := homedir.Dir()
		if err != nil {
			log.Fatal(err)
		}

		// get currently set host
		currentHost, err := utility.GetHostState(home)
		if err != nil {
			fmt.Printf("Please be sure to use 'ssql host set' to set a host from your .ssql file\n")
			log.Fatal(err)
		}

		if viper.IsSet(currentHost.Host) == false {
			fmt.Printf("Host setting does not exist. Please make sure the desired host is configured in your .ssql file.\n")
		}

		// update host state
		if err := utility.SetHostState(home, currentHost.Host, currentHost.Platform, desiredDb); err != nil {
			log.Fatal(err)
		}
	},
}

// cobra func
func init() {
	rootCmd.AddCommand(useCmd)
}
