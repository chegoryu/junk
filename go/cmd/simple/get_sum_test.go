package main

import (
	"fmt"
	"testing"
)

func TestGetSum(t *testing.T) {
	var tests = []struct {
		a      int
		b      int
		result int
	}{
		{0, 0, 0},
		{1, 2, 3},
		{-7, 11, 4},
		{-5, -1, -6},
		{234, 434, 668},
	}

	for _, test := range tests {
		testName := fmt.Sprintf("%d + %d", test.a, test.b)
		t.Run(testName, func(t *testing.T) {
			result := GetSum(test.a, test.b)
			if result != test.result {
				t.Errorf("expected %d, got %d", test.result, result)
			}
		})
	}
}
