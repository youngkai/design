package main

import "fmt"

type AbstractMessage interface {
	SendMessage(text, to string)
}

type MessageImplementer interface {
	Send(text, to string)
}

type MessageSMS struct{}

func ViaMessageSMS() MessageImplementer {
	return &MessageSMS{}
}

func (m *MessageSMS) Send(text, to string) {
	fmt.Printf("send %s to %s via SMS\n", text, to)
}

type MessageEmail struct{}

func ViaMessageEmail() MessageImplementer {
	return &MessageEmail{}
}

func (m *MessageEmail) Send(text, to string) {
	fmt.Printf("send %s to %s via Email\n", text, to)
}

type CommonMessage struct {
	method MessageImplementer
}

func NewCommonMessage(method MessageImplementer) *CommonMessage {
	return &CommonMessage{
		method: method,
	}
}

func (c *CommonMessage) Send(text, to string) {
	c.method.Send(text, to)
}

type UrgencyMessage struct {
	method MessageImplementer
}

func NewUrgencyMessage(method MessageImplementer) *UrgencyMessage {
	return &UrgencyMessage{
		method: method,
	}
}

func (u *UrgencyMessage) Send(text, to string) {
	u.method.Send(text, to)
}

func main() {
	c1 := NewCommonMessage(ViaMessageEmail())
	c1.Send("xxxxxxxxxxxx", "tom")
	c2 := NewCommonMessage(ViaMessageSMS())
	c2.Send("xxxxxxxxxxxx", "tom1")
	c3 := NewUrgencyMessage(ViaMessageEmail())
	c3.Send("xxxxxxxxxxxx", "tom2")
	c4 := NewUrgencyMessage(ViaMessageSMS())
	c4.Send("xxxxxxxxxxxx", "tom3")
}
