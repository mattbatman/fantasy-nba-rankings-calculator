package rank

import (
	"sort"

	"github.com/mattbatman/fantasy-nba-rankings-calculator/models"
)

// SortRanks sorts the passed rankings in descending order of the total fantasy points
func SortRanks(data []models.BioStatsPoints) []models.BioStatsPoints {
	sort.Slice(data, func(i, j int) bool {
		return data[i].Points.Total > data[j].Points.Total
	})

	return data
}
