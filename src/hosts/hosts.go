package hosts

import (
	"fmt"
	"io/ioutil"
	"path"
	"path/filepath"

	"../paths"
	"../types"
	"gopkg.in/yaml.v2"
)

func readYaml(filePath string, storage interface{}) error {
	dat, err := ioutil.ReadFile(filePath)
	if err != nil {
		return err
	}
	err = yaml.Unmarshal(dat, *&storage)
	return err
}

func readGlobalConfig(filePath string) types.GlobalConfig {
	var config types.GlobalConfig
	err := readYaml(filePath, &config)
	if err != nil {
		fmt.Println("Error: cannot read config file")
		panic(err)
	}
	return config
}

func readNamespaceHosts(filePath string) types.NamespaceConfig {
	var config types.NamespaceConfig
	err := readYaml(filePath, &config)
	if err != nil {
		fmt.Println("Error: cannot read hosts file")
		panic(err)
	}
	return config
}

// GetHosts returns finalized list of hosts
func GetHosts() {
	config := readGlobalConfig(paths.GlobalConfig)
	fmt.Println(config)
	hosts, _ := filepath.Glob(path.Join(paths.HostsDirectory, "*"))
	for _, filePath := range hosts {
		host := readNamespaceHosts(filePath)
		fmt.Println(host)
	}
}
