package hosts

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/mishamyrt/dive/v1/pkg/paths"
	"github.com/mishamyrt/dive/v1/pkg/types"
	"github.com/mishamyrt/dive/v1/pkg/yaml"
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
		hosts[namespaceName+":"+hostName] = hostConfig
	}
	return hosts
}

func getHosts() (map[string]types.Host, []string) {
	hostMap := make(map[string]types.Host)
	config := readGlobalConfig(paths.GlobalConfig)
	hosts, _ := filepath.Glob(path.Join(paths.HostsDirectory, "*"))
	var namespaces []string
	for _, filePath := range hosts {
		host := readNamespaceHosts(filePath)
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
	if strings.Contains(restString, ":") {
		namespaceParts := strings.Split(hostName, ":")
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
	if _, err := os.Stat(".dive.yaml"); os.IsNotExist(err) {
		return localConfig, err
	}
	err := yaml.ReadFile(".dive.yaml", &localConfig)
	return localConfig, err
}

// MatchHost finds requested host in list
func MatchHost(hostName string) types.Host {
	var hostNamePart = getHostNameParts(hostName)
	var hostConfig types.Host
	hosts, namespaces := getHosts()
	localConfig, err := getLocalConfig()
	if err == nil {
		if replace, ok := localConfig.Aliases[hostNamePart.Host]; ok {
			hostNamePart.Host = replace
		}
	}
	if len(hostNamePart.Namespace) > 0 {
		if config, ok := hosts[hostNamePart.Namespace+":"+hostNamePart.Host]; ok {
			hostConfig = config
		}
	}
	for _, namespace := range namespaces {
		if config, ok := hosts[namespace+":"+hostNamePart.Host]; ok {
			hostConfig = config
		}
	}
	if len(hostConfig.Host) == 0 {
		hostConfig.Host = hostNamePart.Host
	}
	if len(hostNamePart.User) > 0 {
		hostConfig.User = hostNamePart.User
	}
	return hostConfig
}

func GetNamespaces() []string {
	_, namespaces := getHosts()
	return namespaces
}