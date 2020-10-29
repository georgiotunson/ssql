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
	"os"
	"ssql/utility"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	setExample = `# Set the current host to 'host1' from .ssql config file.
ssql host set host1

# Set the current host to 'host2' from ./.ssql2 config file.
ssql host set host2 --config ./.ssql2`

	setLong = `Set the db host that you would like to work with. The host
name must be configured in your .ssql file. You may only 
set one host at a time.`
)

// setCmd represents the set command
var setCmd = &cobra.Command{
	Use:                   "set",
	DisableFlagsInUseLine: true,
	Short:                 "Set the db host",
	Long:                  setLong,
	Example:               setExample,

	Run: func(cmd *cobra.Command, args []string) {
		// validation
		utility.CheckErr("arg validation ERROR:", validateArgsSet(args))
		// user's desired host
		desiredHost := args[0]
		// check host config exists
		utility.CheckErr(fmt.Sprintf("Could not get config for host %s.", desiredHost), utility.CheckHostConfig(desiredHost))
		// if config exists, get settings for desired host
		hostSettings := viper.GetStringMapString(desiredHost)
		// get user's $HOME dir
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		// set user's host state
		if err := utility.SetHostState(home, desiredHost, hostSettings["platform"], ""); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}

// cobra's func (no description)
func init() {
	hostCmd.AddCommand(setCmd)
}

// validateArgsSet asserts that the required HOST_NAME arg is present and
// that more than one argument is not passed to the command.
func validateArgsSet(args []string) error {
	if len(args) == 0 {
		return fmt.Errorf(`'host set' requires a host name argument.
Examples:
  ssql host set [HOST_NAME]`)
	}
	if len(args) > 1 {
		return fmt.Errorf(`You may only set one host at a time.
Examples:
  ssql host set [HOST_NAME]`)
	}
	return nil
}
