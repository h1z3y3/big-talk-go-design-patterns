package observer

import (
	"fmt"
)

type Subject struct {
	observers []Observer
	context   string
}

func NewSubject() *Subject {
	return &Subject{
		observers: make([]Observer, 0),
	}
}

func (s *Subject) register(o Observer) {
	s.observers = append(s.observers, o)
}

func (s *Subject) notify() {
	for _, o := range s.observers {
		o.Update(s)
	}
}

func (s *Subject) UpdateContext(ctx string) {
	s.context = ctx
	s.notify()
}

type Observer interface {
	Update(*Subject)
}

type ConcreteObserver struct {
	name string
}

func NewConcreteObserver(name string) *ConcreteObserver {
	return &ConcreteObserver{name: name}
}

func (c *ConcreteObserver) Update(s *Subject) {
	fmt.Printf("%s receive %s\n", c.name, s.context)
}
