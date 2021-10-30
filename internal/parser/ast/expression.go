package ast

type Expression interface {
	Eval() float64
}
