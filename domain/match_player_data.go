package domain

type MatchPlayerData struct {
	MatchID          string
	Team             string
	Name             string
	FoursAsBatter    int
	SixesAsBatter    int
	BallsAsBatter    int
	DotBallsAsBowler int
	MatchName        string
	MatchResult      string
	NameWDTitles     string
	OnFieldUmpires1  string
	OnFieldUmpires2  string
	OnFieldUmpires3  string
	OutDetails       string
	OversAsBowler    float64
	PlayerRole       string
	RunsAsBatter     int
	RunsAsBowler     int
	Venue            string
	WicketAsBowler   int
}
