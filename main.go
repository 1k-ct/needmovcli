package main

import (
	"os"

	"github.com/1k-ct/clonefile/prac/cmd"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		// fmt.Println(os.Stderr, "%s: %v\n", os.Args[0], err)
		os.Exit(-1)
	}
}
