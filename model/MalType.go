package model

type MalType int

const (
	MalTypeInvalid MalType = iota
	MalTypeNil
	MalTypeBool
	MalTypeSymbol
	MalTypeString
	MalTypeNumber
	MalTypeList
	MalTypeClosure

	MalTypeFunc
)
