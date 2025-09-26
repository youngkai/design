package main

import "fmt"

type Operator interface {
	SetA(int)
	SetB(int)
	Result() int
}

type OperatorFactory interface {
	Create() Operator
}

type OperatorBase struct {
	A int
	B int
}

func (o *OperatorBase) SetA(a int) {
	o.A = a
}

func (o *OperatorBase) SetB(b int) {
	o.B = b
}

type PlusOperatorFactory struct{}

func (p *PlusOperatorFactory) Create() Operator {
	return &PlusOperator{
		OperatorBase: &OperatorBase{},
	}
}

type PlusOperator struct {
	*OperatorBase
}

func (p *PlusOperator) Result() int {
	return p.A + p.B
}

type MinusOperatorFactory struct{}

func (m *MinusOperatorFactory) Create() Operator {
	return &MinusOperator{
		OperatorBase: &OperatorBase{},
	}
}

type MinusOperator struct {
	*OperatorBase
}

func (m *MinusOperator) Result() int {
	return m.A - m.B
}

func main() {
	plus := &PlusOperatorFactory{}
	pOp := plus.Create()
	pOp.SetA(1)
	pOp.SetB(2)
	res1 := pOp.Result()
	fmt.Println(res1)
	minus := &MinusOperatorFactory{}
	mOp := minus.Create()
	mOp.SetA(1)
	mOp.SetB(2)
	res2 := mOp.Result()
	fmt.Println(res2)
}
