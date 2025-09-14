package observer

import (
	"testing"
	"time"
)

func TestObserver(t *testing.T) {
	subject := &ConcreteSubject{}
	observer1 := &ConcreteObserver{name: "observer1"}
	observer2 := &ConcreteObserver{name: "observer2"}

	subject.Attach(observer1)
	subject.Attach(observer2)

	subject.Notify("test data")

	if observer1.GetData() != "test data" {
		t.Error("Observer1 not updated")
	}
	if observer2.GetData() != "test data" {
		t.Error("Observer2 not updated")
	}
}

func TestEventBus(t *testing.T) {
	bus := NewEventBus()
	received := make(chan string, 1)

	bus.Subscribe("test", func(data interface{}) {
		received <- data.(string)
	})

	bus.Publish("test", "hello")

	select {
	case msg := <-received:
		if msg != "hello" {
			t.Errorf("Expected 'hello', got %s", msg)
		}
	case <-time.After(time.Second):
		t.Error("Timeout waiting for event")
	}
}
