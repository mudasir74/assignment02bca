package main

import (
	"os"

	"github.com/mudasir74/assignment02bca/cli"
)

func main() {
	defer os.Exit(0)
	cmd := cli.CommandLine{}
	cmd.Run()
}
