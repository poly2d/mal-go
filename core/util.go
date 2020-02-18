package core

import (
	"fmt"

	"github.com/poly2d/malgo/model"
)

func argLenCheck(args []model.MalForm, expectedLen int) error {
	actLen := len(args)
	expLen := expectedLen
	if actLen != expLen {
		msg := fmt.Sprintf("Incorrect num of args (expected %d, actual %d)", expLen, actLen)
		return model.MalErr(msg)
	}
	return nil
}

func argCheck(args []model.MalForm, expectedTypes ...model.MalType) error {
	if err := argLenCheck(args, len(expectedTypes)); err != nil {
		return err
	}
	for i := range args {
		actType := args[i].Type
		expType := expectedTypes[i]
		if actType != expType {
			msg := fmt.Sprintf("Incorrect type for arg at index %d (expected %d, actual %d)", i, expType, actType)
			return model.MalErr(msg)
		}
	}
	return nil
}

func boolForm(val bool) model.MalForm {
	return model.MalForm{
		Type:  model.MalTypeBool,
		Value: val,
	}
}

func numForm(val int) model.MalForm {
	return model.MalForm{
		Type:  model.MalTypeNumber,
		Value: val,
	}
}

func errForm(err error) model.MalForm {
	return model.MalForm{
		Err: err,
	}
}
