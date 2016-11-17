package chop_test

import (
	"testing"

	"github.com/mcandre/go-chop"
)

func TestChopTable(t *testing.T) {
	testCases := []struct{ Expected, Input string }{
		{"", ""},
		{"", "A"},
		{"A", "AB"},
		{"A", "A\n"},
		{"A\r", "A\r\n"},
		{"\r", "\r\n"},
	}

	for _, testCase := range testCases {
		observed := chop.Chop(testCase.Input)

		if observed != testCase.Expected {
			t.Errorf("expected chop.Chop(%v) => %v, got %v", testCase.Input, testCase.Expected, observed)
		}
	}
}

func TestChompTable(t *testing.T) {
	testCases := []struct{ Expected, Input string }{
		{"", ""},
		{"A", "A"},
		{"AB", "AB"},
		{"A", "A\n"},
		{"A", "A\r\n"},
		{"", "\r\n"},
	}

	for _, testCase := range testCases {
		observed := chop.Chomp(testCase.Input)

		if observed != testCase.Expected {
			t.Errorf("expected chop.Chomp(%v) => %v, got %v", testCase.Input, testCase.Expected, observed)
		}
	}
}
