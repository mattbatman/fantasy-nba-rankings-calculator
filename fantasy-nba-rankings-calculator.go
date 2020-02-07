package main

import (
	"flag"
	"fmt"

	"github.com/mattbatman/fantasy-nba-rankings-calculator/models"
	"github.com/mattbatman/fantasy-nba-rankings-calculator/rank"
	"github.com/mattbatman/fantasy-nba-rankings-calculator/sheets"
)

// readCommandLineParams parses command line parameters, e.g. -mode="fetch"
func readCommandLineParams() models.Parameters {
	var mode = flag.String("mode", "help", "Can be fetch or calculate")

	flag.Parse()

	var params models.Parameters

	if *mode != "help" {
		params.Mode = *mode
	}

	return params
}

// dispatch will trigger a path of logic to either fetch data from google sheets
// or calculate fantasy nba rankings and write them as a JSON file
func dispatch(parameters *models.Parameters) {
	switch parameters.Mode {
	case "fetch":
		sheets.FetchSheetsData()
	case "calculate":
		rank.CalculateRankings()
	default:
		fmt.Println("No mode recognized.")
	}
}

// main is the function that will run from executing the binary
func main() {
	parameters := readCommandLineParams()
	dispatch(&parameters)
}
