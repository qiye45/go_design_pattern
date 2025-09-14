package decorator

import "fmt"

// 组件接口
type Component interface {
	Operation() string
}

// 具体组件
type ConcreteComponent struct{}

func (c *ConcreteComponent) Operation() string {
	return "ConcreteComponent"
}

// 装饰器基类
type BaseDecorator struct {
	component Component
}

func (d *BaseDecorator) Operation() string {
	return d.component.Operation()
}

// 具体装饰器A
type ConcreteDecoratorA struct {
	BaseDecorator
}

func NewDecoratorA(component Component) *ConcreteDecoratorA {
	return &ConcreteDecoratorA{
		BaseDecorator: BaseDecorator{component: component},
	}
}

func (d *ConcreteDecoratorA) Operation() string {
	return fmt.Sprintf("DecoratorA(%s)", d.BaseDecorator.Operation())
}

// 具体装饰器B
type ConcreteDecoratorB struct {
	BaseDecorator
}

func NewDecoratorB(component Component) *ConcreteDecoratorB {
	return &ConcreteDecoratorB{
		BaseDecorator: BaseDecorator{component: component},
	}
}

func (d *ConcreteDecoratorB) Operation() string {
	return fmt.Sprintf("DecoratorB(%s)", d.BaseDecorator.Operation())
}
