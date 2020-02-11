package rank

import (
	"encoding/json"
	"io/ioutil"
	"testing"

	"github.com/mattbatman/fantasy-nba-rankings-calculator/models"
)

func TestPuntTurnoversFiftyPercent(t *testing.T) {
	file, e := ioutil.ReadFile("../test-data/sheets-data.json")

	if e != nil {
		t.Error(e)
	}

	var data []models.Spreadsheet
	var amountAndCategory models.AmountAndCategory
	var playerRanks []models.BioStatsPoints
	var discountedTurnoverRanks []models.BioStatsPoints

	_ = json.Unmarshal(file, &data)

	amountAndCategory = ByAmountAndCategory(data)
	playerRanks = DistributePlayerPoints(amountAndCategory, data)
	discountedTurnoverRanks = PuntTurnoversFiftyPercent(playerRanks)

	if playerRanks[0].Points.Turnovers/2 != discountedTurnoverRanks[0].Points.Turnovers {
		t.Error("Turnovers not reduced by 50%")
	}

	if playerRanks[1].Points.Turnovers/2 != discountedTurnoverRanks[1].Points.Turnovers {
		t.Error("Turnovers not reduced by 50%")
	}

	if playerRanks[2].Points.Turnovers/2 != discountedTurnoverRanks[2].Points.Turnovers {
		t.Error("Turnovers not reduced by 50%")
	}

	getExpectedTotal := func(player models.BioStatsPoints) float64 {
		expectedTotal := player.Points.Total - player.Points.Turnovers/2

		return expectedTotal
	}

	expectedFirstTotal := getExpectedTotal(playerRanks[0])
	expectedSecondTotal := getExpectedTotal(playerRanks[1])
	expectedThirdTotal := getExpectedTotal(playerRanks[2])

	if discountedTurnoverRanks[0].Points.Total != expectedFirstTotal {
		t.Error("Total not calculated correctly")
	}

	if discountedTurnoverRanks[1].Points.Total != expectedSecondTotal {
		t.Error("Total not calculated correctly")
	}

	if discountedTurnoverRanks[2].Points.Total != expectedThirdTotal {
		t.Error("Total not calculated correctly")
	}
}
