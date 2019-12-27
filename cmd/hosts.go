package cmd

import (
	"errors"
	"fmt"
	"os"

	"github.com/mishamyrt/mysh/v1/pkg/hosts"
	"github.com/mishamyrt/mysh/v1/pkg/types"
	"github.com/spf13/cobra"
)

func printHost(host types.Host) {
	fmt.Println("Host:", host.Host)
	if len(host.User) > 0 {
		fmt.Println("User:", host.User)
	}
	if len(host.Port) > 0 {
		fmt.Println("Port:", host.Port)
	}
}

func printHosts(hosts map[string]types.Host) {
	fmt.Println("Hosts:")
	for host := range hosts {
		fmt.Printf("- %s\n", host)
	}
}

var hostsCmd = &cobra.Command{
	Use:   "hosts",
	Short: "display all hosts",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) > 0 {
			return errors.New("no argument is required")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		hosts, _ := hosts.GetHosts(true)
		if len(hosts) == 0 {
			fmt.Println("No hosts were found")
		} else {
			printHosts(hosts)
		}
	},
}

var showCmd = &cobra.Command{
	Use:   "show <host>",
	Short: "display host information",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("requires a host argument")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		host, err := hosts.MatchHost(os.Args[2], true)
		if err != nil {
			fmt.Println("Host not found")
		} else {
			printHost(host)
		}
	},
}

func init() {
	rootCmd.AddCommand(showCmd)
	rootCmd.AddCommand(hostsCmd)
}
