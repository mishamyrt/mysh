package ssh

import (
	"errors"
	"fmt"
	"os"

	"github.com/mishamyrt/mysh/v1/pkg/types"
)

func fallbackIfEmpty(value, fallback string) string {
	if len(value) > 0 {
		return value
	}

	return fallback
}

// BuildSSHCommand builds command for SSH
func BuildSSHCommand(hostConfig types.Host) (string, error) {
	if len(hostConfig.Host) == 0 {
		return "", errors.New("empty host passed")
	}
	user := fallbackIfEmpty(hostConfig.User, os.Getenv("USER"))
	port := fallbackIfEmpty(hostConfig.Port, "22")
	sshString := fmt.Sprintf("ssh %s@%s -p %s", user, hostConfig.Host, port)
	if len(hostConfig.Key) > 0 {
		sshString += fmt.Sprintf(" -i %s", hostConfig.Key)
	}
	return sshString, nil
}

// BuildRSyncPath builds part of rsync command
func BuildRSyncPath(remoteFile *types.RemoteFile) (string, error) {
	if len(remoteFile.Host.Host) == 0 {
		return "", errors.New("empty host passed")
	}
	user := fallbackIfEmpty(remoteFile.Host.User, os.Getenv("USER"))
	port := fallbackIfEmpty(remoteFile.Host.Port, "22")
	return fmt.Sprintf(
		"--rsh='ssh -p%s' %s@%s:%s",
		port, user, remoteFile.Host.Host, remoteFile.FilePath,
	), nil
}
