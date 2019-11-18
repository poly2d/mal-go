package eval

import "github.com/poly2d/mal-go/model"

func evalList(list []model.MalForm, env model.MalEnv) model.MalForm {
	lead := list[0]
	switch lead.Type {
	case model.MalTypeFunc:
		malFunc := lead.ValMalFunc()
		return malFunc(list[1:])
	case model.MalTypeSymbol:
		panic("Unexpected symbol " + lead.ValString())
	}

	// Default case for now: return list as is.
	// Perhaps its more proper to print a err msg.
	return model.MalForm{
		model.MalTypeList,
		list,
	}
}

func EvalAst(ast model.MalForm, env model.MalEnv) model.MalForm {
	switch ast.Type {
	case model.MalTypeList:
		list := ast.ValList()
		if len(list) == 0 {
			return ast
		}
		// Handle special forms
		lead := list[0]
		if lead.IsSpecialForm() {
			sym := lead.ValString()

			// Handle special forms
			switch model.SpecialForm(sym) {
			case model.SpecialFormDef:
				key := list[1].ValString()
				val := EvalAst(list[2], env)
				env.Set(key, val)
				return val

			default:
				panic("Unimplemented special form " + sym)
			}
		}
		newList := []model.MalForm{}
		for _, member := range list {
			newList = append(newList, EvalAst(member, env))
		}
		return evalList(newList, env)

	case model.MalTypeSymbol:
		if ast.IsSpecialForm() {
			return ast
		}
		return env.Get(ast.ValString())
	}

	return ast
}
