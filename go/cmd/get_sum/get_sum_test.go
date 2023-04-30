package main

import (
	"fmt"
	"testing"
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
		testName := fmt.Sprintf("%d + %d", test.A, test.B)
		t.Run(testName, func(t *testing.T) {
			result := GetSum(test.A, test.B)
			if result != test.Result {
				t.Errorf("expected %d, got %d", test.Result, result)
			}
		})
	}
}
