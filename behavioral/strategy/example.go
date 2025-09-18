package main

import "fmt"

// PaymentStrategy 支付策略接口
type PaymentStrategy interface {
	Pay(amount float64)
}

// Alipay 支付宝支付
type Alipay struct{}

func (a *Alipay) Pay(amount float64) {
	fmt.Printf("Paying %.2f using Alipay.\n", amount)
}

// WeChatPay 微信支付
type WeChatPay struct{}

func (w *WeChatPay) Pay(amount float64) {
	fmt.Printf("Paying %.2f using WeChat Pay.\n", amount)
}

// CreditCardPay 信用卡支付
type CreditCardPay struct{}

func (c *CreditCardPay) Pay(amount float64) {
	fmt.Printf("Paying %.2f using Credit Card.\n", amount)
}

// PaymentContext 上下文：支付上下文
type PaymentContext struct {
	Strategy PaymentStrategy
}

func (p *PaymentContext) SetStrategy(strategy PaymentStrategy) {
	p.Strategy = strategy
}

func (p *PaymentContext) ExecutePayment(amount float64) {
	p.Strategy.Pay(amount)
}

func main() {
	// 创建支付上下文
	context := &PaymentContext{}

	// 用户选择支付宝支付
	context.SetStrategy(&Alipay{})
	context.ExecutePayment(100.0)

	// 用户选择微信支付
	context.SetStrategy(&WeChatPay{})
	context.ExecutePayment(200.0)

	// 用户选择信用卡支付
	context.SetStrategy(&CreditCardPay{})
	context.ExecutePayment(300.0)
}
