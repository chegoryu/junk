package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetSum(t *testing.T) {
	var tests = []struct {
		A      int
		B      int
		Result int
	}{
		{0, 0, 0},
		{1, 2, 3},
		{-7, 11, 4},
		{-5, -1, -6},
		{234, 434, 668},
	}

	for _, test := range tests {
		testName := fmt.Sprintf("%dX%d", test.A, test.B)
		t.Run(testName, func(t *testing.T) {
			require.Equal(t, test.Result, GetSum(test.A, test.B))
		})
	}
}
