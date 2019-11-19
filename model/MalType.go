package model

type MalType int

const (
	MalTypeUnit MalType = iota
	MalTypeSymbol
	MalTypeNumber
	MalTypeList

	MalTypeFunc
)
