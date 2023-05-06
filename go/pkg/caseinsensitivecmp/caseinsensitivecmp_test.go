package caseinsensitivecmp

import (
	"fmt"
	"testing"
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
			{
				result := Equal(test.A, test.B)
				if result != test.Result {
					t.Errorf("expected '%t', got '%t'", test.Result, result)
				}
			}

			{
				result := Equal(test.B, test.A)
				if result != test.Result {
					t.Errorf("expected '%t', got '%t'", test.Result, result)
				}
			}
		})
	}
}

func TestLessAndGreaterOrEqual(t *testing.T) {
	var tests = []struct {
		A                  string
		B                  string
		ExpectedLessResult bool
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
			{
				result := Less(test.A, test.B)
				if result != test.ExpectedLessResult {
					t.Errorf("expected '%t', got '%t'", test.ExpectedLessResult, result)
				}
			}

			{
				result := GreaterOrEqual(test.A, test.B)
				if result != !test.ExpectedLessResult {
					t.Errorf("expected '%t', got '%t'", !test.ExpectedLessResult, result)
				}
			}

			{
				result := Less(test.B, test.A)
				expectedResult := !Equal(test.A, test.B) && !test.ExpectedLessResult
				if result != expectedResult {
					t.Errorf("expected '%t', got '%t'", expectedResult, result)
				}
			}

			{
				result := GreaterOrEqual(test.B, test.A)
				expectedResult := Equal(test.A, test.B) || test.ExpectedLessResult
				if result != expectedResult {
					t.Errorf("expected '%t', got '%t'", expectedResult, result)
				}
			}
		})
	}
}

func TestGreaterAndLessOrEqual(t *testing.T) {
	var tests = []struct {
		A                     string
		B                     string
		ExpectedGreaterResult bool
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
			{
				result := Greater(test.A, test.B)
				if result != test.ExpectedGreaterResult {
					t.Errorf("expected '%t', got '%t'", test.ExpectedGreaterResult, result)
				}
			}

			{
				result := LessOrEqual(test.A, test.B)
				if result != !test.ExpectedGreaterResult {
					t.Errorf("expected '%t', got '%t'", !test.ExpectedGreaterResult, result)
				}
			}

			{
				result := Greater(test.B, test.A)
				expectedResult := !Equal(test.A, test.B) && !test.ExpectedGreaterResult
				if result != expectedResult {
					t.Errorf("expected '%t', got '%t'", expectedResult, result)
				}
			}

			{
				result := LessOrEqual(test.B, test.A)
				expectedResult := Equal(test.A, test.B) || test.ExpectedGreaterResult
				if result != expectedResult {
					t.Errorf("expected '%t', got '%t'", expectedResult, result)
				}
			}
		})
	}
}
