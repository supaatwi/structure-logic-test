package util_test

import (
	"logictest/util"
	"reflect"
	"testing"
)

func TestFindTheOddInt(t *testing.T) {
	testCases := map[string]struct {
		input    []int
		expected int
	}{
		"Test Case 1": {
			input:    []int{7},
			expected: 7,
		},
		"Test Case 2": {
			input:    []int{0},
			expected: 0,
		},
		"Test Case 3": {
			input:    []int{1, 1, 2},
			expected: 2,
		},
		"Test Case 4": {
			input:    []int{0, 1, 0, 1, 0},
			expected: 0,
		},
		"Test Case 5": {
			input:    []int{1, 2, 2, 3, 3, 3, 4, 3, 3, 3, 2, 2, 1},
			expected: 4,
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			result := util.FindTheOddInt(tc.input)
			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("unexpected result for input %q: got %v, want %v", tc.input, result, tc.expected)
			}
		})
	}

}
