package player

type MatchPlayer struct {
	MatchID         string       `csv:"match-id"`
	MatchName       string       `csv:"match-name"`
	MatchResult     string       `csv:"match-result"`
	OnFieldUmpires1 string       `csv:"on-field-umpires1"`
	OnFieldUmpires2 string       `csv:"on-field-umpires2"`
	OnFieldUmpires3 string       `csv:"on-field-umpires3"`
	Venue           string       `csv:"venue"`
	PlayerStats     *PlayerStats `csv:"player"`
}
