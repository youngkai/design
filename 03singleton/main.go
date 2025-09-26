package main

import (
	"fmt"
	"sync"
)

type Singleton interface {
	foo() string
}

type aa struct{}

func (a *aa) foo() string {
	return "adsadsadasd"
}

var (
	instance *aa
	once     sync.Once
)

func GetSingleton() Singleton {
	if instance == nil {
		once.Do(func() {
			instance = &aa{}
		})
	}
	return instance
}

func main() {
	res := GetSingleton().foo()
	fmt.Println(res)
}
