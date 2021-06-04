package algorithms

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPageCount(t *testing.T) {
	testCases := map[string]struct {
		expectedOutput int32
		input          [2]int32
	}{
		"1": {
			expectedOutput: 0,
			input:          [2]int32{1, 1},
		},
		"2": {
			expectedOutput: 0,
			input:          [2]int32{2, 1},
		},
		"3": {
			expectedOutput: 0,
			input:          [2]int32{2, 2},
		},
		"4": {
			expectedOutput: 0,
			input:          [2]int32{3, 1},
		},
		"5": {
			expectedOutput: 0,
			input:          [2]int32{3, 2},
		},
		"6": {
			expectedOutput: 0,
			input:          [2]int32{3, 3},
		},
		"7": {
			expectedOutput: 0,
			input:          [2]int32{4, 1},
		},
		"8": {
			expectedOutput: 1,
			input:          [2]int32{4, 2},
		},
		"9": {
			expectedOutput: 1,
			input:          [2]int32{4, 3},
		},
		"10": {
			expectedOutput: 0,
			input:          [2]int32{4, 4},
		},
	}

	for name, test := range testCases {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, test.expectedOutput, pageCount(test.input[0], test.input[1]))
		})
	}
}
