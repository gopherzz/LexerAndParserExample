package ast

import "fmt"

type NumberExpression struct {
	value float64
}

func NewNumberExpresion(value float64) NumberExpression {
	return NumberExpression{
		value: value,
	}
}

func (e NumberExpression) Eval() float64 {
	return e.value
}

func (e NumberExpression) String() string {
	return fmt.Sprintf("%.1f", e.value)
}
