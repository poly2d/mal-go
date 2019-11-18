package model

type SpecialForm string

const (
	SpecialFormDef SpecialForm = "def!"
)

var specialForms = []SpecialForm{
	SpecialFormDef,
}

func IsSpecialForm(str string) bool {
	for _, sp := range specialForms {
		if string(sp) == str {
			return true
		}
	}
	return false
}
