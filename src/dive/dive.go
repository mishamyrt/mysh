package main

import (
	"fmt"
	"os"

	"../hosts"
)

func main() {
	switch os.Args[1] {
	case "get":
		fmt.Println("subcommand 'foo'")
	case "update":
		fmt.Println("subcommand 'bar'")
	default:
		hosts.GetHosts()
	}
}
