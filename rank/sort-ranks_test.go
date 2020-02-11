package rank

import (
	"encoding/json"
	"io/ioutil"
	"testing"

	"github.com/mattbatman/fantasy-nba-rankings-calculator/models"
)

func TestSortRanks(t *testing.T) {
	file, e := ioutil.ReadFile("../test-data/sheets-data.json")

	if e != nil {
		t.Error(e)
	}

	var data []models.Spreadsheet
	var amountAndCategory models.AmountAndCategory
	var playerRanks []models.BioStatsPoints

	_ = json.Unmarshal(file, &data)

	amountAndCategory = ByAmountAndCategory(data)
	playerRanks = DistributePlayerPoints(amountAndCategory, data)
	sortedRanks := SortRanks(playerRanks)

	if sortedRanks[0].Points.Total < sortedRanks[1].Points.Total {
		t.Error("First point total not greater than second")
	}

	if sortedRanks[1].Points.Total < sortedRanks[2].Points.Total {
		t.Error("Second point total not greater than last")
	}
}
