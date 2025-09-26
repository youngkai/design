package main

import "fmt"

//type Subject1 struct {
//	observers []Observer1
//	context   string
//}
//
//func NewSubject() *Subject1 {
//	return &Subject1{
//		observers: make([]Observer1, 0),
//	}
//}
//
//func (s *Subject1) Attach(o Observer1) {
//	s.observers = append(s.observers, o)
//}
//
//func (s *Subject1) notify() {
//	for _, o := range s.observers {
//		o.Update(s)
//	}
//}
//
//func (s *Subject1) UpdateContext(context string) {
//	s.context = context
//	s.notify()
//}
//
//type Observer1 interface {
//	Update(*Subject1)
//}
//
//type Reader struct {
//	name string
//}
//
//func NewReader(name string) *Reader {
//	return &Reader{
//		name: name,
//	}
//}
//
//func (r *Reader) Update(s *Subject1) {
//	fmt.Printf("%s receive %s\n", r.name, s.context)
//}
//
//// 观察者接口
//type Observer interface {
//	Update(message string)
//}
//
//// 被观察者接口
//type Subject interface {
//	Register(o Observer)
//	Unregister(o Observer)
//	Notify(message string)
//}
//
//type Publisher struct {
//	observers []Observer
//}
//
//func (p *Publisher) Register(o Observer) {
//	p.observers = append(p.observers, o)
//}
//
//func (p *Publisher) Unregister(o Observer) {
//	for i, ob := range p.observers {
//		if ob == o {
//			p.observers = append(p.observers[:i], p.observers[i+1:]...)
//		}
//	}
//}
//
//func (p *Publisher) Notify(message string) {
//	for _, ob := range p.observers {
//		ob.Update(message)
//	}
//}
//
//type Subscriber struct {
//	name string
//}
//
//func (s *Subscriber) Update(message string) {
//	fmt.Printf("%s 收到消息: %s\n", s.name, message)
//}

// Observer 观察者接口
type Observer interface {
	Update(temp float64)
}

// Subject 主题接口
type Subject interface {
	RegisterObserver(o Observer)
	RemoveObserver(o Observer)
	NotifyObserver()
}

// WeatherData 具体主题
type WeatherData struct {
	observers   []Observer
	temperature float64
}

func NewWeatherData() *WeatherData {
	return &WeatherData{
		observers: make([]Observer, 0),
	}
}

func (w *WeatherData) RegisterObserver(o Observer) {
	w.observers = append(w.observers, o)
}

func (w *WeatherData) RemoveObserver(o Observer) {
	for i, ob := range w.observers {
		if ob == o {
			w.observers = append(w.observers[:i], w.observers[i+1:]...)
		}
	}
}

func (w *WeatherData) NotifyObserver() {
	for _, ob := range w.observers {
		ob.Update(w.temperature)
	}
}

func (w *WeatherData) SetTemperature(temp float64) {
	fmt.Println("WeatherData: 温度更新为", temp)
	w.temperature = temp
	w.NotifyObserver()
}

// PhoneDisplay 具体观察者
type PhoneDisplay struct {
	id string
}

func (p *PhoneDisplay) Update(temp float64) {
	fmt.Printf("📱 手机显示器[%s]: 当前温度 %.1f°C\n", p.id, temp)
}

type TVDisplay struct{}

func (t *TVDisplay) Update(temp float64) {
	fmt.Printf("📺 电视显示器: 当前温度 %.1f°C\n", temp)
}

func main() {
	//sub := NewSubject()
	//sub.Attach(NewReader("reader1"))
	//sub.Attach(NewReader("reader2"))
	//sub.UpdateContext("new context")
	//
	//news := Publisher{}
	//sub1 := Subscriber{name: "sub1"}
	//sub2 := Subscriber{name: "sub2"}
	//news.Register(&sub1)
	//news.Register(&sub2)
	//news.Notify("new message")
	//news.Notify("xxxxxxxxxxxxxxxxxxxxxx")
	weatherData := NewWeatherData()

	phone1 := &PhoneDisplay{id: "A1"}
	phone2 := &PhoneDisplay{id: "B2"}
	tv := &TVDisplay{}
	weatherData.RegisterObserver(phone1)
	weatherData.RegisterObserver(phone2)
	weatherData.RegisterObserver(tv)
	weatherData.SetTemperature(25.0)

}
