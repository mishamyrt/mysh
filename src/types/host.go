package types

// Host is basic SSH connection config
type Host struct {
	Host string
	Port string `yaml:",omitempty"`
	User string `yaml:",omitempty"`
}
