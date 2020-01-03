package core

import (
	"fmt"

	"github.com/poly2d/malgo/model"
)

func argLenCheck(args []model.MalForm, expectedLen int) {
	actLen := len(args)
	expLen := expectedLen
	if actLen != expLen {
		msg := fmt.Sprintf("Incorrect num of args (expected %d, actual %d)", expLen, actLen)
		panic(msg)
	}
}

func argCheck(args []model.MalForm, expectedTypes ...model.MalType) {
	argLenCheck(args, len(expectedTypes))
	for i := range args {
		actType := args[i].Type
		expType := expectedTypes[i]
		if actType != expType {
			msg := fmt.Sprintf("Incorrect type for arg at index %d (expected %d, actual %d)", i, expType, actType)
			panic(msg)
		}
	}
}

func boolForm(val bool) model.MalForm {
	return model.MalForm{
		model.MalTypeBool,
		val,
	}
}

func numForm(val int) model.MalForm {
	return model.MalForm{
		model.MalTypeNumber,
		val,
	}
}
