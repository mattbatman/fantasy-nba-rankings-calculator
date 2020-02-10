package rank

import (
	"github.com/mattbatman/fantasy-nba-rankings-calculator/models"
)

// GetDataBySeasons takes multi-season data from a Google Sheet and breaks it into
// a map with a key for each season
func GetDataBySeasons(data []models.Spreadsheet) map[string][]models.Spreadsheet {
	dataBySeason := make(map[string][]models.Spreadsheet)

	for _, row := range data {
		season := row.Season

		dataBySeason[season] = append(dataBySeason[season], row)
	}

	return dataBySeason
}
