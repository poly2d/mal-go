package read

import (
	"github.com/poly2d/mal-go/model"
)

func readAtom(r reader) model.MalForm {
	token := r.next()
	return model.MalForm{
		model.MalTypeAtom,
		token,
	}
}

func readList(r reader) model.MalForm {
	list := []model.MalForm{}
	for {
		mf := readForm(r)

		switch mf.Type {
		case model.MalTypeAtom:
			str := mf.Value.(string)
			if str == ")" {
				return model.MalForm{
					model.MalTypeList,
					list,
				}
			}
		}
		list = append(list, mf)
	}
	return model.MalForm{}
}

func readForm(r reader) model.MalForm {
	top := r.peek()
	switch top {
	case "(":
		r.next() // Discard open paren.
		return readList(r)
	default:
		return readAtom(r)

	}
	return model.MalForm{}
}

func ReadStr(in string) model.MalForm {
	reader := newReader(in)
	return readForm(reader)
}
