package rank

import (
	"encoding/json"
	"io/ioutil"
	"testing"

	"github.com/mattbatman/fantasy-nba-rankings-calculator/models"
)

func TestGetSortedListOfValues(t *testing.T) {
	file, e := ioutil.ReadFile("../test-data/sheets-data.json")

	if e != nil {
		t.Error(e)
	}

	var data []models.NBAData

	_ = json.Unmarshal(file, &data)

	actualPoints := GetSortedListOfValues(data, "Points")
	actualTurnovers := GetSortedListOfValues(data, "Turnovers")
	actualFGP := GetSortedListOfValues(data, "FieldGoalPercentage")

	expectedPoints := [3]float64{29.2, 30.1, 37.2}
	expectedTurnovers := [3]float64{4.8, 4.7, 3.4}
	expectedFGP := [3]float64{44.6, 44.6, 55.2}

	for i, v := range actualPoints {
		if v != expectedPoints[i] {
			t.Error("Points were not sorted correctly")
		}
	}

	for i, v := range actualTurnovers {
		if v != expectedTurnovers[i] {
			t.Error("Turnovers were not sorted correctly")
		}
	}

	for i, v := range actualFGP {
		if v != expectedFGP[i] {
			t.Error("Field goal percentage was not sorted correctly")
		}
	}
}

func TestCalculate(t *testing.T) {
	file, e := ioutil.ReadFile("../test-data/sheets-data.json")

	if e != nil {
		t.Error(e)
	}

	var data []models.NBAData

	_ = json.Unmarshal(file, &data)

	var actualPoints map[string]models.PointsDue
	var actualTurnovers map[string]models.PointsDue
	var actualFGP map[string]models.PointsDue

	actualPoints = Calculate(data, "Points")
	actualTurnovers = Calculate(data, "Turnovers")
	actualFGP = Calculate(data, "FieldGoalPercentage")

	if actualPoints["37.2"].DuePer != 3 {
		t.Error("The highest points were not awarded the most fantasy points")
	}

	if actualPoints["29.2"].DuePer != 1 {
		t.Error("The lowest points were not awarded the lowest fantasy points")
	}

	if actualTurnovers["3.4"].DuePer != 3 {
		t.Error("The lowest turnovers were not awarded the most fantasy points")
	}

	if actualTurnovers["4.8"].DuePer != 1 {
		t.Error("The lowest turnovers were not awarded the most fantasy points")
	}

	if actualFGP["55.2"].DuePer != 3 {
		t.Error("The highest field goal percentage was not awarded the most fantasy points")
	}

	if actualFGP["44.6"].DuePer != 1.5 {
		t.Error("The lowest field goal percentage was not awarded the least fantasy points")
	}
}
