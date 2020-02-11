package rank

import "github.com/mattbatman/fantasy-nba-rankings-calculator/models"

func PuntTurnoversFiftyPercent(data []models.BioStatsPoints) []models.BioStatsPoints {
	var turnoversHalved []models.BioStatsPoints

	for _, fullPlayer := range data {
		halfPlayer := fullPlayer

		halfPlayer.Points.Turnovers = halfPlayer.Points.Turnovers / 2
		halfPlayer.Points.Total = SumFantasyPoints(halfPlayer.Points)

		turnoversHalved = append(turnoversHalved, halfPlayer)
	}

	return turnoversHalved
}
