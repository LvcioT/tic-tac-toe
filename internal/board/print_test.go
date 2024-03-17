package board

import (
	"testing"
)

func TestHorizontalSeparator(t *testing.T) {
	testCases := []struct {
		label    string
		n        int
		expected string
	}{
		{"one", 1, ""},
		{"three", 3, "─┼─┼─"},
		{"eight", 8, "─┼─┼─┼─┼─┼─┼─┼─"},
	}

	for _, tc := range testCases {
		actual := horizontalSeparator(tc.n)

		if actual != tc.expected {
			t.Logf("%s: expected '%s', got '%s'", tc.label, tc.expected, actual)
			t.Fail()
		}
	}
}

func TestHorizontalRow(t *testing.T) {
	testCases := []struct {
		label    string
		runes    []Stone
		expected string
	}{
		{"one", []Stone{'A'}, "A"},
		{"three", []Stone{'A', 'B', 'C'}, "A│B│C"},
		{"eight", []Stone{'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H'}, "A│B│C│D│E│F│G│H"},
	}

	for _, tc := range testCases {
		actual := horizontalRow(tc.runes)

		if actual != tc.expected {
			t.Logf("%s: expected '%s', got '%s'", tc.label, tc.expected, actual)
			t.Fail()
		}
	}
}
