package rank

import (
	"encoding/json"
	"io/ioutil"
	"testing"

	"github.com/mattbatman/fantasy-nba-rankings-calculator/models"
)

func TestGetDataBySeasons(t *testing.T) {
	file, e := ioutil.ReadFile("../test-data/sheets-data.json")

	if e != nil {
		t.Error(e)
	}

	var data []models.Spreadsheet
	actual := make(map[string][]models.Spreadsheet)

	_ = json.Unmarshal(file, &data)

	for key := range actual {
		if key != "2019-2020" {
			t.Error("Unexpected season as key")
		}
	}
}
