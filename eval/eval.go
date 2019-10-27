package eval

import (
	"fmt"

	"github.com/poly2d/mal-go/model"
)

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
		return model.MalForm{
			model.MalTypeList,
			newList,
		}

	case model.MalTypeSymbol:
		sym := ast.Value.(string)

		if _, exist := env[sym]; !exist {
			panic(fmt.Sprintf("Symbol %s does not exist in env", sym))
		}

		// Todo: apply the func from env.
		return ast
	}
	return ast
}
