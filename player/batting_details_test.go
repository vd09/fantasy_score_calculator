package player

import (
	"testing"
)

func TestGetBattingScore(t *testing.T) {
	tests := []struct {
		description   string
		battingStats  BattingStats
		expectedScore int
	}{
		{
			description: "Test with Runs less than 30 and no boundaries",
			battingStats: BattingStats{
				Runs:        20,
				TotalFours:  0,
				TotalSixes:  0,
				PlayedBalls: 30,
			},
			expectedScore: 20 - 2, /* strike rate*/
		},
		{
			description: "Test with Runs greater than 50 and some boundaries",
			battingStats: BattingStats{
				Runs:        60,
				TotalFours:  4,
				TotalSixes:  2,
				PlayedBalls: 40,
			},
			expectedScore: 60 + HALF_CENTURY_BONUS + (4 * FOUR_BOUNDARY_BONUS) + (2 * SIX_BOUNDARY_BONUS) + 2, /* strike rate*/
		},
		{
			description: "Test with Duck on 3rd ball",
			battingStats: BattingStats{
				Runs:        0,
				TotalFours:  0,
				TotalSixes:  0,
				PlayedBalls: 3,
				OutDetails:  "catch",
			},
			expectedScore: 0 + (-2),
		},
		{
			description: "Test with no Runs",
			battingStats: BattingStats{
				Runs:        0,
				TotalFours:  0,
				TotalSixes:  0,
				PlayedBalls: 10,
				OutDetails:  "catch",
			},
			expectedScore: 0 + (-2) + (-6),
		},
	}

	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			score := test.battingStats.GetBattingScore()
			if score != test.expectedScore {
				t.Errorf("Expected score %d but got %d", test.expectedScore, score)
			}
		})
	}
}
