# Go设计模式

本仓库包含了23种经典设计模式的Go语言实现，每个模式都有详细的示例代码和说明文档。

## 设计原则

- **单一职责原则(SRP)**: 一个类应该只有一个引起它变化的原因
- **开闭原则(OCP)**: 软件实体应该对扩展开放，对修改关闭
- **里氏替换原则(LSP)**: 子类型必须能够替换掉它们的父类型
- **接口隔离原则(ISP)**: 不应该强迫客户依赖于它们不用的方法
- **依赖倒置原则(DIP)**: 高层模块不应该依赖低层模块，两者都应该依赖抽象

## 设计模式分类

### 创建型模式 (Creational Patterns)
- [单例模式 (Singleton)](./creational/singleton/) ✅
- [工厂模式 (Factory)](./creational/factory/) ✅
- [建造者模式 (Builder)](./creational/builder/) ✅
- [原型模式 (Prototype)](./creational/prototype/) 不常用 ❌

### 结构型模式 (Structural Patterns)
- [代理模式 (Proxy)](structural/proxy/) ✅
- [桥接模式 (Bridge)](./structural/bridge/) ✅
- [装饰器模式 (Decorator)](./structural/decorator/) ✅
- [适配器模式 (Adapter)](./structural/adapter/) ✅
- [门面模式 (Facade)](./structural/facade/) 不常用 ❌
- [组合模式 (Composite)](./structural/composite/) 不常用 ❌
- [享元模式 (Flyweight)](./structural/flyweight/) 不常用 ❌

### 行为型模式 (Behavioral Patterns)
- [观察者模式 (Observer)](./behavioral/observer/) ✅
- [模板模式 (Template Method)](./behavioral/template/) ✅
- [策略模式 (Strategy)](./behavioral/strategy/) ✅
- [职责链模式 (Chain of Responsibility)](./behavioral/chain/) ✅
- [状态模式 (State)](./behavioral/state/) ✅
- [迭代器模式 (Iterator)](./behavioral/iterator/) ✅
- [访问者模式 (Visitor)](./behavioral/visitor/) 不常用 ❌
- [备忘录模式 (Memento)](./behavioral/memento/) 不常用 ❌
- [命令模式 (Command)](./behavioral/command/) 不常用 ❌
- [解释器模式 (Interpreter)](./behavioral/interpreter/) 不常用 ❌
- [中介模式 (Mediator)](./behavioral/mediator/) 不常用 ❌

## 项目结构

```
go_design_pattern/
├── creational/          # 创建型模式
│   ├── singleton/       # 单例模式
│   ├── factory/         # 工厂模式
│   ├── builder/         # 建造者模式
│   └── prototype/       # 原型模式
├── structural/          # 结构型模式
│   ├── proxy/           # 代理模式
│   ├── bridge/          # 桥接模式
│   ├── decorator/       # 装饰器模式
│   ├── adapter/         # 适配器模式
│   ├── facade/          # 门面模式
│   ├── composite/       # 组合模式
│   └── flyweight/       # 享元模式
└── behavioral/          # 行为型模式
    ├── observer/        # 观察者模式
    ├── template/        # 模板模式
    ├── strategy/        # 策略模式
    ├── chain/           # 职责链模式
    ├── state/           # 状态模式
    ├── iterator/        # 迭代器模式
    ├── visitor/         # 访问者模式
    ├── memento/         # 备忘录模式
    ├── command/         # 命令模式
    ├── interpreter/     # 解释器模式
    └── mediator/        # 中介模式
```

## 使用说明

每个设计模式目录包含：
- `README.md` - 模式说明和使用场景
- `example.go` - 示例代码实现
- `example_test.go` - 单元测试

## 运行示例

```bash
# 运行特定模式的测试
go test ./creational/singleton/

# 运行所有测试
go test ./...
```

## 贡献

欢迎提交PR来完善代码实现和文档说明。

## 许可证

MIT License