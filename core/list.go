package core

import "github.com/poly2d/malgo/model"

func list(args []model.MalForm) model.MalForm {
	return model.MalForm{
		Type:  model.MalTypeList,
		Value: args,
	}
}

func listQ(args []model.MalForm) model.MalForm {
	if err := argLenCheck(args, 1); err != nil {
		return errForm(err)
	}
	return boolForm(args[0].Type == model.MalTypeList)
}

func emptyQ(args []model.MalForm) model.MalForm {
	if err := argCheck(args, model.MalTypeList); err != nil {
		return errForm(err)
	}
	return boolForm(len(args[0].ValList()) == 0)
}

func count(args []model.MalForm) model.MalForm {
	if err := argCheck(args, model.MalTypeList); err != nil {
		return errForm(err)
	}
	return numForm(len(args[0].ValList()))
}
