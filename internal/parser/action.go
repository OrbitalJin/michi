package parser

const (
	DEFAULT = iota
	BANG
	SHORTCUT
	SESSION
)

type QueryAction struct {
	Type     int
	Result   *Result
	RawQuery string
}
