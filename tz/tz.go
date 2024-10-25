package tz

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

const localtime = "/etc/localtime"

// LocalTZ returns the current IANA time zone of the system.
// Only works on Unix systems.
func LocalTZ() (string, error) {
	fi, err := os.Lstat(localtime)
	if err != nil {
		err = fmt.Errorf("failed to stat %q: %w", localtime, err)
		return "", err
	}

	if (fi.Mode() & os.ModeSymlink) == 0 {
		err = fmt.Errorf("%q is not a symlink", localtime)
		return "", err
	}

	p, err := os.Readlink(localtime)
	if err != nil {
		return "", err
	}

	name, err := inferTZFromPath(p)
	if err != nil {
		return "", err
	}

	return name, nil
}

func inferTZFromPath(p string) (string, error) {
	parts := strings.Split(p, string(filepath.Separator))
	for i := range parts {
		if parts[i] == "zoneinfo" {
			parts = parts[i+1:]
			break
		}
	}

	if len(parts) < 1 {
		return "", fmt.Errorf("unable to infer time zone from path: %q", p)
	}

	return filepath.Join(parts...), nil
}
