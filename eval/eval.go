package eval

import (
	"fmt"

	"github.com/poly2d/mal-go/model"
)

func getFunc(mf model.MalForm, env model.MalEnv) model.MalFunc {
	if mf.Type != model.MalTypeSymbol {
		panic(fmt.Sprintf("getFunc was called on a non-symbol MalForm mf=%v", mf))
	}

	sym := mf.Value.(string)
	f, exist := env[sym]
	if !exist {
		panic(fmt.Sprintf("Symbol %s does not exist in env", sym))
	}

	return f
}

func evalList(list []model.MalForm, env model.MalEnv) model.MalForm {
	lead := list[0]
	if lead.Type != model.MalTypeSymbol {
		return model.MalForm{
			model.MalTypeList,
			list,
		}
	}

	malFunc := getFunc(lead, env)
	return malFunc(list[1:])
}

func EvalAst(ast model.MalForm, env model.MalEnv) model.MalForm {
	switch ast.Type {
	case model.MalTypeList:
		list := ast.Value.([]model.MalForm)
		if len(list) == 0 {
			return ast
		}

		newList := []model.MalForm{}
		for _, member := range list {
			newList = append(newList, EvalAst(member, env))
		}
		return evalList(newList, env)
	}
	return ast
}
