package main

import "fmt"

// Container 定义一个容器
type Container struct {
	beans map[string]func() any // 保存构造函数
}

func NewContainer() *Container {
	return &Container{beans: make(map[string]func() any)}
}

// Register 注册：名字 + 构造函数
func (c *Container) Register(name string, creator func() any) {
	c.beans[name] = creator
}

// Get 获取：调用构造函数生成对象
func (c *Container) Get(name string) any {
	if creator, ok := c.beans[name]; ok {
		return creator()
	}
	return nil
}

// ================== 示例 ==================

type UserRepo struct{}

func (UserRepo) Find() { fmt.Println("查用户") }

type UserService struct {
	Repo UserRepo
}

func (u UserService) Do() { u.Repo.Find() }

func main() {
	c := NewContainer()

	// 注册依赖
	c.Register("repo", func() any { return UserRepo{} })
	c.Register("service", func() any {
		return UserService{Repo: c.Get("repo").(UserRepo)}
	})

	// 获取对象（自动注入依赖）
	s := c.Get("service").(UserService)
	s.Do()
}
