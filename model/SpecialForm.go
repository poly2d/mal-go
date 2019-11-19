package model

type SpecialForm string

const (
	SpecialFormDef SpecialForm = "def!"
	SpecialFormLet SpecialForm = "let*"
)

var specialForms = []SpecialForm{
	SpecialFormDef,
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
