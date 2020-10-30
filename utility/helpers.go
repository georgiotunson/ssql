package utility

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/viper"
	"io/ioutil"
	"os"
	"os/exec"

	homedir "github.com/mitchellh/go-homedir"
)

type hostState struct {
	Host      string `json:"Host"`
	Platform  string `json:"Platform"`
	CurrentDb string `json:"CurrentDb"`
}

// CommandOrchestrator is the function that all commands call to be exectued. It orchestrates
// validation and configuration for each command.
func CommandOrchestrator(initCommand string, args []string, prevArgCheckMap map[string]bool, checkDbRequired bool, validators ...map[string]bool) {
	// get user's $HOME dir
	home, err := homedir.Dir()
	CheckErr("Could not get $HOME dir:", err)

	// get currently set host
	currentHost, err := GetHostState(home)
	CheckErr("Could not get currently set host:", err)

	// check host config exists
	CheckErr(fmt.Sprintf("Could not get config for host %s", currentHost.Host), CheckHostConfig(currentHost.Host))
	// if config exists, get settings for current host
	hostSettings := viper.GetStringMapString(currentHost.Host)

	// mysql
	if hostSettings["platform"] == "mysql" {
		// find 'mysql' executable path
		mysqlPath, err := exec.LookPath("mysql")
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// isClauseNeedingDb checks if args includes item that requires 'use db'
		var isClauseNeedingDb bool
		if checkDbRequired {
			// if true, the option to not use db exists..
			isClauseNeedingDb = CheckArgsInclude(args, validators...)
			if isClauseNeedingDb {
				if currentHost.CurrentDb == "" {
					fmt.Println("Please use 'ssql use [dbName]' first")
					os.Exit(1)
				}
			}
		} else {
			// else, there is no option, db is needed
			// assert that a db is set
			isClauseNeedingDb = true
			if currentHost.CurrentDb == "" {
				fmt.Println("Please use 'ssql use [dbName]' first")
				os.Exit(1)
			}
		}

		// build query and query command
		dbQueryBase := initCommand
		// select does not need to wrap any values in quotes
		//prevArgCheck := map[string]bool{"placeholder": false}
		dbQueryBase = BuildDbQuery(args, prevArgCheckMap, dbQueryBase)

		var dbQuery string
		if !isClauseNeedingDb {
			dbQuery = dbQueryBase
		} else {
			dbQuery = fmt.Sprintf(`USE %s; %s;`, currentHost.CurrentDb, dbQueryBase)
		}

		query := CreateMysqlQueryCommand(dbQuery, hostSettings, mysqlPath)

		// run the command
		if err := query.Run(); err != nil {
			CheckErr(fmt.Sprintf("Could not run mysql command for query %s", dbQuery), err)
		}

		// postgres
	} else {
		// else do postgres stuff(coming soon)
	}
}

// GetLongDescription creates a long description for the passed command.
func GetLongDescription(commandName string) string {
	return fmt.Sprintf(`The %s command has most of the usual functionality
provided by the sql %s statement.`, commandName, commandName)
}

// GetHostState reads the currently set host from the user's
// $HOME/.ssqlHostState file and returns the unmarshaled json
// string as a data structure.
func GetHostState(home string) (stateData hostState, err error) {
	hostStateFile, err := ioutil.ReadFile(home + "/.ssqlHostState")
	stateData = hostState{}
	_ = json.Unmarshal([]byte(hostStateFile), &stateData)
	return stateData, err
}

// SetHostState uses the user's $HOME path and the user's desired
// host retrieved from args to create a state containing which
// host from the user's config file should be used. The state is
// stored in the file $HOME/.ssqlHostState.
func SetHostState(home, host, platform, currentDb string) (err error) {
	hostData := hostState{
		Host:      host,
		Platform:  platform,
		CurrentDb: currentDb, // current db will always default to empty
	}
	json, err := json.MarshalIndent(hostData, "", "  ")
	if err != nil {
		return err
	}
	_ = ioutil.WriteFile(home+"/.ssqlHostState", json, 0644)
	return
}

// CheckHostConfig takes the user's desired host from args and asserts
// that it is available in the user's .ssql config file.
func CheckHostConfig(desiredHost string) error {
	if viper.IsSet(desiredHost) == false {
		return fmt.Errorf("Host '%s' is not configured. Please make sure the desired host is configured in your .ssql config file.", desiredHost)
	}
	return nil
}

// CheckErr is a general function that can be used to check for errors. If
// an error is present, it will output the error to stdout and terminate the
// program.
func CheckErr(fromMsg string, err error) {
	if e := err; err != nil {
		fmt.Println(fromMsg, e)
		os.Exit(1)
	}
}

// CreateMysqlQueryCommand creates the mysql command out of the user's host config
// and passed args.
func CreateMysqlQueryCommand(dbQuery string, hostSettings map[string]string, commandPath string) *exec.Cmd {
	dbUser := fmt.Sprintf(`-u%s`, hostSettings["user"])
	dbPassword := fmt.Sprintf(`-p%s`, hostSettings["password"])
	showDbCmd := &exec.Cmd{
		Path: commandPath,
		Args: []string{commandPath, `-h`, hostSettings["host"],
			`-P`, hostSettings["port"], dbUser, dbPassword, `-e`, dbQuery},
		Stdout: os.Stdout,
		Stderr: os.Stderr,
	}
	return showDbCmd
}

// CheckArgsInclude is a generic variadic function that checks to see if an array(args)
// contains an element or elements contained in n maps. The maps are checked concurrently
// using goroutines.
func CheckArgsInclude(args []string, validators ...map[string]bool) bool {
	countGoroutine := len(validators)
	c := make(chan bool)

	// initiate n goroutines
	for i := 0; i < len(validators); i++ {
		go checkArgsIncludeChild(args, validators[i], c)
	}

	for countToComplete := 0; countToComplete < countGoroutine; countToComplete++ {
		result := <-c
		if result == true {
			return result
		}
	}
	return false
}

// checkArgsIncludeChild checks if each arg is present in a map and
// writes the results to a channel read by CheckArgsInclude.
func checkArgsIncludeChild(args []string, validator map[string]bool, c chan bool) {
	for iArg := 0; iArg < len(args); iArg++ {
		if validator[args[iArg]] == true {
			c <- true
		}
	}
	c <- false
}

// BuildDbQuery uses args passed by the user and builds a db
// query out of them.
func BuildDbQuery(args []string, prevArgCheck map[string]bool, dbQueryBase string) string {
	for i := 0; i < len(args); i++ {
		// if the previous arg is not less than 0(out of index range) and
		// the previous arg is LIKE, wrap pattern in quotes
		if i > 0 && prevArgCheck[args[i-1]] == true {
			dbQueryBase = fmt.Sprintf(`%s "%s"`, dbQueryBase, args[i])
		} else {
			dbQueryBase = dbQueryBase + " " + args[i]
		}
	}
	return dbQueryBase
}
