package rank

import (
	"encoding/json"
	"io/ioutil"

	"github.com/mattbatman/fantasy-nba-rankings-calculator/models"
)

// CalculateRankings is the driver logic for the -mode="calcuate" flag.
// It reads a JSON file of player data, creates rankings, and writes a new JSON file.
func CalculateRankings() {
	file, _ := ioutil.ReadFile("data/sheets-data.json")

	var data []models.Spreadsheet
	var amountAndCategory models.AmountAndCategory
	var playerRanks []models.BioStatsPoints

	_ = json.Unmarshal(file, &data)

	amountAndCategory = ByAmountAndCategory(data)
	playerRanks = DistributePlayerPoints(amountAndCategory, data)

	WriteRankings(playerRanks)
}

// WriteRankings writes rankings data to file.
func WriteRankings(data []models.BioStatsPoints) {
	mi, _ := json.MarshalIndent(data, "", " ")

	_ = ioutil.WriteFile("data/ranks.json", mi, 0644)
}
