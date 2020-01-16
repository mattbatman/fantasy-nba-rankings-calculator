package models

type NBAData struct {
	SortAtPull            int     `json:"SORT_AT_PULL,string"`
	Player                string  `json:"PLAYER"`
	Team                  string  `json:"TEAM"`
	Age                   int     `json:"AGE,string"`
	GamesPlayed           int     `json:"GP,string"`
	Wins                  int     `json:"W,string"`
	Losses                int     `json:"L,string"`
	Minutes               float64 `json:"MIN,string"`
	Points                float64 `json:"PTS,string"`
	FieldGoalsMade        float64 `json:"FGM,string"`
	FieldGoalsAttempted   float64 `json:"FGA,string"`
	FieldGoalPercentage   float64 `json:"FG%,string"`
	ThreePointsMade       float64 `json:"3PM,string"`
	ThreePointsAttempted  float64 `json:"3PA,string"`
	ThreePointsPercentage float64 `json:"3P%,string"`
	FreeThrowsMade        float64 `json:"FTM,string"`
	FreeThrowsAttempted   float64 `json:"FTA,string"`
	FreeThrowPercentage   float64 `json:"FT%,string"`
	OffensiveRebounds     float64 `json:"OREB,string"`
	DefensiveRebounds     float64 `json:"DREB,string"`
	Rebounds              float64 `json:"REB,string"`
	Assists               float64 `json:"AST,string"`
	Turnovers             float64 `json:"TOV,string"`
	Steals                float64 `json:"STL,string"`
	Blocks                float64 `json:"BLK,string"`
	PersonalFouls         float64 `json:"PF,string"`
	FantasyPoints         float64 `json:"FP,string"`
	DoubleDoubles         float64 `json:"DD2,string"`
	TripleDoubles         float64 `json:"TD3,string"`
	PlusMinus             float64 `json:"+/-,string"`
	Season                string  `json:"SEASON"`
}

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
