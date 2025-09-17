package main

import "fmt"

// Observer 观察者接口
type Observer interface {
	Update(msg string)
}

// User 具体观察者
type User struct{ name string }

func (u *User) Update(msg string) {
	fmt.Printf("%s 收到通知: %s\n", u.name, msg)
}

// WeChat 被观察者
type WeChat struct{ subs []Observer }

func (w *WeChat) Add(o Observer) { w.subs = append(w.subs, o) }
func (w *WeChat) Notify(msg string) {
	for _, o := range w.subs {
		o.Update(msg)
	}
}

func main() {
	wx := &WeChat{}
	wx.Add(&User{"小明"})
	wx.Add(&User{"小红"})
	wx.Notify("新文章上线啦！")
}
