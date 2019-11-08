package model

type MalFunc func([]MalForm) MalForm

func (f MalFunc) AsMalForm() MalForm {
	return MalForm{
		Type:  MalTypeFunc,
		Value: f,
	}
}
