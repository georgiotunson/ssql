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
	"github.com/spf13/cobra"
	"log"
	"ssql/utility"

	homedir "github.com/mitchellh/go-homedir"
)

// currentCmd represents the current command
var currentCmd = &cobra.Command{
	Use:   "current",
	Short: "Show the currently set host",
	Long: `This command shows a description of the currently set host
that is configured from using the 'ssql host set' command. The host 
must be present in your .ssql file.`,

	Run: func(cmd *cobra.Command, args []string) {
		home, err := homedir.Dir()
		if err != nil {
			log.Fatal(err)
		}
		currentHost, err := utility.GetHostState(home)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("Host: %s\nPlatform: %s\nCurrentDb: %s\n", currentHost.Host, currentHost.Platform, currentHost.CurrentDb)
	},
}

// cobra func
func init() {
	hostCmd.AddCommand(currentCmd)
}
