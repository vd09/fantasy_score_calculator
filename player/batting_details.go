package player

type BattingStats struct {
	Runs        int
	TotalFours  int
	TotalSixes  int
	PlayedBalls int
	OutDetails  string
}

func (bt BattingStats) GetBattingScore() int {
	return bt.getRunScore() + bt.getBoundaryScore() + bt.getStrikeRateScore()
}

func (bt BattingStats) getRunScore() int {
	return bt.Runs + bt.getRunBonusScores()
}

func (bt BattingStats) getBoundaryScore() int {
	var boundScore int
	boundScore += bt.TotalFours * FOUR_BOUNDARY_BONUS
	boundScore += bt.TotalSixes * SIX_BOUNDARY_BONUS
	return boundScore
}

func (bt BattingStats) getRunBonusScores() int {
	switch {
	case bt.Runs >= 100:
		return CENTURY_BONUS
	case bt.Runs >= 50:
		return HALF_CENTURY_BONUS
	case bt.Runs >= 30:
		return THIRTY_RUN_BONUS
	case bt.Runs == 0 && bt.isOut():
		return DUCK_BONUS
	default:
		return 0
	}
}

func (bt BattingStats) getStrikeRateScore() int {
	if bt.PlayedBalls < 10 {
		return 0
	}

	strikeRate := float64(bt.Runs) / float64(bt.PlayedBalls)
	switch {
	case strikeRate > 1.7:
		return 6
	case strikeRate > 1.5:
		return 4
	case strikeRate >= 1.3:
		return 2
	case strikeRate > 0.7:
		return 0
	case strikeRate >= 0.6:
		return -2
	case strikeRate >= 0.5:
		return -4
	default:
		return -6
	}
}

func (bt BattingStats) isOut() bool {
	return len(bt.OutDetails) != 0
}
