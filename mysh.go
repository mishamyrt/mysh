package main

import (
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/mishamyrt/mysh/v1/pkg/hosts"
	"github.com/mishamyrt/mysh/v1/pkg/remotes"
	"github.com/mishamyrt/mysh/v1/pkg/ssh"
)

var versionTemplate = `Version:	{{.Version}}
Git commit:	{{.GitCommit}}
Built:	{{.BuildTime}}
{{- end}}`

// GitCommit refers to commit hash at the moment of build
var GitCommit string

// Version of Mysh
var Version string

// BuildTime of Mysh
var BuildTime string

func main() {
	switch os.Args[1] {
	case "get":
		namespaceName, err := remotes.GetConfig(os.Args[2])
		if err != nil {
			return
		}
		fmt.Printf("'%s' repository successfully added\n", namespaceName)
		err = hosts.BuildCompletionList()
		if err != nil {
			fmt.Println("Could not update shell completion")
		}
	case "update":
		remotes.UpdateRemotes()
		err := hosts.BuildCompletionList()
		if err != nil {
			fmt.Println("Could not update shell completion")
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
	case "show":
		host, err := hosts.MatchHost(os.Args[2], true)
		if err != nil {
			fmt.Println("Host not found")
		}
		fmt.Println("Host:", host.Host)
		if len(host.User) > 0 {
			fmt.Println("User:", host.User)
		}
		if len(host.Port) > 0 {
			fmt.Println("Port:", host.Port)
		}
	case "version":
		w := tabwriter.NewWriter(os.Stdout, 14, 1, 1, ' ', 0)
		fmt.Fprintf(w, "Version:\t%s\n", Version)
		fmt.Fprintf(w, "GitCommit:\t%s\n", GitCommit)
		fmt.Fprintf(w, "Built:\t%s\n", BuildTime)
		w.Flush()
	default:
		host, _ := hosts.MatchHost(os.Args[1], false)
		command, err := ssh.BuildSSHCommand(host)
		if err != nil {
			panic(err)
		}
		fmt.Println(command)
	}
}
