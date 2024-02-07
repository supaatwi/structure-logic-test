package util_test

import (
	"logictest/util"
	"reflect"
	"testing"
)

func TestCountSmileyFace(t *testing.T) {
	testCases := map[string]struct {
		input    []string
		expected int
	}{
		"Test Case 1": {
			input:    []string{":)", ";(", ";}", ":-D"},
			expected: 2,
		},
		"Test Case 2": {
			input:    []string{";D", ":-(", ":-)", ";~)"},
			expected: 3,
		},
		"Test Case 3": {
			input:    []string{";]", ":[", ";*", ":$", ";-D"},
			expected: 1,
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			result := util.CountSmileyFace(tc.input)
			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("unexpected result for input %q: got %v, want %v", tc.input, result, tc.expected)
			}
		})
	}
}
