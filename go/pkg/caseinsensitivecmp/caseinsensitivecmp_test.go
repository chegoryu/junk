package caseinsensitivecmp

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestEqual(t *testing.T) {
	var tests = []struct {
		A      string
		B      string
		Result bool
	}{
		{"a", "A", true},
		{"A", "B", false},
		{"aBcDe", "abCDE", true},
		{"a", "aa", false},
		{"012aB345", "012Ab345", true},
		{"012ba345", "012AB345", false},
		{"first", "second", false},
	}

	for _, test := range tests {
		testName := fmt.Sprintf("%sX%s", test.A, test.B)
		t.Run(testName, func(t *testing.T) {
			require.Equal(t, test.Result, Equal(test.A, test.B))
			require.Equal(t, test.Result, Equal(test.B, test.A))
		})
	}
}

func TestLessAndGreaterOrEqual(t *testing.T) {
	var tests = []struct {
		A          string
		B          string
		LessResult bool
	}{
		{"a", "A", false},
		{"A", "B", true},
		{"abc", "abd", true},
		{"abc", "abD", true},
		{"abC", "abd", true},
		{"abC", "abD", true},
		{"abcsdfwe", "abd", true},
		{"abcsdfs", "abDlkjlkjfb", true},
		{"abC234kjoij", "abdljoioih234f", true},
		{"abCskljnsijdf112", "abDlkjlsdkjfu1923", true},
		{"aBcDe", "abCDE", false},
		{"a", "aa", true},
		{"abcdef", "abcde", false},
		{"012aB345dabc", "012Ab345cA", false},
		{"012ba34", "012AB345324", false},
		{"first", "second", true},
	}

	for _, test := range tests {
		testName := fmt.Sprintf("%sX%s", test.A, test.B)
		t.Run(testName, func(t *testing.T) {
			require.Equal(t, test.LessResult, Less(test.A, test.B))
			require.Equal(t, !test.LessResult, GreaterOrEqual(test.A, test.B))
			require.Equal(t, !Equal(test.A, test.B) && !test.LessResult, Less(test.B, test.A))
			require.Equal(t, Equal(test.A, test.B) || test.LessResult, GreaterOrEqual(test.B, test.A))
		})
	}
}

func TestGreaterAndLessOrEqual(t *testing.T) {
	var tests = []struct {
		A             string
		B             string
		GreaterResult bool
	}{
		{"a", "A", false},
		{"A", "B", false},
		{"abc", "abd", false},
		{"abc", "abD", false},
		{"abC", "abd", false},
		{"abC", "abD", false},
		{"abcsdfwe", "abd", false},
		{"abcsdfs", "abDlkjlkjfb", false},
		{"abC234kjoij", "abdljoioih234f", false},
		{"abCskljnsijdf112", "abDlkjlsdkjfu1923", false},
		{"aBcDe", "abCDE", false},
		{"a", "aa", false},
		{"abcdef", "abcde", true},
		{"012aB345dabc", "012Ab345cA", true},
		{"012ba34", "012AB345324", true},
		{"first", "second", false},
	}

	for _, test := range tests {
		testName := fmt.Sprintf("%sX%s", test.A, test.B)
		t.Run(testName, func(t *testing.T) {
			require.Equal(t, test.GreaterResult, Greater(test.A, test.B))
			require.Equal(t, !test.GreaterResult, LessOrEqual(test.A, test.B))
			require.Equal(t, !Equal(test.A, test.B) && !test.GreaterResult, Greater(test.B, test.A))
			require.Equal(t, Equal(test.A, test.B) || test.GreaterResult, LessOrEqual(test.B, test.A))
		})
	}
}
