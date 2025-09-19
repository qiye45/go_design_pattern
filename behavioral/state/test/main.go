package main

import "fmt"

type Order struct {
	num int
}

func (o *Order) change(num int) {
	o.num = num
}

func main() {
	//	改变自己状态
	order := Order{1}
	fmt.Println(order.num)
	order.change(10)
	fmt.Println(order.num)
}
