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
		namespaceName, err := remotes.GetConfig(os.Args[2])
		if err != nil {
			return
		}
		fmt.Printf("'%s' config successfully added\n", namespaceName)
	case "update":
		remotes.UpdateRemotes()
	case "remotes":
		remotes := remotes.GetRemotes()
		fmt.Println("Remote namespaces:")
		for namespace := range remotes.Remotes {
			fmt.Printf("- %s\n", namespace)
		}
	case "namespaces":
		namespaces := hosts.GetNamespaces()
		fmt.Println("Namespaces:")
		for _, namespace := range namespaces {
			fmt.Printf("- %s\n", namespace)
		}
	default:
		host := hosts.MatchHost(os.Args[1])
		command, err := ssh.BuildSSHCommand(host)
		if err != nil {
			panic(err)
		}
		fmt.Println(command)
	}
}
