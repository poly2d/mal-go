package replEnv

import (
	"github.com/poly2d/mal-go/model"
)

var ReplEnv = model.NewMalEnv(
	nil,
	map[string]model.MalForm{
		"+": Add.AsMalForm(),
		"-": Sub.AsMalForm(),
		"*": Mul.AsMalForm(),
		"/": Div.AsMalForm(),
	},
)
