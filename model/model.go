package model

import (
	"fmt"
	"strconv"
	"strings"
)

type MalType int

const (
	MalTypeDefault MalType = iota
	MalTypeSymbol
	MalTypeNumber
	MalTypeList
)

func (mt MalType) isAtomic() bool {
	return mt != MalTypeList
}

type MalForm struct {
	Type  MalType
	Value interface{}
}

func (mf MalForm) Sprint() string {
	var sb strings.Builder

	switch mf.Type {
	case MalTypeSymbol:
		sb.WriteString(mf.Value.(string))
	case MalTypeNumber:
		sb.WriteString(strconv.Itoa(mf.Value.(int)))
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

func (mf MalForm) Print() {
	fmt.Print(mf.Sprint())
}
