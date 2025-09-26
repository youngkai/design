package main

import "fmt"

type Manage interface {
	HaveRight(money int) bool
	HandleFeeRequest(name string, money int) bool
}

type RequestChain struct {
	Manage
	successor *RequestChain
}

func (r *RequestChain) SetSuccessor(m *RequestChain) {
	r.successor = m
}

func (r *RequestChain) HandleFeeRequest(name string, money int) bool {
	if r.Manage.HaveRight(money) {
		return r.Manage.HandleFeeRequest(name, money)
	}
	if r.successor != nil {
		return r.successor.HandleFeeRequest(name, money)
	}
	return false
}

func (r *RequestChain) HaveRight(money int) bool {
	return true
}

type ProjectManager struct{}

func NewProjectManager() *RequestChain {
	return &RequestChain{
		Manage: &ProjectManager{},
	}
}

func (p *ProjectManager) HaveRight(money int) bool {
	return money < 500
}

func (p *ProjectManager) HandleFeeRequest(name string, money int) bool {
	if name == "bob" {
		fmt.Printf("Project manager permit %s %d fee request\n", name, money)
		return true
	}
	fmt.Printf("Project manager don't permit %s %d fee request\n", name, money)
	return false
}

type DepManager struct{}

func NewDepManager() *RequestChain {
	return &RequestChain{
		Manage: &DepManager{},
	}
}

func (d *DepManager) HaveRight(money int) bool {
	return money < 5000
}

func (d *DepManager) HandleFeeRequest(name string, money int) bool {
	if name == "tom" {
		fmt.Printf("Dep manager permit %s %d fee request\n", name, money)
		return true
	}
	fmt.Printf("Dep manager don't permit %s %d fee request\n", name, money)
	return false
}

type GeneralManager struct{}

func NewGeneralManager() *RequestChain {
	return &RequestChain{
		Manage: &GeneralManager{},
	}
}

func (g *GeneralManager) HaveRight(money int) bool {
	return true
}

func (g *GeneralManager) HandleFeeRequest(name string, money int) bool {
	if name == "ada" {
		fmt.Printf("General manager permit %s %d fee request\n", name, money)
		return true
	}
	fmt.Printf("General manager don't permit %s %d fee request\n", name, money)
	return false
}

func main() {
	c1 := NewProjectManager()
	c2 := NewDepManager()
	c3 := NewGeneralManager()

	c1.SetSuccessor(c2)
	c2.SetSuccessor(c3)

	var c Manage = c1
	c.HandleFeeRequest("bob", 400)
	c.HandleFeeRequest("tom", 1400)
	c.HandleFeeRequest("ada", 10000)
	c.HandleFeeRequest("floar", 400)
}
