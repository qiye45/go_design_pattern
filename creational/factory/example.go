package factory

// 产品接口
type Product interface {
	Use() string
}

// 具体产品
type ConcreteProductA struct{}

func (p *ConcreteProductA) Use() string { return "Product A" }

type ConcreteProductB struct{}

func (p *ConcreteProductB) Use() string { return "Product B" }

// 简单工厂
func CreateProduct(productType string) Product {
	switch productType {
	case "A":
		return &ConcreteProductA{}
	case "B":
		return &ConcreteProductB{}
	default:
		return nil
	}
}

// 工厂方法
type Factory interface {
	CreateProduct() Product
}

type FactoryA struct{}

func (f *FactoryA) CreateProduct() Product { return &ConcreteProductA{} }

type FactoryB struct{}

func (f *FactoryB) CreateProduct() Product { return &ConcreteProductB{} }

// DI容器
type Container struct {
	services map[string]interface{}
}

func NewContainer() *Container {
	return &Container{services: make(map[string]interface{})}
}

func (c *Container) Register(name string, service interface{}) {
	c.services[name] = service
}

func (c *Container) Get(name string) interface{} {
	return c.services[name]
}
