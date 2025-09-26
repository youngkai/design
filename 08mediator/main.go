package main

import (
	"fmt"
	"strings"
)

type CDDriver struct {
	Data string
}

func (c *CDDriver) ReadData() {
	c.Data = "music,image"
	fmt.Printf("CDDriver: reading data %s\n", c.Data)
	GetMediatorInstance().changed(c)
}

type CPU struct {
	Video string
	Sound string
}

func (c *CPU) Process(data string) {
	sp := strings.Split(data, ",")
	c.Video = sp[0]
	c.Sound = sp[1]
	fmt.Printf("CPU: split data with Sound %s, Video %s\n", c.Sound, c.Video)
	GetMediatorInstance().changed(c)
}

type VideoCard struct {
	Data string
}

func (v *VideoCard) Display(data string) {
	v.Data = data
	fmt.Printf("SoundCard: play %s\n", v.Data)
	GetMediatorInstance().changed(v)
}

type SoundCard struct {
	Data string
}

func (s *SoundCard) Display(data string) {
	s.Data = data
	fmt.Printf("SoundCard: play %s\n", s.Data)
	GetMediatorInstance().changed(s)
}

type Mediator struct {
	CD    *CDDriver
	CPU   *CPU
	Video *VideoCard
	Sound *SoundCard
}

var mediator *Mediator

func GetMediatorInstance() *Mediator {
	if mediator == nil {
		mediator = &Mediator{}
	}
	return mediator
}

func (m *Mediator) changed(i interface{}) {
	switch inst := i.(type) {
	case *CDDriver:
		m.CPU.Process(inst.Data)
	case *CPU:
		m.Video.Display(inst.Video)
		m.Sound.Display(inst.Sound)
	}
}

type ChatRoom interface {
	SendMessage(user, message string)
}

type SimpleChatRoom struct {
	users map[string]*User
}

func NewSimpleChatRoom() *SimpleChatRoom {
	return &SimpleChatRoom{
		users: make(map[string]*User),
	}
}

func (s *SimpleChatRoom) Register(u *User) {
	s.users[u.name] = u
	u.room = s
}

func (s *SimpleChatRoom) SendMessage(user, message string) {
	for name, u := range s.users {
		if name != user {
			u.Receive(user, message)
		}
	}
}

type User struct {
	name string
	room ChatRoom
}

func NewUser(name string) *User {
	return &User{
		name: name,
	}
}

func (u *User) Send(msg string) {
	if u.room != nil {
		fmt.Printf("[%s] 发送: %s\n", u.name, msg)
		u.room.SendMessage(u.name, msg)
	}
}

func (u *User) Receive(from, msg string) {
	fmt.Printf("[%s] 收到来自 %s 的消息: %s\n", u.name, from, msg)
}

func main() {
	a := GetMediatorInstance()
	a.CD = &CDDriver{}
	a.CPU = &CPU{}
	a.Video = &VideoCard{}
	a.Sound = &SoundCard{}
	a.CD.ReadData()

	chat := NewSimpleChatRoom()
	tom := NewUser("tom")
	alice := NewUser("alice")
	chat.Register(tom)
	chat.Register(alice)
	tom.Send("hello")
	alice.Send("world")
}
