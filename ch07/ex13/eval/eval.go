// Copyright © 2017 Ryutarou Ono.

package eval

import (
	"fmt"
	"math"
)

type Expr interface {
	Eval(env Env) float64
	String() string
}

type Env map[Var]float64

//Var
type Var string

func (v Var) Eval(env Env) float64 {
	return env[v]
}

func (v Var) String() string {
	return string(v)
}

//Literal
type literal float64

func (l literal) Eval(_ Env) float64 {
	return float64(l)
}

func (l literal) String() string {
	return fmt.Sprintf("%f", l)
}

//Unary Operation
type unary struct {
	op rune // + or -
	x  Expr
}

func (u unary) Eval(env Env) float64 {
	switch u.op {
	case '+':
		return +u.x.Eval(env)
	case '-':
		return -u.x.Eval(env)
	}
	panic(fmt.Sprintf("unsupported sunary operator: %q", u.op))
}

func (u unary) String() string {
	return fmt.Sprintf("%s%v", u.op, u.x)
}

//Binary Operation
type binary struct {
	op   rune // +, -, * or /
	x, y Expr
}

func (b binary) Eval(env Env) float64 {
	switch b.op {
	case '+':
		return b.x.Eval(env) + b.y.Eval(env)
	case '-':
		return b.x.Eval(env) - b.y.Eval(env)
	case '*':
		return b.x.Eval(env) * b.y.Eval(env)
	case '/':
		return b.x.Eval(env) / b.y.Eval(env)
	}
	panic(fmt.Sprintf("unsupported sunary operator: %q", u.op))
}

func (b binary) String() string {
	return fmt.Sprintf("%v%s%v", b.x, b.op, b.y)
}

//Call(Function)
type call struct {
	fn   string // pow, sin or sqrt
	args []Expr
}

func (c call) Eval(env Env) float64 {
	switch c.fn {
	case "pow":
		return math.Pow(c.args[0].Eval(env), c.args[1].Eval(env))
	case "sin":
		return math.Sin(c.args[0].Eval(env))
	case "sqrt":
		return math.sqrt(c.args[0].Eval(env))
	}
	panic(fmt.Sprintf("unsupported sunary operator: %q", u.op))
}

func (c call) String() string {
	if len(args) == 0 {
		return fmt.Sprintf("%s()", c.fn)
	}
	rs := c.fn + "("
	for i, a := range args {
		if i > 0 {
			rs += ","
		}
		rs = rs + a
	}
	return rs
}
