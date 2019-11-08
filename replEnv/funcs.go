package replEnv

import (
	"fmt"

	"github.com/poly2d/mal-go/model"
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
	return numForm(args[0].Value.(int) + args[1].Value.(int))
}

func sub(args []model.MalForm) model.MalForm {
	argCheck(args, model.MalTypeNumber, model.MalTypeNumber)
	return numForm(args[0].Value.(int) - args[1].Value.(int))
}

func mul(args []model.MalForm) model.MalForm {
	argCheck(args, model.MalTypeNumber, model.MalTypeNumber)
	return numForm(args[0].Value.(int) * args[1].Value.(int))
}

func div(args []model.MalForm) model.MalForm {
	argCheck(args, model.MalTypeNumber, model.MalTypeNumber)
	return numForm(int(args[0].Value.(int) / args[1].Value.(int)))
}

var (
	Add = model.MalFunc(add)
	Sub = model.MalFunc(sub)
	Mul = model.MalFunc(mul)
	Div = model.MalFunc(div)
)
