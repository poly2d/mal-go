package model

type MalType int

const (
	MalTypeUnit MalType = iota
	MalTypeSymbol
	MalTypeNumber
	MalTypeList

	MalTypeFunc
)

func (mt MalType) isAtomic() bool {
	return mt != MalTypeList
}
