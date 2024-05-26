package domain

type MatchPlayerData struct {
	MatchID          string  `csv:"match-id"`
	Team             string  `csv:"team"`
	Name             string  `csv:"name"`
	FoursAsBatter    int     `csv:"4s-as-batter"`
	SixesAsBatter    int     `csv:"6s-as-batter"`
	BallsAsBatter    int     `csv:"balls-as-batter"`
	DotBallsAsBowler int     `csv:"dotballs-as-bowler"`
	MatchName        string  `csv:"match-name"`
	MatchResult      string  `csv:"match-result"`
	NameWDTitles     string  `csv:"name-wd-titles"`
	OnFieldUmpires1  string  `csv:"on-field-umpires1"`
	OnFieldUmpires2  string  `csv:"on-field-umpires2"`
	OnFieldUmpires3  string  `csv:"on-field-umpires3"`
	OutDetails       string  `csv:"out-details"`
	OversAsBowler    float64 `csv:"overs-as-bowler"`
	PlayerRole       string  `csv:"player-role"`
	RunsAsBatter     int     `csv:"runs-as-batter"`
	RunsAsBowler     int     `csv:"runs-as-bowler"`
	Venue            string  `csv:"venue"`
	WicketAsBowler   int     `csv:"wicket-as-bowler"`
}
