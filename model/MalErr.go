package model

type MalErr string

func (m MalErr) Error() string {
	return string(m)
}
