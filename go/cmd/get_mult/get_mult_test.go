package main

import (
	"fmt"
	"testing"
)

func TestGetMult(t *testing.T) {
	var tests = []struct {
		A      int
		B      int
		Result int
	}{
		{0, 0, 0},
		{1, 2, 2},
		{-7, 11, -77},
		{-5, -1, 5},
		{234, 434, 101556},
	}

	for _, test := range tests {
		testName := fmt.Sprintf("%dX%d", test.A, test.A)
		t.Run(testName, func(t *testing.T) {
			result := GetMult(test.A, test.B)
			if result != test.Result {
				t.Errorf("expected %d, got %d", test.Result, result)
			}
		})
	}
}
