package core

import (
	"fmt"

	"github.com/poly2d/malgo/model"
)

func prn(args []model.MalForm) model.MalForm {
	for _, arg := range args {
		fmt.Print(arg.Sprint())
	}
	return model.MalForm{
		Type: model.MalTypeNil,
	}
}
