package hosts

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/mishamyrt/mysh/v1/pkg/paths"
	"github.com/mishamyrt/mysh/v1/pkg/types"
	"github.com/mishamyrt/mysh/v1/pkg/yaml"
)

func readGlobalConfig(filePath string) types.GlobalConfig {
	var config types.GlobalConfig
	err := yaml.ReadFile(filePath, &config)
	if err != nil {
		fmt.Println("Error: cannot read config file")
		panic(err)
	}
	return config
}

func readNamespaceHosts(filePath string) types.NamespaceConfig {
	var config types.NamespaceConfig
	err := yaml.ReadFile(filePath, &config)
	if err != nil {
		fmt.Println("Error: cannot read hosts file")
		panic(err)
	}
	return config
}

func finalizeNamespacedHosts(
	globalConfig types.GlobalConfig,
	namespaceConfig types.NamespaceConfig,
) map[string]types.Host {
	hosts := make(map[string]types.Host)
	namespaceName := namespaceConfig.Namespace
	var fallbackUser = globalConfig.Namespaces[namespaceName].User
	for hostName, hostConfig := range namespaceConfig.Hosts {
		if len(hostConfig.User) == 0 && len(fallbackUser) > 0 {
			hostConfig.User = fallbackUser
		}
		hosts[namespaceName+"/"+hostName] = hostConfig
	}
	return hosts
}

// BuildCompletionList builds and saves list for shell completion
func BuildCompletionList() error {
	hosts, _ := GetHosts(false)
	var complitions string
	for hostName := range hosts {
		complitions += hostName + " "
	}
	return ioutil.WriteFile(paths.CompletionList, []byte(complitions), 0644)
}

// GetHosts returns final map of hosts
func GetHosts(stripOrig bool) (map[string]types.Host, []string) {
	hostMap := make(map[string]types.Host)
	config := readGlobalConfig(paths.GlobalConfig)
	hosts, _ := filepath.Glob(path.Join(paths.HostsDirectory, "*"))
	namespaces := make([]string, len(hosts))
	for _, filePath := range hosts {
		host := readNamespaceHosts(filePath)
		if !stripOrig {
			for key, value := range host.Hosts {
				hostMap[key] = value
			}
		}
		namespaces = append(namespaces, host.Namespace)
		namespaceHostMap := finalizeNamespacedHosts(config, host)
		for key, value := range namespaceHostMap {
			hostMap[key] = value
		}
	}
	return hostMap, namespaces
}

func getHostNameParts(hostName string) types.HostNameParts {
	var parts types.HostNameParts
	restString := hostName
	if strings.Contains(restString, "/") {
		namespaceParts := strings.Split(hostName, "/")
		parts.Namespace = namespaceParts[0]
		restString = namespaceParts[1]
	}
	if strings.Contains(restString, "@") {
		userParts := strings.Split(hostName, "@")
		parts.User = userParts[0]
		restString = userParts[1]
	}
	parts.Host = restString
	return parts
}

// getLocalConfig loads the local configuration file
// with aliases
func getLocalConfig() (types.LocalConfig, error) {
	var localConfig types.LocalConfig
	if _, err := os.Stat(".mysh.yaml"); os.IsNotExist(err) {
		return localConfig, err
	}
	err := yaml.ReadFile(".mysh.yaml", &localConfig)
	return localConfig, err
}

// MatchRemoteFile finds requested host with file in list
func MatchRemoteFile(remoteFilePath string) types.RemoteFile {
	var remoteFile types.RemoteFile
	remoteFileParts := strings.Split(remoteFilePath, ":")
	remoteFile.FilePath = remoteFileParts[1]
	remoteFile.Host, _ = MatchHost(remoteFileParts[0], false)
	return remoteFile
}

// MatchHost finds requested host in list
func MatchHost(hostName string, strict bool) (types.Host, error) {
	var hostNamePart = getHostNameParts(hostName)
	var hostConfig types.Host
	hosts, namespaces := GetHosts(true)
	localConfig, err := getLocalConfig()
	if err == nil {
		if replace, ok := localConfig.Aliases[hostNamePart.Host]; ok {
			hostNamePart.Host = replace
		}
	}
	if len(hostNamePart.Namespace) > 0 {
		if config, ok := hosts[hostNamePart.Namespace+"/"+hostNamePart.Host]; ok {
			hostConfig = config
		}
	}
	for _, namespace := range namespaces {
		if config, ok := hosts[namespace+"/"+hostNamePart.Host]; ok {
			hostConfig = config
		}
	}
	if len(hostConfig.Host) == 0 {
		if strict {
			return types.Host{Host: ""}, errors.New("not found")
		}
		hostConfig.Host = hostNamePart.Host
	}
	if len(hostNamePart.User) > 0 {
		hostConfig.User = hostNamePart.User
	}
	return hostConfig, nil
}

// GetNamespaces returns list of hosts namespaces
func GetNamespaces() []string {
	_, namespaces := GetHosts(true)
	return namespaces
}
