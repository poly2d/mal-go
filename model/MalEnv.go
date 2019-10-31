package model

type MalEnv struct {
	Outer *MalEnv
	Data  map[string]MalFunc
}
