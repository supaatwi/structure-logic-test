package util_test

import (
	"logictest/util"
	"reflect"
	"testing"
)

func TestPermutation(t *testing.T) {
	testCases := map[string]struct {
		input    string
		expected []string
	}{
		"Test Case 1": {
			input:    "a",
			expected: []string{"a"},
		},
		"Test Case 2": {
			input:    "ab",
			expected: []string{"ab", "ba"},
		},
		"Test Case 3": {
			input:    "abc",
			expected: []string{"abc", "acb", "bac", "bca", "cab", "cba"},
		},
		"Test Case 4": {
			input:    "aabb",
			expected: []string{"aabb", "abab", "abba", "baab", "baba", "bbaa"},
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			result := util.Permutation(tc.input)
			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("unexpected result for input %q: got %v, want %v", tc.input, result, tc.expected)
			}
		})
	}
}
