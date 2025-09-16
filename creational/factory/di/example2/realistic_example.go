// Package main 演示依赖注入在实际Web应用中的使用场景
// 模拟一个用户管理系统: Controller -> Service -> Repository -> Database
package main

import (
	"fmt"
	"log"

	"github.com/qiye45/go_design_pattern/creational/factory/di"
)

// User 用户实体
type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

// Database 数据库接口，便于测试时mock
type Database interface {
	Connect() error
	FindUserByID(id int) (*User, error)
	SaveUser(user *User) error
}

// MySQLDatabase MySQL数据库实现
type MySQLDatabase struct {
	Host string
	Port int
}

// NewMySQLDatabase MySQL数据库构造函数
func NewMySQLDatabase() Database {
	return &MySQLDatabase{
		Host: "localhost",
		Port: 3306,
	}
}

func (db *MySQLDatabase) Connect() error {
	fmt.Printf("连接MySQL数据库: %s:%d\n", db.Host, db.Port)
	return nil
}

func (db *MySQLDatabase) FindUserByID(id int) (*User, error) {
	// 模拟数据库查询
	return &User{
		ID:    id,
		Name:  fmt.Sprintf("用户%d", id),
		Email: fmt.Sprintf("user%d@example.com", id),
	}, nil
}

func (db *MySQLDatabase) SaveUser(user *User) error {
	fmt.Printf("保存用户到数据库: %+v\n", user)
	return nil
}

// UserRepository 用户数据访问层
type UserRepository struct {
	db Database // 依赖数据库接口
}

// NewUserRepository 用户仓库构造函数
func NewUserRepository(db Database) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) GetUser(id int) (*User, error) {
	return r.db.FindUserByID(id)
}

func (r *UserRepository) CreateUser(user *User) error {
	return r.db.SaveUser(user)
}

// UserService 用户业务逻辑层
type UserService struct {
	repo *UserRepository // 依赖用户仓库
}

// NewUserService 用户服务构造函数
func NewUserService(repo *UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) GetUserProfile(id int) (*User, error) {
	// 业务逻辑：获取用户信息
	user, err := s.repo.GetUser(id)
	if err != nil {
		return nil, err
	}

	// 可以在这里添加更多业务逻辑，如权限检查、数据转换等
	fmt.Printf("获取用户资料: %s\n", user.Name)
	return user, nil
}

func (s *UserService) RegisterUser(name, email string) (*User, error) {
	// 业务逻辑：用户注册
	user := &User{
		ID:    100, // 模拟生成ID
		Name:  name,
		Email: email,
	}

	err := s.repo.CreateUser(user)
	if err != nil {
		return nil, err
	}

	fmt.Printf("用户注册成功: %s\n", user.Name)
	return user, nil
}

// UserController 用户控制器层
type UserController struct {
	service *UserService // 依赖用户服务
}

// NewUserController 用户控制器构造函数
func NewUserController(service *UserService) *UserController {
	return &UserController{service: service}
}

func (c *UserController) HandleGetUser(id int) {
	// 模拟HTTP请求处理
	fmt.Printf("处理GET /users/%d请求\n", id)

	user, err := c.service.GetUserProfile(id)
	if err != nil {
		fmt.Printf("错误: %v\n", err)
		return
	}

	// 模拟返回JSON响应
	fmt.Printf("返回用户信息: %+v\n", user)
}

func (c *UserController) HandleCreateUser(name, email string) {
	// 模拟HTTP POST请求处理
	fmt.Printf("处理POST /users请求: name=%s, email=%s\n", name, email)

	user, err := c.service.RegisterUser(name, email)
	if err != nil {
		fmt.Printf("错误: %v\n", err)
		return
	}

	// 模拟返回创建成功响应
	fmt.Printf("用户创建成功: %+v\n", user)
}

// Logger 日志记录器
type Logger struct {
	Level string
}

// NewLogger 日志记录器构造函数
func NewLogger() *Logger {
	return &Logger{Level: "INFO"}
}

func (l *Logger) Info(msg string) {
	fmt.Printf("[%s] %s\n", l.Level, msg)
}

// App 应用程序主结构
type App struct {
	controller *UserController
	logger     *Logger
}

// NewApp 应用程序构造函数
func NewApp(controller *UserController, logger *Logger) *App {
	return &App{
		controller: controller,
		logger:     logger,
	}
}

func (app *App) Start() {
	app.logger.Info("应用程序启动")

	// 模拟处理一些HTTP请求
	app.controller.HandleGetUser(1)
	app.controller.HandleCreateUser("张三", "zhangsan@example.com")

	app.logger.Info("应用程序运行中...")
}

func main() {
	// 创建DI容器
	container := di.New()

	// 注册所有依赖的构造函数
	// 注意：注册顺序不重要，DI容器会自动解析依赖关系
	if err := container.Provide(NewMySQLDatabase); err != nil {
		log.Fatal(err)
	}
	if err := container.Provide(NewUserRepository); err != nil {
		log.Fatal(err)
	}
	if err := container.Provide(NewUserService); err != nil {
		log.Fatal(err)
	}
	if err := container.Provide(NewUserController); err != nil {
		log.Fatal(err)
	}
	if err := container.Provide(NewLogger); err != nil {
		log.Fatal(err)
	}
	if err := container.Provide(NewApp); err != nil {
		log.Fatal(err)
	}

	// 启动应用程序，DI容器会自动创建整个依赖链
	err := container.Invoke(func(app *App) {
		// 此时app及其所有依赖都已经完全初始化
		// 依赖链: App -> (UserController, Logger) -> UserService -> UserRepository -> Database
		app.Start()
	})

	if err != nil {
		log.Fatal(err)
	}
}
