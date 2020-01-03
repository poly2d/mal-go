package core

import "github.com/poly2d/malgo/model"

func list(args []model.MalForm) model.MalForm {
	return model.MalForm{
		Type:  model.MalTypeList,
		Value: args,
	}
}

func listQ(args []model.MalForm) model.MalForm {
	argLenCheck(args, 1)
	return boolForm(args[0].Type == model.MalTypeList)
}

func emptyQ(args []model.MalForm) model.MalForm {
	argCheck(args, model.MalTypeList)
	return boolForm(len(args[0].ValList()) == 0)
}

func count(args []model.MalForm) model.MalForm {
	argCheck(args, model.MalTypeList)
	return numForm(len(args[0].ValList()))
}
