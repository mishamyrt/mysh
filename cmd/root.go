package cmd

import (
	"errors"
	"fmt"
	"os"
	"strconv"

	"github.com/mishamyrt/mysh/v1/pkg/hosts"
	"github.com/mishamyrt/mysh/v1/pkg/paths"
	"github.com/mishamyrt/mysh/v1/pkg/ssh"
	"github.com/spf13/cobra"
)

var (
	port     int
	identity string
	rootCmd  = &cobra.Command{
		Use:   "mysh <host>",
		Short: "Mysh is a tool for improving SSH user experience",
		Long:  `Mys(s)h — wrapper over SSH, which helps not to clog your head with unnecessary things. In Mysh, you can specify a remote repository with SSH hosts and connect to it by knowing only the name.`,
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) != 1 {
				return errors.New("only host is required")
			}
			return nil
		},
		Run: func(cmd *cobra.Command, args []string) {
			host, _ := hosts.MatchHost(args[0], false)
			if len(identity) > 0 {
				host.Key = identity
			}
			if port != 0 {
				host.Port = strconv.Itoa(port)
			}
			command, _ := ssh.BuildSSHCommand(host)
			fmt.Println(command)
		},
		TraverseChildren: true,
	}
)

func init() {
	rootCmd.Flags().IntVarP(&port, "port", "p", 0, "host port")
	rootCmd.Flags().StringVarP(&identity, "identity", "i", "", "identity file")
}

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
