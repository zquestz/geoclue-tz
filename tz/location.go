package tz

import (
	"errors"
	"fmt"
	"os"
	"os/user"
	"strconv"
)

const (
	etcGeolocation = "/etc/geolocation"
)

// Location stores location information.
type Location struct {
	Latitude  float32 `json:"latitude,string"`
	Longitude float32 `json:"longitude,string"`
	Altitude  float32 `json:"altitude,string"`
	Accuracy  float32 `json:"accuracy,string"`
	Name      string  `json:"name"`
}

func (l *Location) WriteGeolocation(dryRun bool) error {
	if dryRun {
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

	err = os.WriteFile(etcGeolocation, []byte(l.Output()), 0600)
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

	fmt.Printf("Successfully updated %s with %s location\n", etcGeolocation, l.Name)

	return nil
}

// Output formats the Location
// for /etc/geolocation.
func (l *Location) Output() string {
	return fmt.Sprintf(
		"%v\n%v\n%v\n%v",
		l.Latitude,
		l.Longitude,
		l.Altitude,
		l.Accuracy,
	)
}
