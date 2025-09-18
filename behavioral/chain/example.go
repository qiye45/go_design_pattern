package main

import "fmt"

// Handler Handler接口
type Handler interface {
	Handle(req int)
	SetNext(h Handler)
}

type Base struct{ next Handler }

func (b *Base) SetNext(h Handler) { b.next = h }
func (b *Base) CallNext(req int) {
	if b.next != nil {
		b.next.Handle(req)
	}
}

// Leader 具体处理器
type Leader struct{ Base }

func (l *Leader) Handle(req int) {
	if req <= 1 {
		fmt.Println("Leader 批准")
	} else {
		l.CallNext(req)
	}
}

type Manager struct{ Base }

func (m *Manager) Handle(req int) {
	if req <= 3 {
		fmt.Println("Manager 批准")
	} else {
		m.CallNext(req)
	}
}

type Boss struct{ Base }

func (b *Boss) Handle(req int) {
	fmt.Println("Boss 批准")
}

func main() {
	leader, manager, boss := &Leader{}, &Manager{}, &Boss{}
	leader.SetNext(manager)
	manager.SetNext(boss)

	leader.Handle(1) // 输出 Leader 批准
	leader.Handle(2)
	leader.Handle(4)
}
