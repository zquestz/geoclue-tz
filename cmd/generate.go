package cmd

import (
	"errors"
	"fmt"
	"os"
	"os/user"
	"strconv"
	"unicode"

	"github.com/spf13/cobra"
)

const (
	appName        = "geoclue-tz"
	version        = "0.0.1"
	etcGeolocation = "/etc/geolocation"
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
		&config.DefaultLatitude, "default-latitude", "", config.DefaultLatitude, "default latitude")
	GenerateCmd.PersistentFlags().Float32VarP(
		&config.DefaultLongitude, "default-longitude", "", config.DefaultLongitude, "default longitude")
	GenerateCmd.PersistentFlags().Float32VarP(
		&config.DefaultAltitude, "default-altitude", "", config.DefaultAltitude, "default altitude")
	GenerateCmd.PersistentFlags().Float32VarP(
		&config.DefaultAccuracy, "default-accuracy", "", config.DefaultAccuracy, "default accuracy")
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

	entry, err := findTimezoneLatLong()
	if err != nil {
		return fmt.Errorf("unable to find location: %s", err)
	}

	err = writeGeolocation(entry)
	if err != nil {
		return fmt.Errorf("unable to write /etc/geolocation: %s", err)
	}

	return nil
}

// Find the lat/long entry for the current timezone
// in /usr/share/zoneinfo/zone.tab.
func findTimezoneLatLong() (*GeoClue, error) {
	tz, err := LocalTZ()
	if err != nil {
		return defaultGeoClue(err)
	}

	if config.Verbose {
		fmt.Printf("Timezone: %s\n", tz)
	}

	entry, err := zoneEntry(tz)
	if err != nil {
		return defaultGeoClue(err)
	}

	if config.Verbose {
		fmt.Printf("GeoClue: %#v\n", *entry)
	}

	return entry, nil
}

func writeGeolocation(g *GeoClue) error {
	if config.DryRun {
		return nil
	}

	currentUser, err := user.Current()
	if err != nil {
		return fmt.Errorf("unable to get current user: %s", err)
	}

	if currentUser.Uid != "0" {
		return errors.New("root access required")
	}

	geoclueUser, err := user.Lookup("geoclue")
	if err != nil {
		return err
	}

	err = os.WriteFile(etcGeolocation, []byte(g.Output()), 0600)
	if err != nil {
		return err
	}

	geoclueUserId, err := strconv.ParseInt(geoclueUser.Uid, 10, 0)
	if err != nil {
		return err
	}

	err = os.Chown(etcGeolocation, int(geoclueUserId), 0)
	if err != nil {
		return err
	}

	return nil
}

// Returns the default values if provided,
// otherwise just bubbles up the error.
func defaultGeoClue(err error) (*GeoClue, error) {
	if config.DefaultLatitude != 0 && config.DefaultLongitude != 0 {
		if config.Verbose {
			fmt.Printf("Error: defaults returned: %s\n", err)
		}

		return &GeoClue{
			Latitude:  config.DefaultLatitude,
			Longitude: config.DefaultLongitude,
			Altitude:  config.DefaultAltitude,
			Accuracy:  config.DefaultAccuracy,
		}, nil
	} else {
		return nil, err
	}
}
