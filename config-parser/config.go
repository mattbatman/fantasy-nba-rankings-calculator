package config

import (
	"encoding/json"
	"errors"
	"os"

	"github.com/mattbatman/fantasy-nba-rankings-calculator/models"
)

// ReadConfig will attempt to decode and return the JSON file at the
// specified path. If this can't be done, it will return an error and
// an empty, default config.
func ReadConfig(path string) (models.Config, error) {
	c := models.Config{
		SpreadsheetID:    "",
		SpreadsheetRange: "",
	}

	cfile, err := os.Open(path)

	if err != nil {
		return c, err
	}

	jsonDecoder := json.NewDecoder(cfile)

	err = jsonDecoder.Decode(&c)

	if err != nil {
		return c, err
	}

	return c, nil
}

// GetConfig retrieves and returns the config JSON file as a struct
func GetConfig() (models.Config, error) {
	config, err := ReadConfig("conf/config.json")

	if err != nil {
		return config, errors.New("Invalid config file")
	}

	return config, nil
}
