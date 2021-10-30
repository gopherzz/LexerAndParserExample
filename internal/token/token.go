package token

import "fmt"

type Token struct {
	tokenType int
	value     string
}

func NewToken(tokenType int, value string) Token {
	return Token{
		tokenType: tokenType,
		value:     value,
	}
}

func (t *Token) GetType() int {
	return t.tokenType
}

func (t *Token) GetValue() string {
	return t.value
}

func (t Token) String() string {
	return fmt.Sprintf("%s %s", tokens[t.tokenType], t.value)
}
