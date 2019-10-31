package replEnv

import (
	"github.com/poly2d/mal-go/model"
)

var ReplEnv = model.NewMalEnv(
	nil,
	map[string]model.MalForm{
		"+": GetMalFormFunc(Add),
		"-": GetMalFormFunc(Sub),
		"*": GetMalFormFunc(Mul),
		"/": GetMalFormFunc(Div),
	},
)
