package model

type MalClosure struct {
	Params MalForm // Parameter list to closure
	Body   MalForm // AST representing the closure func
	Env    *MalEnv // Outer env to base the new env on
}

func (c MalClosure) AsMalForm() MalForm {
	return MalForm{
		Type:  MalTypeClosure,
		Value: c,
	}
}
