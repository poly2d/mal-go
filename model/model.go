package model

import (
	"fmt"
	"strings"
)

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
	fmt.Print(mf.Sprint())
}

func (mf MalForm) Sprint() string {
	var sb strings.Builder

	switch mf.Type {
	case MalTypeAtom:
		sb.WriteString(mf.Value.(string))
	case MalTypeList:
		sb.WriteString("(")

		vals := mf.Value.([]MalForm)
		for i, val := range vals {
			sb.WriteString(val.Sprint())

			if i == len(vals)-1 {
				sb.WriteString(")")
			} else {
				sb.WriteString(" ")
			}
		}
	default:
		panic(fmt.Sprintf("Invalid MalType, mf=%v\n", mf))
	}
	return sb.String()
}
