package eval

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/poly2d/mal-go/read"
	env "github.com/poly2d/mal-go/replEnv"
)

func TestReadStr(t *testing.T) {
	var tests = []struct {
		in       string
		expected string
	}{
		{"abc", "abc"},
		{"123   ", "123"},
		{"(+ 2  3)", "5"},
		{"     (+ 2 (* 3   4 )  )", "14"},

		// Todo: figure out handling of eval of list without symbol.
		{"( 123 456 789 )", "(123 456 789)"},
	}

	// Todo: check for panics on ill-formed inputs
	for _, test := range tests {
		ast := read.ReadStr(test.in)
		actual := EvalAst(ast, *env.ReplEnv)
		assert.Equal(t, test.expected, actual.Sprint())
	}
}
