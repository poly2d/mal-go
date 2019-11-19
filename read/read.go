package read

import (
	"strconv"

	"github.com/poly2d/malgo/model"
)

func readAtom(r reader) model.MalForm {
	token := r.next()

	if num, err := strconv.Atoi(token); err == nil {
		return model.MalForm{
			model.MalTypeNumber,
			num,
		}
	}

	return model.MalForm{
		model.MalTypeSymbol,
		token,
	}
}

func readList(r reader) model.MalForm {
	list := []model.MalForm{}
	for {
		mf := readForm(r)

		switch mf.Type {
		case model.MalTypeSymbol:
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
	}
	return readAtom(r)
}

func ReadStr(in string) model.MalForm {
	reader := newReader(in)
	return readForm(reader)
}
