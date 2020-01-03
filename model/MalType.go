package model

type MalType int

const (
	MalTypeNil MalType = iota
	MalTypeBool
	MalTypeSymbol
	MalTypeNumber
	MalTypeList

	MalTypeFunc
)
