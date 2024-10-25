package cmd

import "fmt"

// GeoClue stores location information.
type GeoClue struct {
	Latitude  float32
	Longitude float32
	Altitude  float32
	Accuracy  float32
}

// Output formats the Geoclue data\
// for /etc/geolocation.
func (g *GeoClue) Output() string {
	return fmt.Sprintf(
		"%f\n%f\n%f\n%f",
		g.Latitude,
		g.Longitude,
		g.Altitude,
		g.Accuracy,
	)
}
