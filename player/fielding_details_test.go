package player

import (
	"testing"
)

func TestGetFieldingScore(t *testing.T) {
	tests := []struct {
		description   string
		fieldingStats FieldingStats
		expectedScore int
	}{
		{
			description:   "Test with no catches, stumps, or run outs",
			fieldingStats: FieldingStats{},
			expectedScore: 0,
		},
		{
			description:   "Test with 2 catches",
			fieldingStats: FieldingStats{Catches: 2},
			expectedScore: 2 * CATCH_SCORE,
		},
		{
			description:   "Test with 3 catches",
			fieldingStats: FieldingStats{Catches: 3},
			expectedScore: (3 * CATCH_SCORE) + BONUS_THREE_CATCH,
		},
		{
			description:   "Test with 1 stumping",
			fieldingStats: FieldingStats{Stumps: 1},
			expectedScore: STUMP_SCORE,
		},
		{
			description:   "Test with 2 run outs",
			fieldingStats: FieldingStats{runOut: 2},
			expectedScore: 2 * RUNOUT_DIRECT_HIT,
		},
		{
			description:   "Test with 2 run outs and 1 stumping",
			fieldingStats: FieldingStats{runOut: 2, Stumps: 1},
			expectedScore: 2*RUNOUT_DIRECT_HIT + 1*STUMP_SCORE,
		},
	}

	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			score := test.fieldingStats.GetFieldingScore()
			if score != test.expectedScore {
				t.Errorf("Expected score %d but got %d", test.expectedScore, score)
			}
		})
	}
}
