package main

import "fmt"

// State 状态接口
type State interface {
	Next(*Order)
	Show()
}

// Created 具体状态
type Created struct{}

func (s *Created) Next(o *Order) { o.state = &Paid{} }
func (s *Created) Show()         { fmt.Println("订单已创建") }

type Paid struct{}

func (s *Paid) Next(o *Order) { o.state = &Shipped{} }
func (s *Paid) Show()         { fmt.Println("订单已支付") }

type Shipped struct{}

func (s *Shipped) Next(o *Order) { o.state = &Done{} }
func (s *Shipped) Show()         { fmt.Println("订单已发货") }

type Done struct{}

func (s *Done) Next(o *Order) {}
func (s *Done) Show()         { fmt.Println("订单完成") }

// Order 上下文
type Order struct{ state State }

func (o *Order) Next() { o.state.Next(o) }
func (o *Order) Show() { o.state.Show() }

func main() {
	order := &Order{state: &Created{}}

	order.Show() // 订单已创建
	order.Next()

	order.Show() // 订单已支付
	order.Next()

	order.Show() // 订单已发货
	order.Next()

	order.Show() // 订单完成
}
