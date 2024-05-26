package player

type PlayerRole string

const (
	BATTER        PlayerRole = "BATTER"
	BOWLER                   = "BOWLER"
	ALL_ROUNDER              = "ALL_ROUNDER"
	WICKET_KEEPER            = "WICKET_KEEPER"
)

type PlayerStats struct {
	Name          string         `csv:"name"`
	NameWithTitle string         `csv:"name_with_title"`
	Team          string         `csv:"team"`
	Score         int            `csv:"score"`
	PlayerRole    PlayerRole     `csv:"role"`
	BattingStats  *BattingStats  `csv:"batting"`
	BowlingStats  *BowlingStats  `csv:"bowling"`
	FieldingStats *FieldingStats `csv:"fielding"`
}

func (py *PlayerStats) CalculatePlayerScore() {
	py.Score = py.BattingStats.GetBattingScore() +
		py.BowlingStats.GetBowlingScore() +
		py.FieldingStats.GetFieldingScore()
}
