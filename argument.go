package cmd_repl

import "strconv"

type argument struct {
	value string
}

func NewArgument(value string) *argument {
	return &argument{
		value: value,
	}
}

func (a *argument) AsInt() (int, error) {
	return strconv.Atoi(a.value)
}

func (a *argument) AsFloat() (float64, error) {
	return strconv.ParseFloat(a.value, 32)
}

func (a *argument) AsString() (string, error) {
	return a.value, nil
}
