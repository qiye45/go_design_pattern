package decorator

import "testing"

func TestDecorator(t *testing.T) {
	component := &ConcreteComponent{}

	// 单层装饰
	decoratorA := NewDecoratorA(component)
	result := decoratorA.Operation()
	expected := "DecoratorA(ConcreteComponent)"
	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}

	// 多层装饰
	decoratorB := NewDecoratorB(decoratorA)
	result = decoratorB.Operation()
	expected = "DecoratorB(DecoratorA(ConcreteComponent))"
	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}
