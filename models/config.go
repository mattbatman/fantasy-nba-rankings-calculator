package models

// Config is the structure of a config JSON file to be used
// by logic fetching data from google sheets
type Config struct {
	SpreadsheetID                string `json:"spreadsheetID"`
	SpreadsheetRange             string `json:"spreadsheetRange"`
	PerGameOrTotalFileIdentifier string `json:"perGameOrTotalFileIdentifier"`
}
