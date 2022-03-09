package path_finder

import (
	"errors"
	"testing"
)

var (
	dataSuccessOne = [][]string{
		{"a", "b"},
	}
	dataSuccessTwo = [][]string{
		{"a", "b"},
		{"b", "c"},
	}
	dataSuccess = [][]string{
		{"e", "f"},
		{"a", "b"},
		{"b", "c"},
		{"h", "i"},
		{"m", "n"},
		{"c", "d"},
		{"j", "k"},
		{"f", "g"},
		{"g", "h"},
		{"d", "e"},
		{"i", "j"},
		{"l", "m"},
		{"k", "l"},
	}
	dataBrokenChain = [][]string{
		{"e", "f"},
		{"a", "b"},
		{"b", "c"},
		{"h", "i"},
		{"m", "n"},
		{"c", "d"},
		{"j", "k"},
		{"g", "h"},
		{"d", "e"},
		{"i", "j"},
		{"l", "m"},
		{"k", "l"},
	}
	dataDuplicates = [][]string{
		{"e", "f"},
		{"a", "b"},
		{"b", "c"},
		{"h", "i"},
		{"m", "n"},
		{"c", "d"},
		{"j", "k"},
		{"f", "g"},
		{"g", "h"},
		{"d", "e"},
		{"i", "j"},
		{"l", "m"},
		{"k", "l"},

		{"a", "b"},
	}
	dataBrokenPair = [][]string{
		{"e", "f"},
		{"a", "b"},
		{"b", "c"},
		{"h"},
		{"m", "n"},
		{"c", "d"},
		{"j", "k"},
		{"f", "g"},
		{"g", "h"},
		{"d", "e"},
		{"i", "j"},
		{"l", "m"},
		{"k", "l"},
	}
)

func TestPath(t *testing.T) {
	tests := []struct {
		name        string
		input       [][]string
		expected    []string
		expectedErr error
	}{
		{"success one", dataSuccessOne, []string{"a", "b"}, nil},
		{"success two", dataSuccessTwo, []string{"a", "c"}, nil},
		{"success", dataSuccess, []string{"a", "n"}, nil},
		{"broken chain", dataBrokenChain, []string{}, errBrokenChain},
		{"duplicates", dataDuplicates, []string{}, errBrokenChain},
		{"broken pair", dataBrokenPair, []string{}, errBrokenPair},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			fl, err := Path(tc.input)
			if !errors.Is(err, tc.expectedErr) {
				t.Errorf("expected errors do not match, actual: %v, expected: %v", err, tc.expectedErr)
				return
			}

			if len(fl) != 0 {
				if fl[0] != tc.expected[0] || fl[1] != tc.expected[1] {
					t.Errorf("expected results do not match, actual: %v; expected: %v", fl, tc.expected)
					return
				}
			}
		})
	}
}
