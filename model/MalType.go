package model

type MalType int

const (
	MalTypeInvalid MalType = iota
	MalTypeNil
	MalTypeBool
	MalTypeSymbol
	MalTypeNumber
	MalTypeList
	MalTypeClosure

	MalTypeFunc
)
