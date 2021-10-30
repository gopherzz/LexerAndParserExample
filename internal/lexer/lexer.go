package lexer

import (
	"strings"
	"unicode"

	"github.com/gopherzz/simpleparserlexer/internal/token"
)

const (
	OPERATORS_RUNES = "+-"
)

var OPERATOR_TOKENS = []int{
	token.TT_PLUS, token.TT_MINUS,
}

type Lexer struct {
	input       string
	tokens      []token.Token
	pos, length int
}

func NewLexer(input string) Lexer {
	return Lexer{
		input:  input,
		tokens: []token.Token{},
		length: len(input),
	}
}

func (l *Lexer) Tokenize() []token.Token {
	for l.pos < l.length {
		curr := l.peek(0)

		if unicode.IsDigit(curr) {
			l.tokenizeNumber()
		} else if strings.ContainsRune(OPERATORS_RUNES, curr) {
			l.tokenizeOperator()
		} else {
			l.next()
		}
	}
	return l.tokens
}

func (l *Lexer) tokenizeNumber() {
	var buf string
	cur := l.peek(0)
	for {
		if !unicode.IsDigit(cur) {
			break
		}
		buf += string(cur)
		cur = l.next()
	}
	l.addToken(token.TT_NUMBER, buf)
}

func (l *Lexer) tokenizeOperator() {
	pos := strings.IndexRune(OPERATORS_RUNES, l.peek(0))
	l.addToken(OPERATOR_TOKENS[pos], "")
	l.next()
}

func (l *Lexer) addToken(tokenType int, value string) {
	l.tokens = append(l.tokens, token.NewToken(tokenType, value))
}

func (l *Lexer) peek(relativePosition int) rune {
	pos := l.pos + relativePosition
	if pos >= l.length {
		return '\n'
	}
	return rune(l.input[pos])
}

func (l *Lexer) next() rune {
	l.pos++
	return l.peek(0)
}
