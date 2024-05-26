package player

type BowlingStats struct {
	Runs        int     `csv:"runs"`
	DotBalls    int     `csv:"dot_balls"`
	Wickets     int     `csv:"wickets"`
	LbwWickets  int     `csv:"lbw_wickets"`
	MaidenOvers int     `csv:"maiden_overs"`
	Overs       float64 `csv:"overs"`
}

func (bw *BowlingStats) GetBowlingScore() int {
	return bw.getWicketScore() + bw.getLbwBonusScore() + bw.getMaidenScore() + bw.getEconomyRateScore()
}

func (bw *BowlingStats) getWicketScore() int {
	return (bw.Wickets * WICKET_SCORE) + bw.getWicketBonus()
}

func (bw *BowlingStats) getLbwBonusScore() int {
	return bw.LbwWickets * LBW_BONUS
}

func (bw *BowlingStats) getWicketBonus() int {
	switch {
	case bw.Wickets >= 5:
		return FIVE_WICKET_BONUS
	case bw.Wickets >= 4:
		return FOUR_WICKET_BONUS
	case bw.Wickets >= 3:
		return THREE_WICKET_BONUS
	default:
		return 0
	}
}

func (bw *BowlingStats) getMaidenScore() int {
	return bw.MaidenOvers * MAIDEN_OVER_SCORE
}

func (bw *BowlingStats) getEconomyRateScore() int {
	if bw.Overs < 2 {
		return 0
	}

	er := bw.getEconomyRate()
	switch {
	case er < 5:
		return 6
	case er < 6:
		return 4
	case er <= 7:
		return 2
	case er < 10:
		return 0
	case er <= 11:
		return -2
	case er <= 12:
		return -4
	default:
		return -6
	}
}

func (bw *BowlingStats) getEconomyRate() float64 {
	lastFullOver := float64(int(bw.Overs))
	balls := (bw.Overs - lastFullOver) * 10
	adjustedOver := lastFullOver + (balls / 6)
	return float64(bw.Runs) / adjustedOver
}
