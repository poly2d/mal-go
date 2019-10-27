package model

import "fmt"

type MalType int

const (
	MalTypeDefault MalType = iota
	MalTypeAtom            // For the purposes of s1 we'll assume all atoms are strings.
	MalTypeList
)

type MalForm struct {
	Type  MalType
	Value interface{}
}

func (mf MalForm) Print() {
	switch mf.Type {
	case MalTypeAtom:
		str := mf.Value.(string)
		fmt.Print(str)
	case MalTypeList:
		vals := mf.Value.([]MalForm)
		fmt.Print("(")
		for i, val := range vals {
			val.Print()

			if i == len(vals)-1 {
				fmt.Print(")")
			} else {
				fmt.Print(" ")
			}
		}
	default:
		panic(fmt.Sprintf("Invalid MalType, mf=%v\n", mf))
	}
}
