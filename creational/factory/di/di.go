// Package di 实现了一个简单的依赖注入容器
// 依赖注入(Dependency Injection)是控制反转(IoC)的一种实现方式
// 它通过外部容器来管理对象的创建和依赖关系，而不是对象自己创建依赖
package di

import (
	"fmt"
	"reflect"
)

// Container DI容器，负责管理所有对象的生命周期和依赖关系
// 使用反射机制来分析构造函数的参数类型，自动解析依赖关系
type Container struct {
	// providers 存储类型到构造函数的映射
	// 假设一种类型只能有一个provider提供，避免歧义
	// key: 对象类型, value: 对应的构造函数信息
	providers map[reflect.Type]provider

	// results 缓存已创建的对象实例，实现单例模式
	// 避免重复创建相同类型的对象，提高性能
	// key: 对象类型, value: 对象实例的反射值
	results map[reflect.Type]reflect.Value
}

// provider 封装构造函数的信息
// 包含构造函数本身和它需要的参数类型列表
type provider struct {
	value  reflect.Value  // 构造函数的反射值
	params []reflect.Type // 构造函数参数的类型列表，用于依赖解析
}

// New 创建一个新的DI容器实例
// 初始化providers和results映射表
func New() *Container {
	return &Container{
		providers: map[reflect.Type]provider{},
		results:   map[reflect.Type]reflect.Value{},
	}
}

// isError 判断给定类型是否是error接口类型
// 用于区分构造函数的正常返回值和错误返回值
// Go语言惯例：函数可以返回(result, error)，需要特殊处理error类型
func isError(t reflect.Type) bool {
	// 获取error接口的类型
	errorType := reflect.TypeOf((*error)(nil)).Elem()
	// 检查是否实现了error接口
	return t.Implements(errorType)
}

// Provide 注册对象的构造函数到容器中
// constructor 必须是一个函数，用作对象的工厂方法
// 函数的参数类型会被解析为依赖，返回值类型会被注册为可提供的类型
// 例如: func NewA(b *B) *A 会注册*A类型，依赖*B类型
func (c *Container) Provide(constructor interface{}) error {
	v := reflect.ValueOf(constructor)

	// 只接受函数类型的构造器，确保类型安全
	if v.Kind() != reflect.Func {
		return fmt.Errorf("constructor must be a func")
	}

	vt := v.Type()

	// 解析构造函数的输入参数类型
	// 这些参数类型就是该对象的依赖项
	params := make([]reflect.Type, vt.NumIn())
	for i := 0; i < vt.NumIn(); i++ {
		params[i] = vt.In(i)
	}

	// 解析构造函数的返回值类型
	// 通常包含目标对象类型和可能的error类型
	results := make([]reflect.Type, vt.NumOut())
	for i := 0; i < vt.NumOut(); i++ {
		results[i] = vt.Out(i)
	}

	// 创建provider实例，封装构造函数信息
	provider := provider{
		value:  v,      // 构造函数的反射值
		params: params, // 依赖的参数类型列表
	}

	// 为每个非error返回值类型注册provider
	for _, result := range results {
		// 跳过error类型，不作为可注入的依赖类型
		if isError(result) {
			continue
		}

		// 检查类型是否已经注册，避免重复注册
		if _, ok := c.providers[result]; ok {
			return fmt.Errorf("%s had a provider", result)
		}

		// 注册类型到provider的映射
		c.providers[result] = provider
	}

	return nil
}

// Invoke 执行指定函数，自动注入其所需的依赖
// function 必须是一个函数，其参数会被自动解析并注入
// 这是DI容器的核心入口点，触发整个依赖解析和对象创建过程
func (c *Container) Invoke(function interface{}) error {
	v := reflect.ValueOf(function)

	// 只接受函数类型，确保可以进行参数注入
	if v.Kind() != reflect.Func {
		return fmt.Errorf("function must be a func")
	}

	vt := v.Type()

	// 为函数的每个参数构建对应的依赖对象
	var err error
	params := make([]reflect.Value, vt.NumIn())
	for i := 0; i < vt.NumIn(); i++ {
		// 递归构建参数，这会触发整个依赖链的解析
		params[i], err = c.buildParam(vt.In(i))
		if err != nil {
			return err
		}
	}

	// 调用函数，传入构建好的参数
	v.Call(params)

	return nil
}

// buildParam 递归构建指定类型的参数对象
// 这是DI容器的核心算法，实现了依赖解析的完整流程：
// 1. 检查缓存，避免重复创建（单例模式）
// 2. 查找对应的provider（构造函数）
// 3. 递归构建provider所需的依赖参数
// 4. 调用provider创建对象
// 5. 缓存结果并返回
func (c *Container) buildParam(param reflect.Type) (val reflect.Value, err error) {
	// 步骤1: 检查是否已经创建过该类型的对象（缓存机制）
	// 实现单例模式，确保同一类型只创建一次
	if result, ok := c.results[param]; ok {
		return result, nil
	}

	// 步骤2: 查找该类型对应的构造函数provider
	provider, ok := c.providers[param]
	if !ok {
		return reflect.Value{}, fmt.Errorf("can not found provider: %s", param)
	}

	// 步骤3: 递归构建provider所需的所有依赖参数
	// 这里实现了依赖链的自动解析，例如 A->B->C 会先创建C，再创建B，最后创建A
	params := make([]reflect.Value, len(provider.params))
	for i, p := range provider.params {
		// 递归调用，构建每个依赖参数
		params[i], err = c.buildParam(p)
		if err != nil {
			return reflect.Value{}, err
		}
	}

	// 步骤4: 调用构造函数，创建目标对象
	results := provider.value.Call(params)

	// 步骤5: 处理构造函数的返回值
	for _, result := range results {
		// 检查是否有错误返回
		if isError(result.Type()) && !result.IsNil() {
			return reflect.Value{}, fmt.Errorf("%+v call err: %+v", provider, result)
		}

		// 缓存非error类型的返回值，供后续使用
		if !isError(result.Type()) && !result.IsNil() {
			c.results[result.Type()] = result
		}
	}

	// 返回缓存中的对象实例
	return c.results[param], nil
}
