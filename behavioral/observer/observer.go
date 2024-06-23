package observer

import "fmt"

type Observer interface {
	Update(stock)
}

type stock struct {
	name      string
	price     float32
	observers []Observer
}

func (s *stock) Subscribe(o Observer) {
	s.observers = append(s.observers, o)
}

func (s *stock) Unsubscribe(observer Observer) {
	for i, o := range s.observers {
		if o == observer {
			s.observers = append(s.observers[:i], s.observers[i+1:]...)
		}
	}
}

func (s *stock) SetPrice(p float32) {
	s.price = p
	s.Notify()
}

func (s *stock) Notify() {
	for _, o := range s.observers {
		o.Update(*s)
	}
}

type investor struct{}

func (i investor) Update(s stock) {
	fmt.Printf("Update received – %s: %.2f", s.name, s.price)
}
