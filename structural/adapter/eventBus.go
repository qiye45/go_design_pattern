package main

import "fmt"

type EventBus struct{ subs map[string][]func(string) }

func NewBus() *EventBus { return &EventBus{subs: map[string][]func(string){}} }
func (b *EventBus) Subscribe(topic string, fn func(string)) {
	b.subs[topic] = append(b.subs[topic], fn)
}
func (b *EventBus) Publish(topic, msg string) {
	for _, fn := range b.subs[topic] {
		fn(msg)
	}
}

func main() {
	bus := NewBus()
	bus.Subscribe("news", func(m string) { fmt.Println("小明:", m) })
	bus.Subscribe("news", func(m string) { fmt.Println("小红:", m) })

	bus.Publish("news", "今晚有更新！")
}
