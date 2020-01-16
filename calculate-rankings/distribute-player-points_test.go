package calculateRankings

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/mattbatman/fantasy-nba-rankings-calculator/models"
)

func TestDistributePlayerPoints(t *testing.T) {
	file, e := ioutil.ReadFile("../test-data/sheets-data.json")

	if e != nil {
		t.Error(e)
	}

	var data []models.NBAData
	var amountAndCategory models.AmountAndCategory
	var playerRanks []models.BioStatsPoints

	_ = json.Unmarshal(file, &data)

	amountAndCategory = ByAmountAndCategory(data)
	playerRanks = DistributePlayerPoints(amountAndCategory, data)

	if playerRanks[0].Points.FieldGoalPercentage != 1.5 {
		t.Error("Unexpected fantasy points for FG%")
	}

	if playerRanks[0].Points.FreeThrowPercentage != 3 {
		t.Error("Unexpected fantasy points for FT%")
	}

	if playerRanks[0].Points.Turnovers != 2 {
		t.Error("Unexpected fantasy points for turnovers")
	}

	if playerRanks[0].Points.Assists != 2 {
		t.Error("Unexpected fantasy points for assists")
	}

	fmt.Println(playerRanks[0].Points)
	if playerRanks[0].Points.Total != 21.5 {
		t.Error("Unexpected fantasy points for total")
	}
}
