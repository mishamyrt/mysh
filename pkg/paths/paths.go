package paths

import (
	"os/user"
	"path"
)

func getHomeDir() string {
	usr, _ := user.Current()
	return usr.HomeDir
}

// DiveDirectory is place, where configs, hosts and remotes are stored
var DiveDirectory = path.Join(getHomeDir(), ".local", "share", "dive")

// GlobalConfig is the main configuration file
var GlobalConfig = path.Join(DiveDirectory, "global.yaml")

// RemotesList is the configuration file with remotes URL
var RemotesList = path.Join(DiveDirectory, "remotes.yaml")

// CompletionList is the file with finalized host list.
// It used for shell completion.
var CompletionList = path.Join(DiveDirectory, "completion")

// HostsDirectory is place, where hosts are stored
var HostsDirectory = path.Join(DiveDirectory, "hosts")
