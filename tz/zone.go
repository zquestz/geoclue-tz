package tz

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

const zonetab = "/usr/share/zoneinfo/zone.tab"

func ZoneEntry(name string, dryRun bool) (*Location, error) {
	file, err := os.Open(zonetab)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		row := strings.Split(scanner.Text(), "\t")

		if len(row) >= 2 && row[2] == name {
			if dryRun {
				fmt.Printf("Zone Entry: %#v\n", row)
			}

			r, err := regexp.Compile(`([+|-]\d+)([+|-]\d+)`)
			if err != nil {
				return nil, err
			}

			matches := r.FindStringSubmatch(row[1])
			if len(matches) < 3 {
				return nil, fmt.Errorf("failed to parse coordinates: %q", row[1])
			}

			lat, err := convertCoordinates(matches[1], 2)
			if err != nil {
				return nil, err
			}

			long, err := convertCoordinates(matches[2], 3)
			if err != nil {
				return nil, err
			}

			return &Location{
				Latitude:  lat,
				Longitude: long,
				Altitude:  0,
				Accuracy:  1000,
				Name:      name,
			}, nil
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return nil, fmt.Errorf("time zone entry not found in %q", zonetab)
}

func convertCoordinates(coordinate string, insertIndex int) (float32, error) {
	coordBytes := []byte(strings.Trim(coordinate, "+"))
	if coordBytes[0] == '-' {
		insertIndex += 1
	}

	coordStr := fmt.Sprintf("%s%s%s", coordBytes[:insertIndex], ".", coordBytes[insertIndex:])

	coord, err := strconv.ParseFloat(coordStr, 32)
	if err != nil {
		return 0, err
	}

	return float32(coord), nil
}
