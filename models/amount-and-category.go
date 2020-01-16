package models

type AmountAndCategory struct {
	FieldGoalPercentage map[string]PointsDue
	FreeThrowPercentage map[string]PointsDue
	ThreePointsMade     map[string]PointsDue
	Points              map[string]PointsDue
	Rebounds            map[string]PointsDue
	Assists             map[string]PointsDue
	Steals              map[string]PointsDue
	Blocks              map[string]PointsDue
	Turnovers           map[string]PointsDue
}

type PointsDue struct {
	DuePer  float64
	DueList []int
}
