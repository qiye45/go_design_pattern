package main

import "fmt"

// Target 目标接口 (新系统需要的)
type Target interface {
	Request()
}

// Adaptee 被适配者 (老系统/第三方库)
type Adaptee struct{}

func (a *Adaptee) SpecificRequest() {
	fmt.Println("老接口的调用")
}

// Adapter 适配器
type Adapter struct {
	adaptee *Adaptee
}

func (ad *Adapter) Request() {
	fmt.Print("适配器转换 -> ")
	ad.adaptee.SpecificRequest()
}

func main() {
	adaptee := &Adaptee{}
	adapter := &Adapter{adaptee}
	adapter.Request()
}
