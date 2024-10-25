package cmd

import (
	"encoding/json"
	"os"
	"path/filepath"

	"github.com/zquestz/geoclue-tz/tz"
	"github.com/zquestz/go-ucl"
)

// Config stores all the application configuration.
type Config struct {
	Locations      []*tz.Location `json:"locations"`
	DisplayVersion bool           `json:"-"`
	Location       string         `json:"location"`
	DryRun         bool           `json:"dryRun,string"`
	Completion     string         `json:"completion"`
}

// Load reads the configuration from /etc/geoclue-tz.conf
// and loads it into the Config struct.
// The config is in UCL format.
func (c *Config) Load() error {
	conf, err := c.loadConfig()
	if err != nil {
		return err
	}

	// Conf is not required.
	if conf != nil {
		err = c.applyConf(conf)
		if err != nil {
			return err
		}
	}

	return nil
}

func (c *Config) loadConfig() ([]byte, error) {
	f, err := os.Open(filepath.Join("/etc", "geoclue-tz.conf"))
	if err != nil {
		if os.IsNotExist(err) {
			return nil, nil
		} else {
			return nil, err
		}
	}
	defer f.Close()

	ucl.Ucldebug = false
	data, err := ucl.NewParser(f).Ucl()
	if err != nil {
		return nil, err
	}

	conf, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	return conf, nil
}

func (c *Config) applyConf(conf []byte) error {
	err := json.Unmarshal(conf, c)
	if err != nil {
		return err
	}

	return nil
}
