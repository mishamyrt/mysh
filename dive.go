package main

import (
	"fmt"
	"os"

	"github.com/mishamyrt/dive/v1/pkg/hosts"
	"github.com/mishamyrt/dive/v1/pkg/remotes"
	"github.com/mishamyrt/dive/v1/pkg/ssh"
)

func main() {
	switch os.Args[1] {
	case "get":
		namespaceName, err := remotes.GetConfig(os.Args[2])
		if err != nil {
			return
		}
		fmt.Printf("'%s' repository successfully added\n", namespaceName)
		err = hosts.BuildComplitionList()
		if err != nil {
			fmt.Println("Could not update shell complition")
		}
	case "update":
		remotes.UpdateRemotes()
		err := hosts.BuildComplitionList()
		if err != nil {
			fmt.Println("Could not update shell complition")
		}
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
	case "hosts":
		hosts, _ := hosts.GetHosts(true)
		fmt.Println("Hosts:")
		for host := range hosts {
			fmt.Printf("- %s\n", host)
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
