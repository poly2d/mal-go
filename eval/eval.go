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
					// TODO check error field before setting
					val := EvalAst(list[2], env)
					if val.Err == nil {
						env.Set(key, val)
					}
					return val

				case model.SpecialFormDo:
					lastIndex := len(list) - 1
					for _, doAst := range list[1:lastIndex] {
						// TODO check error field before moving on
						ea := EvalAst(doAst, env)
						if ea.Err != nil {
							return ea
						}
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
					// TODO check error field before moving on
					val := EvalAst(list[1], env)
					if val.Err != nil {
						return val
					}
					cond := val.ValBool()
					if cond {
						if len(list) < 3 { // Return error if expr for true path is not provided.
							return model.MalForm{
								Err: model.MalErr("No expression provided for true path"),
							}
						}
						ast = list[2]
						continue
					}
					if len(list) < 4 { // Return nil if expr for false path is not provided.
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
							return model.MalForm{
								Err: model.MalErr("Unexpected symbol " + key + " - no value specified"),
							}

						}
						initMap[key] = bindingList[i+1]
					}
					env = model.NewMalEnv(env, initMap)
					ast = list[2]
					continue

				default:
					return model.MalForm{
						Err: model.MalErr("Unimplemented special form " + sym),
					}
				}
			}

			newList := []model.MalForm{}
			for _, member := range list {
				// TODO check error field before moving on
				val := EvalAst(member, env)
				if val.Err != nil {
					return val
				}
				newList = append(newList, val)
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
						errMsg := fmt.Sprintf("Expected %d params in expression for closure, got %d", len(binds), len(exprs))
						return model.MalForm{
							Err: model.MalErr(errMsg),
						}

					}
					initMap := map[string]model.MalForm{}
					for i := range binds {
						key := binds[i].ValString()
						// TODO check error field before moving on
						val := EvalAst(exprs[i], mc.Env)
						if val.Err != nil {
							return val
						}
						initMap[key] = val
					}

					ast = mc.Body
					env = model.NewMalEnv(mc.Env, initMap)
					continue

				case model.MalTypeSymbol:
					return model.MalForm{
						Err: model.MalErr("Unexpected symbol " + lead.ValString()),
					}
				}

				// Default case for now: return list as is.
				// Perhaps its more proper to print a err msg.
				return model.MalForm{
					Type:  model.MalTypeList,
					Value: list,
				}
			}
		}

		return ast
	}
}
