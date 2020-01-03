package eval

import (
	"fmt"

	"github.com/poly2d/malgo/model"
)

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

		// Handle special forms.
		lead := list[0]
		if lead.IsSpecialForm() {
			sym := lead.ValString()
			switch model.SpecialForm(sym) {
			case model.SpecialFormDef:
				key := list[1].ValString()
				val := EvalAst(list[2], env)
				env.Set(key, val)
				return val

			case model.SpecialFormFn:
				c := model.MalClosure{
					Params: list[1],
					Body:   list[2], // Not evaluated
				}
				return c.AsMalForm()

			case model.SpecialFormLet:
				bindingList := list[1].ValList()
				initMap := map[string]model.MalForm{}
				for i := 0; i < len(bindingList); i += 2 {
					key := bindingList[i].ValString()
					if i+1 >= len(bindingList) {
						panic("Unexpected symbol " + key + " - no value specified")
					}
					initMap[key] = bindingList[i+1]
				}
				letEnv := model.NewMalEnv(&env, initMap)
				return EvalAst(list[2], *letEnv)

			default:
				panic("Unimplemented special form " + sym)
			}
		}

		newList := []model.MalForm{}
		for _, member := range list {
			newList = append(newList, EvalAst(member, env))
		}

		// Handle Closures
		if newList[0].Type == model.MalTypeClosure {
			mc := newList[0].ValMalClosure()
			binds := mc.Params.ValList()
			exprs := list[1:]
			if len(binds) != len(exprs) {
				panic(fmt.Sprintf("Expected %d params in expression for closure", len(binds)))
			}
			initMap := map[string]model.MalForm{}
			for i := range binds {
				key := binds[i].ValString()
				val := EvalAst(exprs[i], env)
				initMap[key] = val
			}

			closureEnv := model.NewMalEnv(&env, initMap)
			return EvalAst(mc.Body, *closureEnv)
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
