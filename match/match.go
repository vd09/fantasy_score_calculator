package match

import (
	"fantasy_score_calculator/domain"
	"fantasy_score_calculator/player"
)

type Match struct {
	id                        string
	name                      string
	roleToNameMap             map[player.PlayerRole]string
	playerMap                 map[string]player.Player
	playerNameToFieldingStats map[string]*player.FieldingStats
	playerNameToLbwWickets    map[string]int
}

func (m *Match) AddPlayer(data domain.MatchPlayerData) {
	m.playerMap[data.Name] = player.Player{
		Name:          data.Name,
		NameWithTitle: data.NameWDTitles,
		PlayerRole:    player.ALL_ROUNDER,
		BattingStats:  m.CreateBattingStatsForPlayer(data),
		BowlingStats:  m.CreateBowlingStatsForPlayer(data),
	}
	m.UpdateFieldingStats(data.OutDetails)
}

func (m *Match) CreateBattingStatsForPlayer(data domain.MatchPlayerData) player.BattingStats {
	return player.BattingStats{
		Runs:        data.RunsAsBatter,
		TotalFours:  data.FoursAsBatter,
		TotalSixes:  data.SixesAsBatter,
		PlayedBalls: data.BallsAsBatter,
		OutDetails:  data.OutDetails,
	}
}

func (m *Match) CreateBowlingStatsForPlayer(data domain.MatchPlayerData) player.BowlingStats {
	return player.BowlingStats{
		Runs:        data.RunsAsBowler,
		DotBalls:    data.DotBallsAsBowler,
		Wickets:     data.WicketAsBowler,
		LbwWickets:  0,
		MaidenOvers: 0,
		Overs:       data.OversAsBowler,
	}
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
	if _, ok := m.playerNameToFieldingStats[playerName]; !ok {
		m.playerNameToFieldingStats[playerName] = &player.FieldingStats{}
	}
	return m.playerNameToFieldingStats[playerName]
}

func NewMatchWithFirstPlayer(data domain.MatchPlayerData) *Match {
	matchDetails := NewMatchDetails(data.MatchID, data.MatchName)
	matchDetails.AddPlayer(data)
	return matchDetails
}

func NewMatchDetails(matchId, matchName string) *Match {
	return &Match{
		id:                        matchId,
		name:                      matchName,
		playerMap:                 make(map[string]player.Player),
		playerNameToFieldingStats: make(map[string]*player.FieldingStats),
		playerNameToLbwWickets:    make(map[string]int),
	}
}
