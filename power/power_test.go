package exponentiation

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestN(t *testing.T) {
	for _, tt := range Cases {
		t.Run(tt.name, func(t *testing.T) {
			output := Calculate(tt.input, tt.inputN)

			require.LessOrEqual(t, tt.output*0.99, output)
			require.GreaterOrEqual(t, tt.output*1.01, output)
		})
	}
}
