package cmd

import (
	"fmt"
	"os"
	"unicode"

	"github.com/spf13/cobra"
)

const (
	appName = "geoclue-tz"
	version = "0.0.1"
)

// Stores configuration data.
var config Config

// GenerateCmd is the main command for Cobra.
var GenerateCmd = &cobra.Command{
	Use:   "geoclue-tz",
	Short: "Generate /etc/geolocation from timezone.",
	Long:  `Generate /etc/geolocation from timezone.`,
	Run: func(cmd *cobra.Command, args []string) {
		err := performCommand(cmd, args)
		if err != nil {
			bail(err)
		}
	},
}

func init() {
	err := config.Load()
	if err != nil {
		bail(fmt.Errorf("failed to load configuration: %s", err))
	}

	prepareFlags()
}

func bail(err error) {
	fmt.Fprintf(os.Stderr, "[Error] %s\n", capitalize(err.Error()))
	os.Exit(1)
}

func capitalize(str string) string {
	if len(str) == 0 {
		return ""
	}
	tmp := []rune(str)
	tmp[0] = unicode.ToUpper(tmp[0])
	return string(tmp)
}

func prepareFlags() {
	GenerateCmd.PersistentFlags().BoolVarP(
		&config.DisplayVersion, "version", "", false, "display version")
	GenerateCmd.PersistentFlags().BoolVarP(
		&config.Verbose, "verbose", "v", config.Verbose, "verbose mode")
	GenerateCmd.PersistentFlags().BoolVarP(
		&config.DryRun, "dryrun", "d", config.DryRun, "dryrun mode")
	GenerateCmd.PersistentFlags().Float32VarP(
		&config.DefaultLatitude, "default-lat", "", config.DefaultLatitude, "default latitude")
	GenerateCmd.PersistentFlags().Float32VarP(
		&config.DefaultLongitude, "default-long", "", config.DefaultLongitude, "default longitude")
}

// Where all the work happens.
func performCommand(cmd *cobra.Command, args []string) error {
	if config.DisplayVersion {
		fmt.Printf("%s %s\n", appName, version)
		return nil
	}

	if len(args) != 0 {
		help := cmd.HelpFunc()
		help(cmd, args)
	}

	return nil
}
