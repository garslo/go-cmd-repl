package cmd_repl

type Argument interface {
	AsInt() (int, error)
	AsFloat() (float64, error)
	AsString() (string, error)
}
