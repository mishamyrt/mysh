package ssh

import (
	"errors"
	"fmt"
	"os"

	"../types"
)

func fallbackIfEmpty(value string, fallback string) string {
	if len(value) > 0 {
		return value
	}
	return fallback
}

// BuildSSHCommand builds command for SSH
func BuildSSHCommand(hostConfig types.Host) (string, error) {
	if len(hostConfig.Host) == 0 {
		return "", errors.New("Empty host passed")
	}
	user := fallbackIfEmpty(hostConfig.User, os.Getenv("USER"))
	port := fallbackIfEmpty(hostConfig.Port, "22")
	return fmt.Sprintf("ssh %s@%s -p %s", user, hostConfig.Host, port), nil
}
