package token

const (
	TT_NUMBER = iota

	TT_PLUS
	TT_MINUS

	TT_EOF
)

var tokens map[int]string = map[int]string{
	TT_NUMBER: "NUMBER",
	TT_PLUS:   "PLUS",
	TT_MINUS:  "MINUS",
	TT_EOF:    "EOF",
}
