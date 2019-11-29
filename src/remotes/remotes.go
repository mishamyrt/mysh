package remotes

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"path"

	"../paths"
	"../types"
	"../yaml"
)

func saveRemoteNamespace(namespaceName string, url string) error {
	var remotesList types.RemotesList
	err := yaml.ReadFile(paths.RemotesList, &remotesList)
	if err != nil {
		return err
	}
	if _, ok := remotesList.Remotes[namespaceName]; ok {
		return nil
	}
	remotesList.Remotes[namespaceName] = url
	return yaml.WriteFile(paths.RemotesList, &remotesList)
}

// DownloadConfig downloads remote config
func DownloadConfig(url string) string {
	var config types.NamespaceConfig
	data, err := readRemoteFile(url)
	if err != nil {
		fmt.Println("Could not download remote configuration file")
		panic(err)
	}
	yaml.Parse(data, &config)
	namespaceName := config.Namespace
	writeConfig(namespaceName, data)
	saveRemoteNamespace(namespaceName, url)
	return namespaceName
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

// Create the file
// out, err := os.Create(filepath)
// if err != nil {
// 	return "", err
// }
// defer out.Close()

// // Write the body to file
// _, err = io.Copy(out, resp.Body)
// return "", err
