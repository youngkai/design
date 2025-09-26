package main

import "fmt"

type Week interface {
	Today()
	Next(ctx *DayContext)
}

type DayContext struct {
	today Week
}

func NewDayContext() *DayContext {
	return &DayContext{
		today: &SunDay{},
	}
}

func (d *DayContext) Today() {
	d.today.Today()
}

func (d *DayContext) Next() {
	d.today.Next(d)
}

var _ Week = (*SunDay)(nil)

type SunDay struct{}

func (s *SunDay) Today() {
	fmt.Printf("Sunday\n")
}

func (s *SunDay) Next(ctx *DayContext) {
	ctx.today = &MonDay{}
}

var _ Week = (*MonDay)(nil)

type MonDay struct{}

func (m *MonDay) Today() {
	fmt.Printf("Monday\n")
}

func (m *MonDay) Next(ctx *DayContext) {
	ctx.today = &TueDay{}
}

var _ Week = (*TueDay)(nil)

type TueDay struct{}

func (t *TueDay) Today() {
	fmt.Printf("Tuesday\n")
}

func (t *TueDay) Next(ctx *DayContext) {
	ctx.today = &WedDay{}
}

var _ Week = (*WedDay)(nil)

type WedDay struct{}

func (w *WedDay) Today() {
	fmt.Printf("Wednesday\n")
}

func (w *WedDay) Next(ctx *DayContext) {
	ctx.today = &ThuDay{}
}

var _ Week = (*ThuDay)(nil)

type ThuDay struct{}

func (t *ThuDay) Today() {
	fmt.Printf("Thursday\n")
}

func (t *ThuDay) Next(ctx *DayContext) {
	ctx.today = &FriDay{}
}

var _ Week = (*FriDay)(nil)

type FriDay struct{}

func (f *FriDay) Today() {
	fmt.Printf("Friday\n")
}

func (f *FriDay) Next(ctx *DayContext) {
	ctx.today = &SatDay{}
}

var _ Week = (*SatDay)(nil)

type SatDay struct{}

func (s *SatDay) Today() {
	fmt.Printf("Saturday\n")
}

func (s *SatDay) Next(ctx *DayContext) {
	ctx.today = &SunDay{}
}

func main() {
	ctx := NewDayContext()
	f := func() {
		ctx.Today()
		ctx.Next()
	}
	for i := 0; i < 8; i++ {
		f()
	}
}
