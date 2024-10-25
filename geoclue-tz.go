package main

import (
	"fmt"
	"os"

	"github.com/zquestz/geoclue-tz/cmd"
)

func main() {
	setupSignalHandlers()

	if err := cmd.GenerateCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
