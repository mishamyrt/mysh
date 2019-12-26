package cmd

import (
	"fmt"
	"os"

	"github.com/mishamyrt/mysh/v1/pkg/hosts"
	"github.com/mishamyrt/mysh/v1/pkg/paths"
	"github.com/mishamyrt/mysh/v1/pkg/ssh"
	"github.com/spf13/cobra"
)

func connect(args []string) {
	var privateKey string
	var hostString string
	for i := 0; i < len(args); i++ {
		if args[i] == "-i" {
			privateKey = args[i+1]
			i++
		} else {
			hostString = args[i]
		}
	}
	host, _ := hosts.MatchHost(hostString, false)
	if len(privateKey) > 0 {
		host.Key = privateKey
	}
	command, err := ssh.BuildSSHCommand(host)
	if err != nil {
		panic(err)
	}
	fmt.Println(command)
}

var (
	rootCmd = &cobra.Command{
		Use:   "mysh <host>",
		Short: "Mysh is a tool for improving SSH user experience",
		Long:  `Mys(s)h — wrapper over SSH, which helps not to clog your head with unnecessary things. In Mysh, you can specify a remote repository with SSH hosts and connect to it by knowing only the name.`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(args)
		},
	}
)

// Execute is
func Execute() {
	err := paths.PreapreEnvironment()
	if err != nil {
		fmt.Println("Сan't initialize configuration files")
		os.Exit(1)
	}
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
