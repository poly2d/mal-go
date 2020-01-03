package core

import "github.com/poly2d/malgo/model"

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
