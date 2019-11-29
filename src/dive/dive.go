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
		namespaceName := remotes.GetConfig(os.Args[2])
		fmt.Printf("'%s' config successfully added\n", namespaceName)
	case "update":
		remotes.UpdateRemotes()
	case "remotes":
		remotes := remotes.GetRemotes()
		fmt.Println("Remote namespaces:")
		for namespace := range remotes.Remotes {
			fmt.Printf("- %s\n", namespace)
		}
	default:
		host, _ := hosts.MatchHost(os.Args[1])
		command, err := ssh.BuildSSHCommand(host)
		if err != nil {
			panic(err)
		}
		fmt.Println(command)
	}
}
