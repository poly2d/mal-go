package env

import "github.com/poly2d/mal-go/model"

var ReplEnv = model.MalEnv{
	"+": func(a, b int) int { return a + b },
	"-": func(a, b int) int { return a - b },
	"*": func(a, b int) int { return a * b },
	"/": func(a, b int) int { return int(a / b) },
}
