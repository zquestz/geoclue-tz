package cmd

import (
	"fmt"
	"os"
	"strings"
	"unicode"

	"github.com/spf13/cobra"
	"github.com/zquestz/geoclue-tz/tz"
)

const (
	appName = "geoclue-tz"
	version = "1.0.0"
)

// Stores configuration data.
var config Config

// GenerateCmd is the main command for the application.
var GenerateCmd = &cobra.Command{
	Use:   "geoclue-tz",
	Short: "Generate geoclue /etc/geolocation based on the current time zone.",
	Long:  `Generate geoclue /etc/geolocation based on the current time zone.`,
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
		bail(fmt.Errorf("failed to load configuration: %w", err))
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

func completion(cmd *cobra.Command, c string) {
	switch c {
	case "bash":
		err := cmd.GenBashCompletion(os.Stdout)
		if err != nil {
			bail(fmt.Errorf("failed to generate bash completion: %w", err))
		}
	case "zsh":
		if err := cmd.GenZshCompletion(os.Stdout); err != nil {
			bail(fmt.Errorf("failed to generate zsh completion: %w", err))
		}
	case "fish":
		if err := cmd.GenFishCompletion(os.Stdout, true); err != nil {
			bail(fmt.Errorf("failed to generate fish completion: %w", err))
		}
	default:
		bail(fmt.Errorf("completion not supported: %q", c))
	}
}

func prepareFlags() {
	GenerateCmd.PersistentFlags().BoolVar(
		&config.DisplayVersion, "version", false, "display version")
	GenerateCmd.PersistentFlags().BoolVarP(
		&config.DryRun, "dry-run", "d", config.DryRun, "dry run debug mode")
	GenerateCmd.PersistentFlags().StringVarP(
		&config.Location, "location", "l", config.Location, "enable custom location")
	GenerateCmd.PersistentFlags().StringVarP(
		&config.Completion, "completion", "c", config.Completion, "completion script for bash, zsh or fish")

	err := GenerateCmd.RegisterFlagCompletionFunc("completion", func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return []string{"bash", "fish", "zsh"}, cobra.ShellCompDirectiveNoFileComp
	})
	if err != nil {
		bail(err)
	}

	err = GenerateCmd.RegisterFlagCompletionFunc("location", func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return listLocations(), cobra.ShellCompDirectiveNoFileComp
	})
	if err != nil {
		bail(err)
	}
}

// Where all the work happens.
func performCommand(cmd *cobra.Command, args []string) error {
	if config.Completion != "" {
		completion(cmd, config.Completion)
		return nil
	}

	if config.DisplayVersion {
		fmt.Printf("%s %s\n", appName, version)
		return nil
	}

	if len(args) != 0 {
		help := cmd.HelpFunc()
		help(cmd, args)
	}

	if config.Location != "" {
		location, err := buildLocation(config.Location)
		if err != nil {
			return err
		}

		err = location.WriteGeolocation(config.DryRun)
		if err != nil {
			return fmt.Errorf("unable to write %q: %w", tz.EtcGeolocation, err)
		}

		return nil
	}

	location, err := Location()
	if err != nil {
		return fmt.Errorf("unable to find location: %w", err)
	}

	err = location.WriteGeolocation(config.DryRun)
	if err != nil {
		return fmt.Errorf("unable to write %q: %w", tz.EtcGeolocation, err)
	}

	return nil
}

// Find the lat/long entry for the current time zone
// in /usr/share/zoneinfo/zone.tab.
func Location() (*tz.Location, error) {
	tzName, err := tz.LocalTZ()
	if err != nil {
		return nil, err
	}

	if config.DryRun {
		fmt.Printf("Time Zone: %q\n", tzName)
	}

	entry, err := tz.ZoneEntry(tzName, config.DryRun)
	if err != nil {
		return nil, err
	}

	if config.DryRun {
		fmt.Printf("Location: %#v\n", *entry)
	}

	return entry, nil
}

// listLocations returns a list of custom locations.
func listLocations() []string {
	locations := []string{}
	for _, loc := range config.Locations {
		locations = append(locations, loc.Name)
	}

	return locations
}

// Returns the custom location,
func buildLocation(location string) (*tz.Location, error) {
	for _, loc := range config.Locations {
		if strings.ToLower(loc.Name) == strings.ToLower(location) {
			if loc.Latitude != 0 && loc.Longitude != 0 {
				l := &tz.Location{
					Latitude:  loc.Latitude,
					Longitude: loc.Longitude,
					Altitude:  loc.Altitude,
					Accuracy:  loc.Accuracy,
					Name:      loc.Name,
				}

				if config.DryRun {
					fmt.Printf("Location: %#v\n", *l)
				}

				return l, nil
			}

			return nil, fmt.Errorf("location lat/long not provided: %q", location)
		}
	}

	return nil, fmt.Errorf("location not found: %q", location)
}
