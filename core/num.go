package core

import (
	"fmt"

	"github.com/poly2d/malgo/model"
)

func argCheck(args []model.MalForm, expectedTypes ...model.MalType) {
	actLen := len(args)
	expLen := len(expectedTypes)
	if actLen != expLen {
		msg := fmt.Sprintf("Incorrect num of args (expected %d, actual %d)", expLen, actLen)
		panic(msg)
	}

	for i := range args {
		actType := args[i].Type
		expType := expectedTypes[i]
		if actType != expType {
			msg := fmt.Sprintf("Incorrect type for arg at index %d (expected %d, actual %d)", i, expType, actType)
			panic(msg)
		}
	}
}

func numForm(val int) model.MalForm {
	return model.MalForm{
		model.MalTypeNumber,
		val,
	}
}

func add(args []model.MalForm) model.MalForm {
	argCheck(args, model.MalTypeNumber, model.MalTypeNumber)
	return numForm(args[0].ValInt() + args[1].ValInt())
}

func sub(args []model.MalForm) model.MalForm {
	argCheck(args, model.MalTypeNumber, model.MalTypeNumber)
	return numForm(args[0].ValInt() - args[1].ValInt())
}

func mul(args []model.MalForm) model.MalForm {
	argCheck(args, model.MalTypeNumber, model.MalTypeNumber)
	return numForm(args[0].ValInt() * args[1].ValInt())
}

func div(args []model.MalForm) model.MalForm {
	argCheck(args, model.MalTypeNumber, model.MalTypeNumber)
	return numForm(int(args[0].ValInt() / args[1].ValInt()))
}
