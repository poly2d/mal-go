package core

import "github.com/poly2d/malgo/model"

func getCoreFuncs() map[string]model.MalForm {
	coreFuncs := map[string](func(args []model.MalForm) model.MalForm){
		"+": add,
		"-": sub,
		"*": mul,
		"/": div,

		"<":  lt,
		"<=": lte,
		">":  gt,
		">=": gte,

		"list": list,
		"list?": listQ,
		"empty?": emptyQ,
		"count": count,

		"prn": prn,
	}

	mfMap := map[string]model.MalForm{}
	for sym, f := range coreFuncs {
		mfMap[sym] = model.MalFunc(f).AsMalForm()
	}
	return mfMap
}

var Env = model.NewMalEnv(nil, getCoreFuncs())
