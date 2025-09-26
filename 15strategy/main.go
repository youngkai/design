package main

import "fmt"

type Payment struct {
	context  *PaymentContext
	strategy PaymentStrategy
}

type PaymentContext struct {
	Name, CardId string
	Money        int
}

func NewPayment(name, cardId string, money int, strategy PaymentStrategy) *Payment {
	return &Payment{
		context: &PaymentContext{
			Name:   name,
			CardId: cardId,
			Money:  money,
		},
		strategy: strategy,
	}
}

func (p *Payment) Pay() {
	p.strategy.Pay(p.context)
}

type PaymentStrategy interface {
	Pay(ctx *PaymentContext)
}

type Cash struct{}

func (c *Cash) Pay(ctx *PaymentContext) {
	fmt.Printf("Pay $%d to %s by cash\n", ctx.Money, ctx.Name)
}

type Bank struct{}

func (b *Bank) Pay(ctx *PaymentContext) {
	fmt.Printf("Pay $%d to %s by bank account %s\n", ctx.Money, ctx.Name, ctx.CardId)
}

func main() {
	payment := NewPayment("Ada", "", 123, &Cash{})
	payment.Pay()
	payment = NewPayment("xxxx", "123", 10, &Bank{})
	payment.Pay()
}
