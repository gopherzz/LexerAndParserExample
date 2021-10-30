package parser

import (
	"strconv"

	"github.com/gopherzz/simpleparserlexer/internal/parser/ast"
	"github.com/gopherzz/simpleparserlexer/internal/token"
)

var EOF token.Token = token.NewToken(token.TT_EOF, "")

type Parser struct {
	tokens      []token.Token
	pos, length int
}

func (p *Parser) Parse() []ast.Expression {
	var result []ast.Expression
	for !p.match(token.TT_EOF) {
		result = append(result, p.expression())
	}
	return result
}

func NewParser(tokens []token.Token) Parser {
	return Parser{
		tokens: tokens,
		length: len(tokens),
	}
}

func (p *Parser) expression() ast.Expression {
	return p.additive()
}

func (p *Parser) additive() ast.Expression {
	result := p.multiplicative()
	for {
		if p.match(token.TT_PLUS) {
			result = ast.NewBinaryExpression('+', result, p.multiplicative())
			continue
		} else if p.match(token.TT_MINUS) {
			result = ast.NewBinaryExpression('-', result, p.multiplicative())
			continue
		}
		break
	}
	return result
}

func (p *Parser) multiplicative() ast.Expression {
	return p.unary()
}

func (p *Parser) unary() ast.Expression {
	return p.primary()
}

func (p *Parser) primary() ast.Expression {
	curr := p.peek(0)
	if p.match(token.TT_NUMBER) {
		val, _ := strconv.ParseFloat(curr.GetValue(), 64)
		return ast.NewNumberExpresion(val)
	}
	panic("Unknown Operation")
}

func (p *Parser) peek(relativePosition int) token.Token {
	pos := p.pos + relativePosition
	if pos >= p.length {
		return EOF
	}
	return p.tokens[pos]
}

func (p *Parser) match(tokenType int) bool {
	curToken := p.peek(0)
	if tokenType != curToken.GetType() {
		return false
	}
	p.pos++
	return true
}
