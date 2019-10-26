package read

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Temporary test to check regexp FindStringAll
// outputs what I am expecting.
func TestTokenize(t *testing.T) {
	tokens := tokenize("     (+ 2 (* 3 4))")
	assert.Equal(t, 9, len(tokens))

	var tokenTests = []struct {
		expected string
		actual   string
	}{
		{"(", tokens[0]},
		{"+", tokens[1]},
		{"2", tokens[2]},
		{"(", tokens[3]},
		{"*", tokens[4]},
		{"3", tokens[5]},
		{"4", tokens[6]},
		{")", tokens[7]},
		{")", tokens[8]},
	}
	for _, test := range tokenTests {
		assert.Equal(t, test.actual, test.expected)
	}
}
