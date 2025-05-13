package datahelper

import (
	"path/filepath"
	"testing"
)

func assertStrings(t *testing.T, result, expected string) {
	if result != expected {
		t.Fatalf("'%+v' not equal to '%+v'", result, expected)
	}
}

func setCorrectPositionsFilepath() {
	currentDir, _ := filepath.Abs(".")
	parentDir := filepath.Clean(filepath.Join(currentDir, "../.."))
	setPositionsFilepath(filepath.Join(parentDir, filepath.Clean(POSITIONS_FILENAME)))
}

func Test_Data_Helper(t *testing.T) {
	t.Run("Sanitize String", func(t *testing.T) {

		var testCases = []struct {
			text     string
			expected string
		}{
			{"", ""},
			{"abc123", "abc123"},
			{`abc
			123!@\\#$ _-+/`, "abc123"},
		}

		for _, testCase := range testCases {
			assertStrings(t, SanitizeString(testCase.text), testCase.expected)
		}
	})

	t.Run("Parse Positions", func(t *testing.T) {

		setCorrectPositionsFilepath()
		positions, err := parsePositions()

		if err != nil {
			t.Fatal(err)
		}

		expected := 3
		if len(positions) != expected {
			t.Fatalf("Positions len: '%+v' expected to be '%+v'", len(positions), expected)
		}
	})

}
