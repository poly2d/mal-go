package read

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// For verifying behavior of regexp FindStringAll.
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
		assert.Equal(t, test.expected, test.actual)
	}
}

func TestReaderImpl(t *testing.T) {
	var readerTests = []struct {
		testStr        string
		expectedTokens []string
	}{
		{"abc", []string{"abc"}},
		{"123", []string{"123"}},
		{"( 123 456 789 )", []string{"(", "123", "456", "789", ")"}},
		{"     (+ 2 (* 3 4))", []string{"(", "+", "2", "(", "*", "3", "4", ")", ")"}},
	}

	for _, test := range readerTests {
		testStr := test.testStr
		expectedTokens := test.expectedTokens

		reader := newReader(testStr)
		for _, expected := range expectedTokens {
			assert.Equal(t, expected, reader.peek())
			assert.Equal(t, expected, reader.next())
		}
	}

}
