package paths

import (
	"fmt"
	"github.com/mishamyrt/mysh/v1/pkg/types"
	"github.com/mishamyrt/mysh/v1/pkg/yaml"
	"os"
	"os/user"
	"path"
)

func getHomeDir() string {
	usr, _ := user.Current()
	return usr.HomeDir
}

func createIfNotExists(path string, description string, data interface{}) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		err := yaml.WriteFile(path, &data)
		if err != nil {
			fmt.Println("Сan't write " + description + " config")
			os.Exit(1)
		}
	}
}

// MyshDirectory is place, where configs, hosts and remotes are stored
var MyshDirectory = path.Join(getHomeDir(), ".local", "share", "mysh")

// GlobalConfig is the main configuration file
var GlobalConfig = path.Join(MyshDirectory, "global.yaml")

// RemotesList is the configuration file with remotes URL
var RemotesList = path.Join(MyshDirectory, "remotes.yaml")

// CompletionList is the file with finalized host list.
// It used for shell completion.
var CompletionList = path.Join(MyshDirectory, "completion")

// HostsDirectory is place, where hosts are stored
var HostsDirectory = path.Join(MyshDirectory, "hosts")

// PreapreEnvironment creates all necessary folders and files
func PreapreEnvironment() {
	if _, err := os.Stat(MyshDirectory); os.IsNotExist(err) {
		os.MkdirAll(MyshDirectory, os.ModePerm)
	}
	if _, err := os.Stat(RemotesList); os.IsNotExist(err) {
		var remotesConfig types.RemotesList
		yaml.WriteFile(RemotesList, &remotesConfig)
	}
	var remotesConfig types.RemotesList
	createIfNotExists(RemotesList, "remotes", remotesConfig)
	var globalConfig types.GlobalConfig
	createIfNotExists(GlobalConfig, "global", globalConfig)
}
