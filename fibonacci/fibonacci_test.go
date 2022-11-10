package fibonacci

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCalculate(t *testing.T) {
	for _, tt := range Cases {
		t.Run(tt.name, func(t *testing.T) {
			output := Calculate(tt.input)

			require.LessOrEqual(t, tt.output*0.99, output)
			require.GreaterOrEqual(t, tt.output*1.01, output)
		})
	}
}

func TestCalculateRecursive(t *testing.T) {
	Cases := Cases[:7]

	for _, tt := range Cases {
		t.Run(tt.name, func(t *testing.T) {
			output := CalculateRecursive(tt.input)

			require.LessOrEqual(t, tt.output*0.99, output)
			require.GreaterOrEqual(t, tt.output*1.01, output)
		})
	}
}

func TestCalculateGold(t *testing.T) {
	//	Cases := Cases[:7]

	for _, tt := range Cases {
		t.Run(tt.name, func(t *testing.T) {
			output := CalculateGold(tt.input)

			require.LessOrEqual(t, tt.output*0.99, output)
			require.GreaterOrEqual(t, tt.output*1.01, output)
		})
	}
}
