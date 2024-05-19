package player

type PlayerRole int

const (
	BATTER PlayerRole = iota
	BOWLER
	ALL_ROUNDER
	WICKET_KEEPER
)

type Player struct {
	Name          string
	NameWithTitle string
	PlayerRole
	BattingStats
	BowlingStats
	FieldingStats
}

func (py Player) GetPlayerScore() int {
	return py.BattingStats.GetBattingScore() +
		py.BowlingStats.GetBowlingScore() +
		py.FieldingStats.GetFieldingScore()
}
