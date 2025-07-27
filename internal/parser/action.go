package parser

const (
	BANG = iota
	SHORTCUT
	SESSION
	SEARCH
)

type QueryAction struct {
	Type     int
	Result   *Result
	RawQuery string
}
