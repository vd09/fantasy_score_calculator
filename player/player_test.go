package player

import (
	"testing"
)

func TestGetPlayerScore(t *testing.T) {
	tests := []struct {
		description   string
		player        Player
		expectedScore int
	}{
		{
			description: "Test player with only batting stats",
			player: Player{
				BattingStats: BattingStats{
					Runs:        50,
					TotalFours:  6,
					TotalSixes:  2,
					PlayedBalls: 30,
				},
			},
			expectedScore: 50 + (6 * FOUR_BOUNDARY_BONUS) + (2 * SIX_BOUNDARY_BONUS) + 8 + 4,
		},
		{
			description: "Test player with only bowling stats",
			player: Player{
				BowlingStats: BowlingStats{
					Runs:        25,
					DotBalls:    20,
					Wickets:     3,
					LbwWickets:  2,
					MaidenOvers: 2,
					Overs:       4.2,
				},
			},
			expectedScore: (3 * WICKET_SCORE) + THREE_WICKET_BONUS + (2 * LBW_BONUS) + (2 * MAIDEN_OVER_SCORE) + 4,
		},
		{
			description: "Test player with only fielding stats",
			player: Player{
				FieldingStats: FieldingStats{
					Catches: 3,
					Stumps:  1,
					runOut:  2,
				},
			},
			expectedScore: (3 * CATCH_SCORE) + BONUS_THREE_CATCH + STUMP_SCORE + (2 * RUNOUT_DIRECT_HIT),
		},
		{
			description: "Test player with all stats",
			player: Player{
				BattingStats: BattingStats{
					Runs:        60,
					TotalFours:  8,
					TotalSixes:  4,
					PlayedBalls: 40,
				},
				BowlingStats: BowlingStats{
					Runs:        35,
					DotBalls:    15,
					Wickets:     3,
					LbwWickets:  2,
					MaidenOvers: 1,
					Overs:       5,
				},
				FieldingStats: FieldingStats{
					Catches: 2,
					Stumps:  1,
					runOut:  1,
				},
			},
			expectedScore: (60 + (8 * FOUR_BOUNDARY_BONUS) + (4 * SIX_BOUNDARY_BONUS) + 8 + 2) + ((3 * WICKET_SCORE) + (THREE_WICKET_BONUS) + (2 * LBW_BONUS) + (1 * MAIDEN_OVER_SCORE) + 2) + ((2 * CATCH_SCORE) + STUMP_SCORE + RUNOUT_DIRECT_HIT),
		},
	}

	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			score := test.player.GetPlayerScore()
			if score != test.expectedScore {
				t.Errorf("Expected score %d but got %d", test.expectedScore, score)
			}
		})
	}
}
