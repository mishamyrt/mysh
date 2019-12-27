package paths

import (
	"os"
	"os/user"
	"path"

	"github.com/mishamyrt/mysh/v1/pkg/types"
	"github.com/mishamyrt/mysh/v1/pkg/yaml"
)

func getHomeDir() string {
	usr, _ := user.Current()
	return usr.HomeDir
}

func createIfNotExists(filePath string, data interface{}) error {
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return yaml.WriteFile(filePath, &data)
	}
	return nil
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
func PreapreEnvironment() error {
	var remotesConfig types.RemotesList
	var globalConfig types.GlobalConfig
	var err error
	if _, err = os.Stat(HostsDirectory); os.IsNotExist(err) {
		err = os.MkdirAll(HostsDirectory, os.ModePerm)
	}
	if err != nil {
		return err
	}
	err = createIfNotExists(RemotesList, remotesConfig)
	if err != nil {
		return err
	}
	err = createIfNotExists(GlobalConfig, globalConfig)
	return err
}
