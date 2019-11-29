package hosts

import (
	"fmt"
	"path"
	"path/filepath"
	"strings"

	"../paths"
	"../types"
	"../yaml"
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

// MatchHost finds requested host in list
func MatchHost(hostName string) types.Host {
	var hostNamePart = getHostNameParts(hostName)
	var hostConfig types.Host
	hosts, namespaces := getHosts()
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
