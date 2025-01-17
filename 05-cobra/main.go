package main

import (
	cmd "05-cobra/cmds"
	"os"
)

func main() {
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
