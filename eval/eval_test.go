package eval

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/poly2d/mal-go/read"
	env "github.com/poly2d/mal-go/replEnv"
)

type testCase struct {
	in          string
	expectOut   string
	expectPanic bool
}

func runEval(t *testing.T, tests []testCase) {
	for _, test := range tests {
		ast := read.ReadStr(test.in)
		if test.expectPanic {
			testFunc := func() {
				EvalAst(ast, *env.ReplEnv)
			}
			assert.Panics(t, testFunc, "EvalAst did not panic")
			continue
		}

		actual := EvalAst(ast, *env.ReplEnv)
		assert.Equal(t, test.expectOut, actual.Sprint())
	}
}

func TestReadStr(t *testing.T) {
	runEval(t, []testCase{
		{
			in:          "abc",
			expectPanic: true,
		},
		{
			in:        "123   ",
			expectOut: "123",
		},
		{
			in:        "(+ 2  3)",
			expectOut: "5",
		},
		{
			in:        "     (+ 2 (* 3   4 )  )",
			expectOut: "14",
		},
		{
			// Todo: figure out handling of eval of list without symbol.
			in:        "( 123 456 789 )",
			expectOut: "(123 456 789)",
		},
	})
}

func TestSpecialForms(t *testing.T) {
	runEval(t, []testCase{
		{
			in:          "a",
			expectPanic: true,
		},
		{
			in:        "(def! a 6)",
			expectOut: "6",
		},
		{
			in:        "a",
			expectOut: "6",
		},
		{
			in:        "(def! b (+ a 2))",
			expectOut: "8",
		},
		{
			in:        "(+ a b)",
			expectOut: "14",
		},
		{
			in:        "(let* (c 5) c)",
			expectOut: "5",
		},
		{
			in:          "c",
			expectPanic: true,
		},
		{
			in:        "(let* (c 2 d 4) (+ b (* c a)))",
			expectOut: "20",
		},
	})
}

func TestSpecialFormLet(t *testing.T) {
	runEval(t, []testCase{})
}
