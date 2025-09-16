// Package main 演示依赖注入(DI)模式的使用
// 依赖注入是一种设计模式，用于实现控制反转(IoC)
// 它允许对象的依赖关系由外部容器管理，而不是对象自己创建依赖
package main

import (
	"fmt"

	"github.com/qiye45/go_design_pattern/creational/factory/di"
)

// A 结构体依赖于B
// 依赖关系链: A -> B -> C
// 这种链式依赖在实际项目中很常见，比如 Controller -> Service -> Repository
type A struct {
	B *B // A依赖B，通过构造函数注入
}

// NewA 是A的构造函数
// 接收B作为参数，实现依赖注入
// DI容器会自动解析并注入B的实例
func NewA(b *B) *A {
	return &A{
		B: b,
	}
}

// B 结构体依赖于C
type B struct {
	C *C // B依赖C，通过构造函数注入
}

// NewB 是B的构造函数
// 接收C作为参数，DI容器会自动解析并注入C的实例
func NewB(c *C) *B {
	return &B{C: c}
}

// C 结构体是依赖链的底层，不依赖其他对象
type C struct {
	Num int // 简单的数据字段
}

// NewC 是C的构造函数
// 不需要任何依赖，是依赖链的根节点
func NewC() *C {
	return &C{
		Num: 1, // 初始化默认值
	}
}

func main() {
	// 创建DI容器实例
	// 容器负责管理所有对象的生命周期和依赖关系
	container := di.New()

	// 向容器注册构造函数
	// Provide方法会分析构造函数的参数类型，建立依赖关系图
	if err := container.Provide(NewA); err != nil {
		panic(err)
	}
	if err := container.Provide(NewB); err != nil {
		panic(err)
	}
	if err := container.Provide(NewC); err != nil {
		panic(err)
	}

	// 调用函数并自动注入依赖
	// Invoke会分析函数参数，自动创建并注入A的实例
	// 容器会递归解析A的依赖(B)，B的依赖(C)，按正确顺序创建对象
	err := container.Invoke(func(a *A) {
		// 此时a已经完全初始化，包含完整的依赖链 A->B->C
		fmt.Printf("对象A: %+v, C的值: %d\n", a, a.B.C.Num)
	})
	if err != nil {
		panic(err)
	}

	// DI模式的优势:
	// 1. 解耦: 对象不需要知道如何创建依赖
	// 2. 可测试: 容易进行单元测试，可以注入mock对象
	// 3. 灵活: 可以在运行时改变依赖的实现
	// 4. 单一职责: 对象专注于业务逻辑，不关心依赖创建
}
