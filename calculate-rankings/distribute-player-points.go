package calculateRankings

import (
	"strconv"

	"github.com/mattbatman/fantasy-nba-rankings-calculator/models"
)

// DistributePlayerPoints takes the fantasy points due for each category at
// each amount, iterates over the statistical NBA data, and returns a data
// structure that will eventually be written as a JSON file
func DistributePlayerPoints(amountAndCategory models.AmountAndCategory, NBAData []models.NBAData) []models.BioStatsPoints {
	var outputData []models.BioStatsPoints

	for _, nd := range NBAData {
		var playerData models.BioStatsPoints

		playerData.Bio.Player = nd.Player
		playerData.Bio.Team = nd.Team
		playerData.Bio.Age = nd.Age
		playerData.Bio.GamesPlayed = nd.GamesPlayed
		playerData.Bio.Wins = nd.Wins
		playerData.Bio.Losses = nd.Losses
		playerData.Bio.Minutes = nd.Minutes

		playerData.Points.FieldGoalPercentage = amountAndCategory.FieldGoalPercentage[strconv.FormatFloat(nd.FieldGoalPercentage, 'f', -1, 64)].DuePer
		playerData.Points.FreeThrowPercentage = amountAndCategory.FreeThrowPercentage[strconv.FormatFloat(nd.FreeThrowPercentage, 'f', -1, 64)].DuePer
		playerData.Points.ThreePointsMade = amountAndCategory.ThreePointsMade[strconv.FormatFloat(nd.ThreePointsMade, 'f', -1, 64)].DuePer
		playerData.Points.Points = amountAndCategory.Points[strconv.FormatFloat(nd.Points, 'f', -1, 64)].DuePer
		playerData.Points.Rebounds = amountAndCategory.Rebounds[strconv.FormatFloat(nd.Rebounds, 'f', -1, 64)].DuePer
		playerData.Points.Assists = amountAndCategory.Assists[strconv.FormatFloat(nd.Assists, 'f', -1, 64)].DuePer
		playerData.Points.Blocks = amountAndCategory.Blocks[strconv.FormatFloat(nd.Blocks, 'f', -1, 64)].DuePer
		playerData.Points.Steals = amountAndCategory.Steals[strconv.FormatFloat(nd.Steals, 'f', -1, 64)].DuePer
		playerData.Points.Turnovers = amountAndCategory.Turnovers[strconv.FormatFloat(nd.Turnovers, 'f', -1, 64)].DuePer

		playerData.Points.Total = SumFantasyPoints(playerData.Points)

		playerData.Stats.FieldGoalPercentage = nd.FieldGoalPercentage
		playerData.Stats.FreeThrowPercentage = nd.FreeThrowPercentage
		playerData.Stats.ThreePointsMade = nd.ThreePointsMade
		playerData.Stats.Points = nd.Points
		playerData.Stats.Rebounds = nd.Rebounds
		playerData.Stats.Assists = nd.Assists
		playerData.Stats.Blocks = nd.Blocks
		playerData.Stats.Steals = nd.Steals
		playerData.Stats.Turnovers = nd.Turnovers

		outputData = append(outputData, playerData)
	}

	return outputData
}

// SumFantasyPoints returns the total fantasy points due
func SumFantasyPoints(cats models.Points) float64 {
	t := 0.0

	t += cats.FieldGoalPercentage
	t += cats.FreeThrowPercentage
	t += cats.ThreePointsMade
	t += cats.Points
	t += cats.Rebounds
	t += cats.Assists
	t += cats.Blocks
	t += cats.Steals
	t += cats.Turnovers

	return t
}
