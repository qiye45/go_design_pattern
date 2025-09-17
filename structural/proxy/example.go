package main

import "fmt"

// Service 接口
type Service interface {
	DoSomething()
}

// RealService 真实对象
type RealService struct{}

func (r *RealService) DoSomething() {
	fmt.Println("执行真正的业务逻辑")
}

// Proxy 代理对象
type Proxy struct {
	real *RealService
}

func (p *Proxy) DoSomething() {
	fmt.Println("代理：先检查权限")
	p.real.DoSomething()
	fmt.Println("代理：记录日志")
}

func main() {
	service := &Proxy{real: &RealService{}}
	service.DoSomething()
}
