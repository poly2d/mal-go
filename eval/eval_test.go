package eval

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/poly2d/malgo/core"
	"github.com/poly2d/malgo/read"
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
				EvalAst(ast, *core.Env)
			}
			assert.Panics(t, testFunc, "EvalAst did not panic")
			continue
		}

		actual := EvalAst(ast, *core.Env)
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

func TestSpecialFormsDefLet(t *testing.T) {
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

func TestSpecialFormFn(t *testing.T) {
	runEval(t, []testCase{
		{
			in:        "(fn* (a) a)",
			expectOut: "#<function>",
		},
		{
			in:        "( (fn* (a) a) 7)",
			expectOut: "7",
		},
		{
			in:        "( (fn* (a) (+ a 1)) 10)",
			expectOut: "11",
		},
		{
			in:        "( (fn* (a b) (+ a b)) 2 3)",
			expectOut: "5",
		},
		{
			in:          "( (fn* (a b) (+ a b)) 2 3 6)",
			expectPanic: true,
		},
		{
			in:          "( (fn* (a b) (+ a )) 2 3)",
			expectPanic: true,
		},
	})
}

func TestSpecialFormIf(t *testing.T) {
	runEval(t, []testCase{
		{
			in:        "(if true 1 0)",
			expectOut: "1",
		},
		{
			in:        "(if false 1 0)",
			expectOut: "0",
		},
		{
			in:        "(if (> 2 1) 1 0)",
			expectOut: "1",
		},
		{
			in:        "(if (> 1 2) 1 0)",
			expectOut: "0",
		},
		{
			in:        "(if (> 1 2) 1 (* 2 2))",
			expectOut: "4",
		},
		{
			in:        "(if true (* 2 2))",
			expectOut: "4",
		},
		{
			in:        "(if false (* 2 2))",
			expectOut: "<nil>",
		},
		{
			in:          "(if (> 2 1))",
			expectPanic: true,
		},
	})
}

func TestSpecialFormDo(t *testing.T) {
	runEval(t, []testCase{
		{
			in:        "(do true)",
			expectOut: "true",
		},
		{
			in:        "(do true 1)",
			expectOut: "1",
		},
		{
			in:        "(do true 1 0)",
			expectOut: "0",
		},
		{
			in:        "(do (prn 4) 2)",
			expectOut: "2",
		},
		{
			in:        "(do (prn 4))",
			expectOut: "<nil>",
		},
	})
}

func TestCoreComp(t *testing.T) {
	runEval(t, []testCase{
		{
			in:        "(= 1 2)",
			expectOut: "false",
		},
		{
			in:        "(= 1 false)",
			expectOut: "false",
		},
		{
			in:        "(= 0 nil)",
			expectOut: "false",
		},
		{
			in:        "(= 2 2)",
			expectOut: "true",
		},
		{
			in:        "(= (+ 2 2) (* 2 2))",
			expectOut: "true",
		},
		{
			in:        "(= (1 2) (1 2))",
			expectOut: "true",
		},
		{
			in:        "(= (1 2) (1))",
			expectOut: "false",
		},
		{
			in:        "(= (1 2) (1 3))",
			expectOut: "false",
		},
		{
			in:        "(= (1 2 3) (1 3))",
			expectOut: "false",
		},
		{
			in:        "(< 2 3)",
			expectOut: "true",
		},
		{
			in:        "(< 3 3)",
			expectOut: "false",
		},
		{
			in:        "(< 4 3)",
			expectOut: "false",
		},
		{
			in:        "(<= 2 3)",
			expectOut: "true",
		},
		{
			in:        "(<= 3 3)",
			expectOut: "true",
		},
		{
			in:        "(<= 4 3)",
			expectOut: "false",
		},
		{
			in:        "(> 2 3)",
			expectOut: "false",
		},
		{
			in:        "(> 3 3)",
			expectOut: "false",
		},
		{
			in:        "(> 4 3)",
			expectOut: "true",
		},
		{
			in:        "(>= 2 3)",
			expectOut: "false",
		},
		{
			in:        "(>= 3 3)",
			expectOut: "true",
		},
		{
			in:        "(>= 4 3)",
			expectOut: "true",
		},
		{
			in:          "(> 3)",
			expectPanic: true,
		},
	})
}

func TestCoreList(t *testing.T) {
	runEval(t, []testCase{
		{
			in:        "(list)",
			expectOut: "()",
		},
		{
			in:        "(list 2 3 4)",
			expectOut: "(2 3 4)",
		},
		{
			in:        "(list (1 2) 3 4)",
			expectOut: "((1 2) 3 4)",
		},
		{
			in:        "(list? 4)",
			expectOut: "false",
		},
		{
			in:        "(list? nil)",
			expectOut: "false",
		},
		{
			in:        "(list? ())",
			expectOut: "true",
		},
		{
			in:        "(list? (1 2 3))",
			expectOut: "true",
		},
		{
			in:          "(list? (1 2 3) (45))",
			expectPanic: true,
		},
		{
			in:        "(empty? ())",
			expectOut: "true",
		},
		{
			in:        "(empty? (1 2))",
			expectOut: "false",
		},
		{
			in:          "(empty? (1 2) (3))",
			expectPanic: true,
		},
		{
			in:        "(count ())",
			expectOut: "0",
		},
		{
			in:        "(count (1 2))",
			expectOut: "2",
		},
		{
			in:          "(count (1 2) (3))",
			expectPanic: true,
		},
	})
}
