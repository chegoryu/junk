package main

import (
	"fmt"
	"testing"
)

func TestGetMult(t *testing.T) {
	var tests = []struct {
		a      int
		b      int
		result int
	}{
		{0, 0, 0},
		{1, 2, 2},
		{-7, 11, -77},
		{-5, -1, 5},
		{234, 434, 101556},
	}

	for _, test := range tests {
		testName := fmt.Sprintf("%d * %d", test.a, test.b)
		t.Run(testName, func(t *testing.T) {
			result := GetMult(test.a, test.b)
			if result != test.result {
				t.Errorf("expected %d, got %d", test.result, result)
			}
		})
	}
}
