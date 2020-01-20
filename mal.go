package main

import (
	"bufio"
	"fmt"
	"io"
	"os"

	"github.com/poly2d/malgo/core"
	"github.com/poly2d/malgo/eval"
	"github.com/poly2d/malgo/model"
	"github.com/poly2d/malgo/read"
)

func mRead(in string) model.MalForm {
	return read.ReadStr(in)
}

func mEval(mf model.MalForm) model.MalForm {
	return eval.EvalAst(mf, core.Env)
}

func mPrint(in model.MalForm) model.MalForm {
	in.Print()
	fmt.Println()
	return in
}

func rep(in string) {
	readRes := mRead(in)
	evalRes := mEval(readRes)
	mPrint(evalRes)
}

func main() {
	const PROMPT string = "user> "
	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print(PROMPT)
		in, err := reader.ReadString('\n')
		if err == io.EOF {
			fmt.Println("\ngoodbye")
			break
		}
		rep(in)
	}
}
