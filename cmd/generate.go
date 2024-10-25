package cmd

import (
	"errors"
	"fmt"
	"os"
	"unicode"

	"github.com/spf13/cobra"
	"github.com/zquestz/geoclue-tz/tz"
)

const (
	appName = "geoclue-tz"
	version = "0.5.0"
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
	GenerateCmd.PersistentFlags().BoolVar(
		&config.DisplayVersion, "version", false, "display version")
	GenerateCmd.PersistentFlags().BoolVarP(
		&config.Verbose, "verbose", "v", config.Verbose, "verbose mode")
	GenerateCmd.PersistentFlags().BoolVarP(
		&config.DryRun, "dry-run", "d", config.DryRun, "dry run mode")
	GenerateCmd.PersistentFlags().Float32VarP(
		&config.HomeLatitude, "home-latitude", "", config.HomeLatitude, "home latitude")
	GenerateCmd.PersistentFlags().Float32VarP(
		&config.HomeLongitude, "home-longitude", "", config.HomeLongitude, "home longitude")
	GenerateCmd.PersistentFlags().Float32VarP(
		&config.HomeAltitude, "home-altitude", "", config.HomeAltitude, "home altitude")
	GenerateCmd.PersistentFlags().Float32VarP(
		&config.HomeAccuracy, "home-accuracy", "", config.HomeAccuracy, "home accuracy")
	GenerateCmd.PersistentFlags().BoolVar(
		&config.Home, "home", config.Home, "enable home configuration")
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

	if config.DryRun {
		config.Verbose = true
	}

	if config.Home {
		location, err := homeLocation()
		if err != nil {
			return err
		}

		err = location.WriteGeolocation(config.DryRun)
		if err != nil {
			return fmt.Errorf("unable to write /etc/geolocation: %s", err)
		}

		return nil
	}

	location, err := Location()
	if err != nil {
		return fmt.Errorf("unable to find location: %s", err)
	}

	err = location.WriteGeolocation(config.DryRun)
	if err != nil {
		return fmt.Errorf("unable to write /etc/geolocation: %s", err)
	}

	return nil
}

// Find the lat/long entry for the current timezone
// in /usr/share/zoneinfo/zone.tab.
func Location() (*tz.Location, error) {
	tzName, err := tz.LocalTZ()
	if err != nil {
		return nil, err
	}

	if config.Verbose {
		fmt.Printf("Timezone: %s\n", tzName)
	}

	entry, err := tz.ZoneEntry(tzName, config.Verbose)
	if err != nil {
		return nil, err
	}

	if config.Verbose {
		fmt.Printf("Location: %#v\n", *entry)
	}

	return entry, nil
}

// Returns the home location values if provided,
func homeLocation() (*tz.Location, error) {
	if config.HomeLatitude != 0 && config.HomeLongitude != 0 {
		return &tz.Location{
			Latitude:  config.HomeLatitude,
			Longitude: config.HomeLongitude,
			Altitude:  config.HomeAltitude,
			Accuracy:  config.HomeAccuracy,
			Name:      "Home",
		}, nil
	} else {
		return nil, errors.New("home lat/long not provided")
	}
}
