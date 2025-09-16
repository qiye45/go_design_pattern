# 依赖注入(DI)实际开发示例

本目录包含了两个依赖注入的示例：

## 1. 基础示例 (example.go)
演示了简单的依赖链 A -> B -> C，适合理解DI的基本概念。

## 2. 实际开发示例 (realistic_example.go)
模拟了一个真实的Web应用架构，展示了DI在实际项目中的应用。

### 架构层次
```
App
├── UserController (控制器层)
│   └── UserService (业务逻辑层)
│       └── UserRepository (数据访问层)
│           └── Database (数据库层)
└── Logger (日志服务)
```

### 依赖关系
- **App** 依赖 UserController 和 Logger
- **UserController** 依赖 UserService
- **UserService** 依赖 UserRepository
- **UserRepository** 依赖 Database 接口
- **MySQLDatabase** 实现 Database 接口

### 运行示例
```bash
go run realistic_example.go
```

### 预期输出
```
[INFO] 应用程序启动
连接MySQL数据库: localhost:3306
处理GET /users/1请求
获取用户资料: 用户1
返回用户信息: &{ID:1 Name:用户1 Email:user1@example.com}
处理POST /users请求: name=张三, email=zhangsan@example.com
保存用户到数据库: &{ID:100 Name:张三 Email:zhangsan@example.com}
用户注册成功: 张三
用户创建成功: &{ID:100 Name:张三 Email:zhangsan@example.com}
[INFO] 应用程序运行中...

=== DI模式在实际开发中的优势 ===
1. 分层架构清晰: Controller -> Service -> Repository -> Database
2. 依赖解耦: 每层只依赖接口，不依赖具体实现
3. 易于测试: 可以轻松注入Mock对象进行单元测试
4. 配置集中: 所有依赖关系在main函数中统一配置
5. 扩展性好: 新增功能只需注册新的构造函数
```

## 3. 单元测试示例 (realistic_example_test.go)
展示了如何使用DI进行单元测试，通过注入Mock对象来隔离测试。

### 测试特点
- **MockDatabase**: 模拟数据库操作，避免依赖真实数据库
- **单元测试**: 测试单个组件的功能
- **集成测试**: 测试整个应用流程
- **错误场景**: 测试异常情况的处理

### 运行测试
```bash
go test -v
```

## DI模式的优势

### 1. 解耦合
- 每个组件只依赖接口，不依赖具体实现
- 可以轻松替换实现，如从MySQL切换到PostgreSQL

### 2. 可测试性
- 可以注入Mock对象进行单元测试
- 测试时不需要真实的数据库连接
- 可以模拟各种错误场景

### 3. 配置集中化
- 所有依赖关系在main函数中统一配置
- 便于管理和修改依赖关系

### 4. 单一职责
- 每个组件专注于自己的业务逻辑
- 不需要关心依赖对象的创建和管理

### 5. 扩展性
- 新增功能只需要注册新的构造函数
- 符合开闭原则：对扩展开放，对修改关闭

## 实际应用场景

### Web框架
- Gin + DI: 注入数据库连接、配置、服务等
- Echo + DI: 管理中间件、处理器的依赖

### 微服务
- 服务间依赖管理
- 配置管理
- 监控和日志服务注入

### 数据库操作
- ORM框架集成
- 事务管理
- 连接池管理

### 第三方服务
- Redis客户端
- 消息队列
- 外部API客户端

这个示例展示了DI模式在实际Go项目中的最佳实践，帮助开发者构建可维护、可测试的应用程序。