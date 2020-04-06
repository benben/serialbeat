package main

import (
	"os"

	"github.com/suda/serialbeat/cmd"

	_ "github.com/suda/serialbeat/include"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
