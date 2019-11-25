package hosts

import (
	"fmt"
	"path"
	"path/filepath"

	"../paths"
	"../types"
	"gopkg.in/yaml.v2"
)

var data = `
host: 'airport'
port: '12312312'
`

func GetHosts() {
	hosts, _ := filepath.Glob(path.Join(paths.HostsDirectory, "*"))
	for _, filePath := range hosts {
		fmt.Println(filePath)
	}

	t := types.Host{}

	_ = yaml.Unmarshal([]byte(data), &t)
	fmt.Println(t.Host)
}
