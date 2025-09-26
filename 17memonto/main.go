package main

import "fmt"

type Memento interface{}

type Game struct {
	hp, mp int
}

type gameMemento struct {
	hp, mp int
}

func (g *Game) Play(hpDelta, mpDelta int) {
	g.hp += hpDelta
	g.mp += mpDelta
}

func (g *Game) Save() Memento {
	return &gameMemento{
		hp: g.hp,
		mp: g.mp,
	}
}

func (g *Game) Load(m Memento) {
	gm := m.(*gameMemento)
	g.hp = gm.hp
	g.mp = gm.mp
}

func (g *Game) Status() {
	fmt.Printf("Current HP:%d, MP:%d\n", g.hp, g.mp)
}

func main() {
	g := &Game{
		hp: 100,
		mp: 100,
	}
	g.Status()
	process := g.Save()
	g.Play(-10, -10)
	g.Status()
	g.Load(process)
	g.Status()
}
