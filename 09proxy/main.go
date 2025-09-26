package main

import "fmt"

type Subject interface {
	Do() string
}

type RealSubject struct{}

func (r *RealSubject) Do() string {
	return "real"
}

type Proxy struct {
	real *RealSubject
}

func (p *Proxy) Do() string {
	var res string
	res += "before:"
	res += p.real.Do()
	res += ":after"
	return res
}

func main() {
	a := &RealSubject{}
	b := &Proxy{real: a}
	res := b.Do()
	fmt.Println(res)
}
