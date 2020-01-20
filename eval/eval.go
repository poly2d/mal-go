package eval

import (
	"fmt"

	"github.com/poly2d/malgo/model"
)

func EvalAst(ast model.MalForm, env *model.MalEnv) model.MalForm {
	for { // Loop is used here for Tail-call optimization.
		switch ast.Type {

		case model.MalTypeSymbol:
			if ast.IsSpecialForm() {
				return ast
			}
			return env.Get(ast.ValString())

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

				case model.SpecialFormDo:
					lastIndex := len(list) - 1
					for _, doAst := range list[1:lastIndex] {
						EvalAst(doAst, env)
					}
					ast = list[lastIndex]
					continue

				case model.SpecialFormFn:
					c := model.MalClosure{
						Params: list[1],
						Body:   list[2], // Not evaluated
						Env:    env,
					}
					return c.AsMalForm()

				case model.SpecialFormIf:
					cond := EvalAst(list[1], env).ValBool()
					if cond {
						ast = list[2]
						continue
					}
					if len(list) != 4 { // Return nil if expr for false path is not provided.
						return model.MalForm{
							Type: model.MalTypeNil,
						}
					}
					ast = list[3]
					continue

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
					env = model.NewMalEnv(env, initMap)
					ast = list[2]
					continue

				default:
					panic("Unimplemented special form " + sym)
				}
			}

			newList := []model.MalForm{}
			for _, member := range list {
				newList = append(newList, EvalAst(member, env))
			}

			{
				// This block was initially a separate func (evalList), but was moved here for TCO.
				list := newList
				lead := list[0]
				switch lead.Type {

				case model.MalTypeFunc:
					malFunc := lead.ValMalFunc()
					return malFunc(list[1:])

				case model.MalTypeClosure:
					mc := lead.ValMalClosure()
					binds := mc.Params.ValList()
					exprs := list[1:]
					if len(binds) != len(exprs) {
						panic(fmt.Sprintf("Expected %d params in expression for closure", len(binds)))
					}
					initMap := map[string]model.MalForm{}
					for i := range binds {
						key := binds[i].ValString()
						val := EvalAst(exprs[i], mc.Env)
						initMap[key] = val
					}

					ast = mc.Body
					env = model.NewMalEnv(mc.Env, initMap)
					continue

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
		}

		return ast
	}
}
