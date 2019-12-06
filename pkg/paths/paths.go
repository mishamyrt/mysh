package paths

import (
	"os/user"
	"path"
)

func getHomeDir() string {
	usr, _ := user.Current()
	return usr.HomeDir
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
