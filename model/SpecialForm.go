package model

type SpecialForm string

const (
	SpecialFormDef SpecialForm = "def!"
	SpecialFormFn  SpecialForm = "fn*"
	SpecialFormIf  SpecialForm = "if"
	SpecialFormLet SpecialForm = "let*"
)

var specialForms = []SpecialForm{
	SpecialFormDef,
	SpecialFormFn,
	SpecialFormIf,
	SpecialFormLet,
}

func IsSpecialForm(str string) bool {
	for _, sp := range specialForms {
		if string(sp) == str {
			return true
		}
	}
	return false
}
