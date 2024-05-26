package player

type FieldingStats struct {
	Catches        int `csv:"catches"`
	Stumps         int `csv:"stumps"`
	DirectRunOut   int `csv:"direct_run_out"`
	IndirectRunOut int `csv:"indirect_run_out"`
}

func (fs *FieldingStats) GetFieldingScore() int {
	return fs.getCatchScore() + fs.getStumpingScore() + fs.getRunOutScore()
}

func (fs *FieldingStats) getCatchScore() int {
	catchPoints := fs.Catches * CATCH_SCORE
	if fs.Catches >= 3 {
		catchPoints += BONUS_THREE_CATCH
	}
	return catchPoints
}

func (fs *FieldingStats) getStumpingScore() int {
	return fs.Stumps * STUMP_SCORE
}

func (fs *FieldingStats) getRunOutScore() int {
	return (fs.DirectRunOut * RUNOUT_DIRECT_HIT) + (fs.IndirectRunOut * RUNOUT_INDIRECT_HIT)
}
