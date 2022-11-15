package bits

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestKing(t *testing.T) {
	chess := Chess{}

	for _, tt := range casesKing {
		t.Run(tt.Name, func(t *testing.T) {
			mask := chess.GetKingBitboardMoves(tt.Input)
			count := chess.Popcnt(mask)
			countOpt := chess.PopcntOpt(mask)

			require.Equal(t, tt.Output, mask)
			require.Equal(t, tt.Count, count)
			require.Equal(t, tt.Count, countOpt)
		})
	}
}

func TestHorse(t *testing.T) {
	chess := Chess{}

	for _, tt := range casesHorse {
		t.Run(tt.Name, func(t *testing.T) {
			mask := chess.GetHorseBitboardMoves(tt.Input)
			count := chess.Popcnt(mask)
			countOpt := chess.PopcntOpt(mask)

			require.Equal(t, tt.Output, mask)
			require.Equal(t, tt.Count, count)
			require.Equal(t, tt.Count, countOpt)
		})
	}
}
