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
	"github.com/spf13/viper"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all configured hosts",
	Long: `Lists all of the hosts that you have configured in you
.ssql file. Password values are not shown.`,

	Run: func(cmd *cobra.Command, args []string) {
		// TODO add functionality to pass flags to only show individual
		// keys for each host. Or do only show 1 host. 1 host and specific keys
		// etc etc
		cfg := viper.AllSettings()
		for k := range cfg {
			hostSettings := viper.GetStringMapString(k)
			fmt.Printf("%s : %s => -h %s -P %s -u %s -p *******\n", k, hostSettings["platform"], hostSettings["host"], hostSettings["port"], hostSettings["user"])
		}
	},
}

// cobra func
func init() {
	hostCmd.AddCommand(listCmd)
}
