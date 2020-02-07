package models

type BioStatsPoints struct {
	Bio    Bio    `json:"bio"`
	Stats  Stats  `json:"stats"`
	Points Points `json:"points"`
}

type Bio struct {
	Player      string  `json:"PLAYER"`
	Team        string  `json:"TEAM"`
	Age         int     `json:"AGE,string"`
	GamesPlayed int     `json:"GP,string"`
	Wins        int     `json:"W,string"`
	Losses      int     `json:"L,string"`
	Minutes     float64 `json:"MIN,string"`
}

type Stats struct {
	FieldGoalPercentage float64 `json:"FG%,string"`
	FreeThrowPercentage float64 `json:"FT%,string"`
	ThreePointsMade     float64 `json:"3PM,string"`
	Points              float64 `json:"PTS,string"`
	Rebounds            float64 `json:"REB,string"`
	Assists             float64 `json:"AST,string"`
	Steals              float64 `json:"STL,string"`
	Blocks              float64 `json:"BLK,string"`
	Turnovers           float64 `json:"TOV,string"`
}

type Points struct {
	FieldGoalPercentage float64 `json:"FG%,string"`
	FreeThrowPercentage float64 `json:"FT%,string"`
	ThreePointsMade     float64 `json:"3PM,string"`
	Points              float64 `json:"PTS,string"`
	Rebounds            float64 `json:"REB,string"`
	Assists             float64 `json:"AST,string"`
	Steals              float64 `json:"STL,string"`
	Blocks              float64 `json:"BLK,string"`
	Turnovers           float64 `json:"TOV,string"`
	Total               float64 `json:"TOT,string"`
}
