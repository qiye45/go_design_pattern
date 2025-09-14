package chain

import "testing"

func TestChainOfResponsibility(t *testing.T) {
	handlerA := &ConcreteHandlerA{}
	handlerB := &ConcreteHandlerB{}

	handlerA.SetNext(handlerB)

	// 测试A处理
	result := handlerA.Handle("A")
	if result != "Handled by A" {
		t.Errorf("Expected 'Handled by A', got %s", result)
	}

	// 测试B处理
	result = handlerA.Handle("B")
	if result != "Handled by B" {
		t.Errorf("Expected 'Handled by B', got %s", result)
	}

	// 测试无法处理
	result = handlerA.Handle("C")
	if result != "" {
		t.Errorf("Expected empty string, got %s", result)
	}
}

func TestMiddlewareChain(t *testing.T) {
	chain := &MiddlewareChain{}
	chain.Use(LoggingMiddleware)
	chain.Use(AuthMiddleware)

	result := chain.Execute("test")
	expected := "Auth(test)"
	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}
