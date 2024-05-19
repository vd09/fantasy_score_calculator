package player

import (
	"testing"
)

func TestGetBowlingScore(t *testing.T) {
	tests := []struct {
		description   string
		bowlingStats  BowlingStats
		expectedScore int
	}{
		{
			description: "Test with no wickets and maiden overs",
			bowlingStats: BowlingStats{
				Runs:        25,
				DotBalls:    20,
				Wickets:     0,
				LbwWickets:  0,
				MaidenOvers: 2,
				Overs:       4.2,
			},
			expectedScore: 0 + 0 + (2 * MAIDEN_OVER_SCORE) + 4,
		},
		{
			description: "Test with 3 wickets and LBW wickets",
			bowlingStats: BowlingStats{
				Runs:        35,
				DotBalls:    15,
				Wickets:     3,
				LbwWickets:  2,
				MaidenOvers: 1,
				Overs:       5,
			},
			expectedScore: (3 * WICKET_SCORE) + (3 * THREE_WICKET_BONUS) + (2 * LBW_BONUS) + 0 + 6,
		},
		{
			description: "Test with high economy rate",
			bowlingStats: BowlingStats{
				Runs:        60,
				DotBalls:    10,
				Wickets:     1,
				LbwWickets:  0,
				MaidenOvers: 0,
				Overs:       3.4,
			},
			expectedScore: (1 * WICKET_SCORE) + 0 + 0 + (-6),
		},
	}

	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			score := test.bowlingStats.GetBowlingScore()
			if score != test.expectedScore {
				t.Errorf("Expected score %d but got %d", test.expectedScore, score)
			}
		})
	}
}
