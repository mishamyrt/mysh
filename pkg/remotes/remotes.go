package remotes

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"path"

	"github.com/mishamyrt/dive/v1/pkg/paths"
	"github.com/mishamyrt/dive/v1/pkg/types"
	"github.com/mishamyrt/dive/v1/pkg/yaml"
)

// UpdateRemotes updates all saved remotes configuration files
func UpdateRemotes() {
	remotes := GetRemotes()
	for namespace, url := range remotes.Remotes {
		// TODO: Handle namespace change
		_, err := downloadConfig(url)
		if err != nil {
			fmt.Printf("Could not update '%s'\n", namespace)
		} else {
			fmt.Printf("'%s' is updated\n", namespace)
		}
	}
}

// GetRemotes returns list of return namespaces
func GetRemotes() types.RemotesList {
	var remotesList types.RemotesList
	err := yaml.ReadFile(paths.RemotesList, &remotesList)
	if err != nil {
		fmt.Println("Could not read remotes configuration file")
		panic(err)
	}
	return remotesList
}

func saveRemoteNamespace(namespaceName string, url string) error {
	remotesList := GetRemotes()
	if _, ok := remotesList.Remotes[namespaceName]; ok {
		return nil
	}
	remotesList.Remotes[namespaceName] = url
	return yaml.WriteFile(paths.RemotesList, &remotesList)
}

func downloadConfig(url string) (string, error) {
	var config types.NamespaceConfig
	data, err := readRemoteFile(url)
	if err != nil {
		fmt.Println("Could not download remote configuration file")
		return "", err
	}
	err = yaml.Parse(data, &config)
	if err != nil {
		fmt.Println("Could not parse downloaded configuration file")
		return "", err
	}
	err = writeConfig(config.Namespace, data)
	if err != nil {
		fmt.Println("Could not save downloaded configuration file")
		return "", err
	}
	return config.Namespace, nil
}

// GetConfig downloads remote config
func GetConfig(url string) (string, error) {
	namespaceName, err := downloadConfig(url)
	if err != nil {
		return "", err
	}
	saveRemoteNamespace(namespaceName, url)
	return namespaceName, err
}

func writeConfig(namespaceName string, content []byte) error {
	return ioutil.WriteFile(path.Join(paths.HostsDirectory, namespaceName+".yaml"), content, 0644)
}

func readRemoteFile(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return []byte{}, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	return body, err
}
