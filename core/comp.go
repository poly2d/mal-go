package core

import "github.com/poly2d/malgo/model"

func lt(args []model.MalForm) model.MalForm {
	argCheck(args, model.MalTypeNumber, model.MalTypeNumber)
	return boolForm(args[0].ValInt() < args[1].ValInt())
}

func lte(args []model.MalForm) model.MalForm {
	argCheck(args, model.MalTypeNumber, model.MalTypeNumber)
	return boolForm(args[0].ValInt() <= args[1].ValInt())
}

func gt(args []model.MalForm) model.MalForm {
	argCheck(args, model.MalTypeNumber, model.MalTypeNumber)
	return boolForm(args[0].ValInt() > args[1].ValInt())
}

func gte(args []model.MalForm) model.MalForm {
	argCheck(args, model.MalTypeNumber, model.MalTypeNumber)
	return boolForm(args[0].ValInt() >= args[1].ValInt())
}
