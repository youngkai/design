package main

import "fmt"

type Api interface {
	Say(name string) string
}

func NewApi(t int) Api {
	if t == 1 {
		return &hiApi{}
	}
	return &helloApi{}
}

type hiApi struct{}

func (*hiApi) Say(name string) string {
	return fmt.Sprintf("Hi, %s", name)
}

type helloApi struct{}

func (*helloApi) Say(name string) string {
	return fmt.Sprintf("Hello, %s", name)
}

func main() {
	factory := NewApi(2)
	res := factory.Say("tom")
	fmt.Println(res)
}
