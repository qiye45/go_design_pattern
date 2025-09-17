package main

import "fmt"

// Sender 实现层接口
type Sender interface {
	Send(msg string)
}

// EmailSender 不同实现
type EmailSender struct{}

func (e *EmailSender) Send(msg string) { fmt.Println("邮件发送:", msg) }

type SmsSender struct{}

func (s *SmsSender) Send(msg string) { fmt.Println("短信发送:", msg) }

// Message 抽象层
type Message struct {
	sender Sender
}

func (m *Message) SetSender(s Sender) { m.sender = s }

type NormalMessage struct{ Message }

func (m *NormalMessage) Send(msg string) { m.sender.Send("[普通]" + msg) }

type UrgentMessage struct{ Message }

func (m *UrgentMessage) Send(msg string) { m.sender.Send("[加急]" + msg) }

func main() {
	normal := &NormalMessage{}
	normal.SetSender(&EmailSender{})
	normal.Send("开会通知")
	normal.SetSender(&SmsSender{})
	normal.Send("开会通知 短信")

	urgent := &UrgentMessage{}
	urgent.SetSender(&SmsSender{})
	urgent.Send("服务器宕机！")
}
