package observer

import "sync"

// 观察者接口
type Observer interface {
	Update(data interface{})
}

// 主题接口
type Subject interface {
	Attach(Observer)
	Detach(Observer)
	Notify(data interface{})
}

// 具体主题
type ConcreteSubject struct {
	observers []Observer
	mu        sync.RWMutex
}

func (s *ConcreteSubject) Attach(o Observer) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.observers = append(s.observers, o)
}

func (s *ConcreteSubject) Detach(o Observer) {
	s.mu.Lock()
	defer s.mu.Unlock()
	for i, observer := range s.observers {
		if observer == o {
			s.observers = append(s.observers[:i], s.observers[i+1:]...)
			break
		}
	}
}

func (s *ConcreteSubject) Notify(data interface{}) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	for _, observer := range s.observers {
		observer.Update(data)
	}
}

// 具体观察者
type ConcreteObserver struct {
	name string
	data interface{}
}

func (o *ConcreteObserver) Update(data interface{}) {
	o.data = data
}

func (o *ConcreteObserver) GetData() interface{} {
	return o.data
}

// EventBus实现
type EventBus struct {
	handlers map[string][]func(interface{})
	mu       sync.RWMutex
}

func NewEventBus() *EventBus {
	return &EventBus{
		handlers: make(map[string][]func(interface{})),
	}
}

func (eb *EventBus) Subscribe(event string, handler func(interface{})) {
	eb.mu.Lock()
	defer eb.mu.Unlock()
	eb.handlers[event] = append(eb.handlers[event], handler)
}

func (eb *EventBus) Publish(event string, data interface{}) {
	eb.mu.RLock()
	defer eb.mu.RUnlock()
	for _, handler := range eb.handlers[event] {
		go handler(data)
	}
}
