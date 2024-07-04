package main

import (
	"os"

	"github.com/pascaliske/magicmirror/cmd"
	"github.com/pascaliske/magicmirror/logger"
)

func main() {
	if err := cmd.Execute(); err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}
}
