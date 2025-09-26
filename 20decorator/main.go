package main

import "fmt"

type Component interface {
	Calc() int
}

type ConcreteComponent struct{}

func (c *ConcreteComponent) Calc() int {
	return 0
}

type MulDecorator struct {
	Component
	num int
}

func WrapMulDecorator(c Component, num int) Component {
	return &MulDecorator{
		Component: c,
		num:       num,
	}
}

func (m *MulDecorator) Calc() int {
	return m.Component.Calc() * m.num
}

type AddDecorator struct {
	Component
	num int
}

func WrapAddDecorator(c Component, num int) Component {
	return &AddDecorator{
		Component: c,
		num:       num,
	}
}

func (a *AddDecorator) Calc() int {
	return a.Component.Calc() + a.num
}

func main() {
	var c Component = &ConcreteComponent{}
	c = WrapAddDecorator(c, 1)
	c = WrapMulDecorator(c, 3)
	res := c.Calc()
	fmt.Println(res)
}
