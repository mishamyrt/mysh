package yaml

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

// Parse YAML raw bytes data to given link
func Parse(data []byte, storage interface{}) error {
	return yaml.Unmarshal(data, *&storage)
}

// ReadFile reads YAML file to given link
func ReadFile(filePath string, storage interface{}) error {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return err
	}
	return Parse(data, *&storage)
}

// WriteFile writes given data to YAML file
func WriteFile(filePath string, content interface{}) error {
	data, err := yaml.Marshal(*&content)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(filePath, data, 0644)
	return err
}
