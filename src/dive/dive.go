package main

import (
	"fmt"
	"os"

	"../hosts"
	"../remotes"
	"../ssh"
)

func main() {
	switch os.Args[1] {
	case "get":
		remotes.DownloadConfig(os.Args[2])
	case "update":
		fmt.Println("subcommand 'bar'")
	default:
		host, _ := hosts.MatchHost(os.Args[1])
		command, err := ssh.BuildSSHCommand(host)
		if err != nil {
			panic(err)
		}
		fmt.Println(command)
	}
}
