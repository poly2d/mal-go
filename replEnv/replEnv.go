package replEnv

import (
	"github.com/poly2d/mal-go/model"
)

var ReplEnv = model.MalEnv{
	Outer: nil,
	Data: map[string]model.MalFunc{
		"+": Add,
		"-": Sub,
		"*": Mul,
		"/": Div,
	},
}
