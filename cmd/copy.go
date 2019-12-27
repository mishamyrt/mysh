package cmd

import (
	"errors"
	"fmt"
	"strings"

	"github.com/mishamyrt/mysh/v1/pkg/hosts"
	"github.com/mishamyrt/mysh/v1/pkg/ssh"
	"github.com/mishamyrt/mysh/v1/pkg/types"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(copyCmd)
}

func isFlag(arg string) bool {
	return arg[0:1] == "-"
}

func isHostPath(arg string) bool {
	return strings.Contains(arg, ":")
}

func getFlags(args []string) string {
	var flags string
	for _, arg := range args {
		if arg[0:1] == "-" {
			flags += args[2][1:len(args[2])]
		}
	}
	return flags
}

// if (1 || 2) flags: err

// if 1 == remote: target{type: "remote", path: 1}
// else: source{type: "local", path: 1}

// if 2 == remote: target{type: "remote", path: 2}
// else: target{type: "local", path: 2}

// ______

// if (source && target) remote: rsync reverse channel
// else: rsync

func identifyPath(path string) (types.RSyncFile, error) {
	var file types.RSyncFile
	file.IsRemote = isHostPath(path)
	if file.IsRemote {
		pathMatch := hosts.MatchRemoteFile(path)
		path, err := ssh.BuildRSyncPath(&pathMatch)
		file.Path = path
		return file, err
	}
	file.Path = path
	return file, nil
}

var (
	source     types.RSyncFile
	target     types.RSyncFile
	rsyncFlags = "av8h"
	copyCmd    = &cobra.Command{
		Use:   "copy <source> <target>",
		Short: "copy file to, from or between servers",
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) < 2 {
				return errors.New("too few arguments")
			}
			if isFlag(args[0]) {
				return errors.New("source host looks like flag")
			}
			if isFlag(args[1]) {
				return errors.New("target host looks like flag")
			}
			var err error
			source, err = identifyPath(args[0])
			if err != nil {
				return errors.New("can't find source host")
			}
			target, err = identifyPath(args[1])
			if err != nil {
				return errors.New("can't find target host")
			}
			return nil
		},
		Run: func(cmd *cobra.Command, args []string) {
			rsyncFlags += getFlags(args)
			if source.IsRemote && target.IsRemote {
				fmt.Println("host to host copy is not implemented yet:(")
				return
			}
			fmt.Printf("rsync -%s %s %s\n", rsyncFlags, source.Path, target.Path)
		},
		DisableFlagParsing: true,
	}
)

// case "copy":
//     copyFile(os.Args)

//     func copyFile(args []string) {
//         usage := "\tmysh copy [args] <source> <destination>"
//         pathOffset := 2
//         rsyncArgs := "av8h"
//         if args[2][0:1] == "-" {
//             pathOffset++
//             rsyncArgs += args[2][1:len(args[2])]
//         }
//         var source string
//         var target string
//         if len(args) < 3 {
//             fmt.Println("Source not provided. Usage:")
//             fmt.Println(usage)
//             return
//         } else if len(args) < 4 {
//             fmt.Println("Target not provided. Usage:")
//             fmt.Println(usage)
//             return
//         }
//         source, err := getRSyncFile(args[pathOffset])
//         if err != nil {
//             fmt.Println("Source not provided. Usage:")
//             fmt.Println(usage)
//         }
//         target, err = getRSyncFile(args[pathOffset+1])
//         if err != nil {
//             fmt.Println("Target not provided. Usage:")
//             fmt.Println(usage)
//         }
//         fmt.Printf("rsync -%s %s %s\n", rsyncArgs, source, target)
//     }
