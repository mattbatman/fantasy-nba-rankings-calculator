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
	var fullPlayerRanks []models.BioStatsPoints
	var puntTOPlayerRanks []models.BioStatsPoints

	amountAndCategory = ByAmountAndCategory(data)
	fullPlayerRanks = DistributePlayerPoints(amountAndCategory, data)
	puntTOPlayerRanks = PuntTurnoversFiftyPercent(fullPlayerRanks)

	fullSortedPlayerRanks := SortRanks(fullPlayerRanks)
	puntTOSortedPlayerRanks := SortRanks(puntTOPlayerRanks)

	WriteRankings(fullSortedPlayerRanks, season, false)
	WriteRankings(puntTOSortedPlayerRanks, season, true)
}

// WriteRankings writes rankings data to file.
func WriteRankings(data []models.BioStatsPoints, season string, isPuntTO bool) {
	appConfig, configErr := config.GetConfig()

	if configErr != nil {
		fmt.Printf("Config File Error: %s\n", configErr)
		os.Exit(1)
	}

	datatype := appConfig.PerGameOrTotalFileIdentifier

	mi, _ := json.MarshalIndent(data, "", " ")
	outPath := "data/ranks-" + datatype + "-"

	if isPuntTO {
		outPath = outPath + "to-halved-"
	}

	outPath = outPath + season + ".json"

	_ = ioutil.WriteFile(outPath, mi, 0644)
}
