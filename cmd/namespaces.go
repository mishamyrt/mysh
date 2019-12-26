package cmd

import (
	"errors"
	"fmt"

	"github.com/mishamyrt/mysh/v1/pkg/hosts"
	"github.com/mishamyrt/mysh/v1/pkg/printer"
	"github.com/mishamyrt/mysh/v1/pkg/remotes"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(remotesCmd)
	rootCmd.AddCommand(updateCmd)
	rootCmd.AddCommand(namespacesCmd)
	rootCmd.AddCommand(getCmd)
}

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "refresh hosts from added remote repositories",
	Args:  noArgumentValidator,
	Run: func(cmd *cobra.Command, args []string) {
		remotes.UpdateRemotes()
		err := hosts.BuildCompletionList()
		if err != nil {
			fmt.Println("Could not update shell completion")
		}
	},
}

var remotesCmd = &cobra.Command{
	Use:   "remotes",
	Short: "display host information",
	Args:  noArgumentValidator,
	Run: func(cmd *cobra.Command, args []string) {
		remotes := remotes.GetRemotes()
		if len(remotes.Remotes) == 0 {
			fmt.Println("No remotes were found")
		} else {
			printer.Map(remotes.Remotes)
		}
	},
}

var namespacesCmd = &cobra.Command{
	Use:   "namespaces",
	Short: "display host information",
	Args:  noArgumentValidator,
	Run: func(cmd *cobra.Command, args []string) {
		namespaces := hosts.GetNamespaces()
		if len(namespaces) == 0 {
			fmt.Println("No namespaces were found")
		} else {
			printer.List(namespaces)
		}
	},
}

var getCmd = &cobra.Command{
	Use:   "get <url>",
	Short: "add repository and download hosts from it",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("URL not provided")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		namespaceName, err := remotes.GetConfig(args[0])
		if err != nil {
			fmt.Println("can't load URL")
			return
		}
		fmt.Printf("'%s' repository successfully added\n", namespaceName)
		err = hosts.BuildCompletionList()
		if err != nil {
			fmt.Println("can't update shell completion")
		}
	},
}

// func getRemote(args []string) {
// if len(args) < 3 {
// 	fmt.Println("Can't load repository. URL not provided. Usage:")
// 	fmt.Println("\tmysh get <url>")
// 	return
// }
// namespaceName, err := remotes.GetConfig(args[2])
// if err != nil {
// 	return
// }
// fmt.Printf("'%s' repository successfully added\n", namespaceName)
// err = hosts.BuildCompletionList()
// if err != nil {
// 	fmt.Println("Could not update shell completion")
// }
// }

func noArgumentValidator(cmd *cobra.Command, args []string) error {
	if len(args) > 0 {
		return errors.New("no argument is required")
	}
	return nil
}
