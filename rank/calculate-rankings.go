package rank

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/mattbatman/fantasy-nba-rankings-calculator/config-parser"
	"github.com/mattbatman/fantasy-nba-rankings-calculator/models"
)

// CalculateRankings is the driver logic for the -mode="calcuate" flag.
// It reads a JSON file of player data, creates rankings, and writes a new JSON file.
// A separate file will be made for each season.
func CalculateRankings() {
	file, _ := ioutil.ReadFile("data/sheets-data.json")

	var data []models.Spreadsheet
	dataBySeasons := make(map[string][]models.Spreadsheet)

	_ = json.Unmarshal(file, &data)

	dataBySeasons = GetDataBySeasons(data)

	for season, seasonalData := range dataBySeasons {
		CalculateRankingsPerSeason(seasonalData, season)
	}
}

// CalculateRankingsPerSeason takes data for a single season, calculates rankings
// for that season, and writes the data to file.
func CalculateRankingsPerSeason(data []models.Spreadsheet, season string) {
	var amountAndCategory models.AmountAndCategory
	var playerRanks []models.BioStatsPoints

	amountAndCategory = ByAmountAndCategory(data)
	playerRanks = DistributePlayerPoints(amountAndCategory, data)

	sortedPlayerRanks := SortRanks(playerRanks)

	WriteRankings(sortedPlayerRanks, season)
}

// WriteRankings writes rankings data to file.
func WriteRankings(data []models.BioStatsPoints, season string) {
	appConfig, configErr := config.GetConfig()

	if configErr != nil {
		fmt.Printf("Config File Error: %s\n", configErr)
		os.Exit(1)
	}

	datatype := appConfig.PerGameOrTotalFileIdentifier

	mi, _ := json.MarshalIndent(data, "", " ")

	_ = ioutil.WriteFile("data/ranks-"+datatype+"-"+season+".json", mi, 0644)
}
