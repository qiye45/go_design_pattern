package strategy

// 策略接口
type Strategy interface {
	Execute(a, b int) int
}

// 具体策略
type AddStrategy struct{}

func (s *AddStrategy) Execute(a, b int) int { return a + b }

type SubtractStrategy struct{}

func (s *SubtractStrategy) Execute(a, b int) int { return a - b }

type MultiplyStrategy struct{}

func (s *MultiplyStrategy) Execute(a, b int) int { return a * b }

// 上下文
type Context struct {
	strategy Strategy
}

func (c *Context) SetStrategy(strategy Strategy) {
	c.strategy = strategy
}

func (c *Context) ExecuteStrategy(a, b int) int {
	return c.strategy.Execute(a, b)
}

// 支付策略示例
type PaymentStrategy interface {
	Pay(amount float64) string
}

type CreditCardPayment struct {
	cardNumber string
}

func (p *CreditCardPayment) Pay(amount float64) string {
	return "Paid with credit card"
}

type PayPalPayment struct {
	email string
}

func (p *PayPalPayment) Pay(amount float64) string {
	return "Paid with PayPal"
}

type PaymentContext struct {
	strategy PaymentStrategy
}

func (p *PaymentContext) SetPaymentStrategy(strategy PaymentStrategy) {
	p.strategy = strategy
}

func (p *PaymentContext) Pay(amount float64) string {
	return p.strategy.Pay(amount)
}
