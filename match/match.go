package match

import (
	"fantasy_score_calculator/player"
	"strings"
)

type Match struct {
	Id                        string
	roleToNamesMap            map[player.PlayerRole][]string
	MatchPlayerMap            map[string]*player.MatchPlayer
	playerNameToFieldingStats map[string]*player.FieldingStats
	playerNameToLbwWickets    map[string]int
}

func (m *Match) UpdateScores() {
	for name, pl := range m.MatchPlayerMap {
		pl.PlayerStats.FieldingStats = m.getFieldingStatsFromPlayerName(name)
		pl.PlayerStats.BowlingStats.LbwWickets = m.getLbwWicketsFromPlayerName(name)
		pl.PlayerStats.CalculatePlayerScore()
	}
}

func (m *Match) AddPlayer(data *player.MatchPlayer) {
	m.MatchPlayerMap[data.PlayerStats.Name] = data
	m.UpdateFieldingStats(data.PlayerStats.BattingStats.OutDetails)
}

func (m *Match) UpdateFieldingStats(outDetails string) {
	wicketInfo := player.ParseWicketInfo(outDetails)
	if wicketInfo.LBW {
		m.playerNameToLbwWickets[wicketInfo.Bowler]++
	}

	if wicketInfo.Stumped != "" {
		m.getFieldingStatsFromPlayerName(wicketInfo.Stumped).Stumps++
	}
	if wicketInfo.Caught != "" {
		m.getFieldingStatsFromPlayerName(wicketInfo.Caught).Catches++
	}
	if wicketInfo.RunOutDirect != "" {
		if wicketInfo.RunOutIndirect != "" {
			m.getFieldingStatsFromPlayerName(wicketInfo.RunOutDirect).IndirectRunOut++
		} else {
			m.getFieldingStatsFromPlayerName(wicketInfo.RunOutDirect).DirectRunOut++
		}
	}
}

func (m *Match) getFieldingStatsFromPlayerName(playerName string) *player.FieldingStats {
	playerNameKey := m.getKeyFromPlayerName(playerName)
	if _, ok := m.playerNameToFieldingStats[playerNameKey]; !ok {
		m.playerNameToFieldingStats[playerNameKey] = &player.FieldingStats{}
	}
	return m.playerNameToFieldingStats[playerNameKey]
}

func (m *Match) getLbwWicketsFromPlayerName(playerName string) int {
	playerNameKey := m.getKeyFromPlayerName(playerName)
	if _, ok := m.playerNameToLbwWickets[playerNameKey]; !ok {
		m.playerNameToLbwWickets[playerNameKey] = 0
	}
	return m.playerNameToLbwWickets[playerNameKey]
}

func (m *Match) getKeyFromPlayerName(name string) string {
	name = strings.ReplaceAll(name, " ", "")
	return strings.ToUpper(name)
	//return name
}

func NewMatchWithFirstPlayer(data *player.MatchPlayer) *Match {
	matchDetails := NewMatch(data)
	matchDetails.AddPlayer(data)
	return matchDetails
}

func NewMatch(matchPlayerData *player.MatchPlayer) *Match {
	return &Match{
		Id:                        matchPlayerData.MatchID,
		MatchPlayerMap:            make(map[string]*player.MatchPlayer),
		playerNameToFieldingStats: make(map[string]*player.FieldingStats),
		playerNameToLbwWickets:    make(map[string]int),
	}
}
