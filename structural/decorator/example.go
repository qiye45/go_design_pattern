package main

import "fmt"

// Component 接口
type Component interface {
	Operation()
}

// ConcreteComponent 具体实现
type ConcreteComponent struct{}

func (c *ConcreteComponent) Operation() { fmt.Println("基础功能") }

// Decorator 装饰器基类
type Decorator struct {
	component Component
}

func (d *Decorator) Operation() {
	if d.component != nil {
		d.component.Operation()
	}
}

// LogDecorator 具体装饰器1：加日志
type LogDecorator struct{ Decorator }

func (d *LogDecorator) Operation() {
	fmt.Println("日志: 调用前")
	d.Decorator.Operation()
	fmt.Println("日志: 调用后")
}

// AuthDecorator 具体装饰器2：加权限
type AuthDecorator struct{ Decorator }

func (d *AuthDecorator) Operation() {
	fmt.Println("权限检查")
	d.Decorator.Operation()
}

func main() {
	base := &ConcreteComponent{}
	withLog := &LogDecorator{Decorator{base}}
	withAuth := &AuthDecorator{Decorator{withLog}}
	withAuth.Operation()
}
