package builder

// 产品
type House struct {
	foundation string
	walls      string
	roof       string
}

// 建造者接口
type Builder interface {
	BuildFoundation() Builder
	BuildWalls() Builder
	BuildRoof() Builder
	GetResult() *House
}

// 具体建造者
type ConcreteBuilder struct {
	house *House
}

func NewBuilder() *ConcreteBuilder {
	return &ConcreteBuilder{house: &House{}}
}

func (b *ConcreteBuilder) BuildFoundation() Builder {
	b.house.foundation = "concrete foundation"
	return b
}

func (b *ConcreteBuilder) BuildWalls() Builder {
	b.house.walls = "brick walls"
	return b
}

func (b *ConcreteBuilder) BuildRoof() Builder {
	b.house.roof = "tile roof"
	return b
}

func (b *ConcreteBuilder) GetResult() *House {
	return b.house
}

// 指挥者
type Director struct {
	builder Builder
}

func (d *Director) SetBuilder(builder Builder) {
	d.builder = builder
}

func (d *Director) BuildHouse() *House {
	return d.builder.BuildFoundation().BuildWalls().BuildRoof().GetResult()
}
