package read

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadStr(t *testing.T) {
	var tests = []struct {
		in       string
		expected string
	}{
		{"abc", "abc"},
		{"123   ", "123"},
		{"( 123 456 789 )", "(123 456 789)"},
		{"     (+ 2 (* 3   4 )  )", "(+ 2 (* 3 4))"},
		{"nil", "<nil>"},
	}

	for _, test := range tests {
		mf := ReadStr(test.in)
		assert.Equal(t, test.expected, mf.Sprint())
	}
}
