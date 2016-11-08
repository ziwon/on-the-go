// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.

package eval

type Expr interface {
	Eval(env Env) float64
}

type Var string

type literal float64

type unary struct {
	op rune // one of '+', '-'
	x Expr
}

type binary struct {
	op rune // one of '+', '-', '*'
	x, y Expr
}

type call struct {
	fn string
	args []Expr
}

