package hosts

import (
	"errors"
	"fmt"
	"path"
	"path/filepath"

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

func getHosts() map[string]types.Host {
	hostMap := make(map[string]types.Host)
	config := readGlobalConfig(paths.GlobalConfig)
	hosts, _ := filepath.Glob(path.Join(paths.HostsDirectory, "*"))
	for _, filePath := range hosts {
		host := readNamespaceHosts(filePath)
		namespaceHostMap := finalizeNamespacedHosts(config, host)
		for key, value := range namespaceHostMap {
			hostMap[key] = value
		}
	}
	return hostMap
}

// MatchHost finds requested host in list
func MatchHost(hostName string) (types.Host, error) {
	hosts := getHosts()
	fmt.Println(hosts)
	if hostConfig, ok := hosts[hostName]; ok {
		return hostConfig, nil
	}
	return types.Host{Host: ""}, errors.New("Host not found")
}
