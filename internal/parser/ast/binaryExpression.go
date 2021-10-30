package ast

import "fmt"

type BinaryExpression struct {
	op           rune
	expr1, expr2 Expression
}

func NewBinaryExpression(operator rune, expr1, expr2 Expression) BinaryExpression {
	return BinaryExpression{
		op:    operator,
		expr1: expr1,
		expr2: expr2,
	}
}

func (e BinaryExpression) Eval() float64 {
	num1, num2 := e.expr1.Eval(), e.expr2.Eval()
	switch e.op {
	case '-':
		return num1 - num2
	case '+':
		fallthrough
	default:
		return num1 + num2
	}
}

func (e BinaryExpression) String() string {
	return fmt.Sprintf("%s %c %s", e.expr1, e.op, e.expr2)
}
