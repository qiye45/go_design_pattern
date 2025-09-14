package factory

import "testing"

func TestSimpleFactory(t *testing.T) {
	productA := CreateProduct("A")
	if productA.Use() != "Product A" {
		t.Error("Expected 'Product A'")
	}
}

func TestFactoryMethod(t *testing.T) {
	factoryA := &FactoryA{}
	productA := factoryA.CreateProduct()
	if productA.Use() != "Product A" {
		t.Error("Expected 'Product A'")
	}
}

func TestDIContainer(t *testing.T) {
	container := NewContainer()
	container.Register("productA", &ConcreteProductA{})

	product := container.Get("productA").(Product)
	if product.Use() != "Product A" {
		t.Error("Expected 'Product A'")
	}
}
