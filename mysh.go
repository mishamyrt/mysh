package main

import "github.com/mishamyrt/mysh/v1/cmd"

// GitCommit refers to commit hash at the moment of build
var GitCommit string

// Version of Mysh
var Version string

// BuildTime of Mysh
var BuildTime string

func main() {
	cmd.Execute()
}
