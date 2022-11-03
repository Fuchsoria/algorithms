package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestN(t *testing.T) {
	cases = cases[:4]
	lucky := Lucky{}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			result := lucky.GetLuckyCount(tt.input)
			require.Equal(t, tt.output, result)
			fmt.Printf("Result: %d\n", result)
		})
	}
}
