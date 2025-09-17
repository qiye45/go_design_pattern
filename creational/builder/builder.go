package main

import "fmt"

// Computer 产品：电脑
type Computer struct {
	CPU    string
	Memory int
	Disk   int
	GPU    string
}

// ComputerBuilder Builder
type ComputerBuilder struct {
	cpu    string
	memory int
	disk   int
	gpu    string
}

// NewComputerBuilder 链式调用设置参数
func NewComputerBuilder() *ComputerBuilder { return &ComputerBuilder{} }
func (b *ComputerBuilder) SetCPU(cpu string) *ComputerBuilder {
	b.cpu = cpu
	return b
}
func (b *ComputerBuilder) SetMemory(m int) *ComputerBuilder {
	b.memory = m
	return b
}
func (b *ComputerBuilder) SetDisk(d int) *ComputerBuilder {
	b.disk = d
	return b
}
func (b *ComputerBuilder) SetGPU(g string) *ComputerBuilder {
	b.gpu = g
	return b
}

// Build 构建最终对象
func (b *ComputerBuilder) Build() Computer {
	fmt.Println("开始构建电脑...")
	if b.cpu == "" {
		b.cpu = "Intel i7"
	}
	if b.memory == 0 {
		b.memory = 16
	}
	if b.disk == 0 {
		b.disk = 512
	}
	if b.gpu == "" {
		b.gpu = "RTX 3080"
	}
	return Computer{CPU: b.cpu, Memory: b.memory, Disk: b.disk, GPU: b.gpu}
}

func main() {
	// 灵活选择参数，链式调用
	pc := NewComputerBuilder().
		SetCPU("Intel i9").
		SetMemory(32).
		Build()

	fmt.Printf("电脑配置: %+v\n", pc)
}
