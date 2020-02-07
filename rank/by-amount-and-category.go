package rank

import (
	"fmt"
	"sort"
	"strconv"

	"github.com/mattbatman/fantasy-nba-rankings-calculator/models"
)

// ByAmountAndCategory calls Calculate for every field in the
// AmountAndCategory struct
func ByAmountAndCategory(data []models.Spreadsheet) models.AmountAndCategory {
	var amountAndCategory models.AmountAndCategory

	amountAndCategory.Points = Calculate(data, "Points")
	amountAndCategory.FieldGoalPercentage = Calculate(data, "FieldGoalPercentage")
	amountAndCategory.FreeThrowPercentage = Calculate(data, "FreeThrowPercentage")
	amountAndCategory.ThreePointsMade = Calculate(data, "ThreePointsMade")
	amountAndCategory.Rebounds = Calculate(data, "Rebounds")
	amountAndCategory.Assists = Calculate(data, "Assists")
	amountAndCategory.Steals = Calculate(data, "Steals")
	amountAndCategory.Blocks = Calculate(data, "Blocks")
	amountAndCategory.Turnovers = Calculate(data, "Turnovers")

	return amountAndCategory
}

// Calculate takes data and a category and returns the points due as a slice for
// each amount. The statistical amount as a string is used as a key, pointing to the
// DuePer and DueList.
func Calculate(data []models.Spreadsheet, category string) map[string]models.PointsDue {
	pointsDue := make(map[string]models.PointsDue)

	slv := GetSortedListOfValues(data, category)

	for i, s := range slv {
		k := strconv.FormatFloat(s, 'f', -1, 64)

		if v, ok := pointsDue[k]; ok {
			v.DueList = append(v.DueList, i+1)

			sum := 0.0
			for _, dl := range v.DueList {
				sum += float64(dl)
			}
			average := sum / float64(len(v.DueList))

			v.DuePer = average

			pointsDue[k] = v
		} else {
			var pd models.PointsDue
			pd.DuePer = float64(i + 1)
			pd.DueList = append(pd.DueList, i+1)
			pointsDue[k] = pd
		}
	}

	return pointsDue
}

// GetSortedListOfValues creates a list of values for the provided category
func GetSortedListOfValues(data []models.Spreadsheet, category string) []float64 {
	var listOfValues []float64

	switch category {
	case "Assists":
		for _, d := range data {
			listOfValues = append(listOfValues, d.Assists)
		}
	case "Blocks":
		for _, d := range data {
			listOfValues = append(listOfValues, d.Blocks)
		}
	case "FieldGoalPercentage":
		for _, d := range data {
			listOfValues = append(listOfValues, d.FieldGoalPercentage)
		}
	case "FreeThrowPercentage":
		for _, d := range data {
			listOfValues = append(listOfValues, d.FreeThrowPercentage)
		}
	case "Points":
		for _, d := range data {
			listOfValues = append(listOfValues, d.Points)
		}
	case "Rebounds":
		for _, d := range data {
			listOfValues = append(listOfValues, d.Rebounds)
		}
	case "Steals":
		for _, d := range data {
			listOfValues = append(listOfValues, d.Steals)
		}
	case "ThreePointsMade":
		for _, d := range data {
			listOfValues = append(listOfValues, d.ThreePointsMade)
		}
	case "Turnovers":
		for _, d := range data {
			listOfValues = append(listOfValues, d.Turnovers)
		}
	default:
		fmt.Println("Unexpected category provided.")
	}

	switch category {
	case "Turnovers":
		sort.Sort(sort.Reverse(sort.Float64Slice(listOfValues)))
	default:
		sort.Float64s(listOfValues)
	}

	return listOfValues
}
