package strategy

import "testing"

func TestStrategy(t *testing.T) {
	context := &Context{}

	// 测试加法策略
	context.SetStrategy(&AddStrategy{})
	if result := context.ExecuteStrategy(5, 3); result != 8 {
		t.Errorf("Expected 8, got %d", result)
	}

	// 测试减法策略
	context.SetStrategy(&SubtractStrategy{})
	if result := context.ExecuteStrategy(5, 3); result != 2 {
		t.Errorf("Expected 2, got %d", result)
	}
}

func TestPaymentStrategy(t *testing.T) {
	paymentContext := &PaymentContext{}

	// 测试信用卡支付
	paymentContext.SetPaymentStrategy(&CreditCardPayment{cardNumber: "1234"})
	result := paymentContext.Pay(100.0)
	if result != "Paid with credit card" {
		t.Errorf("Expected 'Paid with credit card', got %s", result)
	}

	// 测试PayPal支付
	paymentContext.SetPaymentStrategy(&PayPalPayment{email: "test@example.com"})
	result = paymentContext.Pay(100.0)
	if result != "Paid with PayPal" {
		t.Errorf("Expected 'Paid with PayPal', got %s", result)
	}
}
