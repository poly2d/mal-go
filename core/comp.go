package core

import (
	"reflect"

	"github.com/poly2d/malgo/model"
)

func eq(args []model.MalForm) model.MalForm {
	if err := argLenCheck(args, 2); err != nil {
		return errForm(err)
	}

	if args[0].Type != args[1].Type {
		return boolForm(false)
	}
	if args[0].Type == model.MalTypeList {
		list0 := args[0].ValList()
		list1 := args[1].ValList()
		if len(list0) != len(list1) {
			return boolForm(false)
		}
		for i := range list0 {
			if !eq([]model.MalForm{list0[i], list1[i]}).ValBool() {
				return boolForm(false)
			}
		}
	}
	// Todo: look into whether the use of reflect warrants
	// re-thinking the MalForm/MalType setup.
	return boolForm(reflect.DeepEqual(args[0].Value, args[1].Value))
}

func lt(args []model.MalForm) model.MalForm {
	if err := argCheck(args, model.MalTypeNumber, model.MalTypeNumber); err != nil {
		return errForm(err)
	}
	return boolForm(args[0].ValInt() < args[1].ValInt())
}

func lte(args []model.MalForm) model.MalForm {
	if err := argCheck(args, model.MalTypeNumber, model.MalTypeNumber); err != nil {
		return errForm(err)
	}
	return boolForm(args[0].ValInt() <= args[1].ValInt())
}

func gt(args []model.MalForm) model.MalForm {
	if err := argCheck(args, model.MalTypeNumber, model.MalTypeNumber); err != nil {
		return errForm(err)
	}
	return boolForm(args[0].ValInt() > args[1].ValInt())
}

func gte(args []model.MalForm) model.MalForm {
	if err := argCheck(args, model.MalTypeNumber, model.MalTypeNumber); err != nil {
		return errForm(err)
	}
	return boolForm(args[0].ValInt() >= args[1].ValInt())
}
