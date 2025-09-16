package main

import (
	"errors"
	"testing"

	"github.com/qiye45/go_design_pattern/creational/factory/di"
)

// MockDatabase 模拟数据库，用于测试
type MockDatabase struct {
	users       map[int]*User
	shouldError bool
}

func NewMockDatabase() Database {
	return &MockDatabase{
		users:       make(map[int]*User),
		shouldError: false,
	}
}

func (db *MockDatabase) Connect() error {
	return nil
}

func (db *MockDatabase) FindUserByID(id int) (*User, error) {
	if db.shouldError {
		return nil, errors.New("数据库连接失败")
	}

	if user, exists := db.users[id]; exists {
		return user, nil
	}
	return nil, errors.New("用户不存在")
}

func (db *MockDatabase) SaveUser(user *User) error {
	if db.shouldError {
		return errors.New("保存失败")
	}

	db.users[user.ID] = user
	return nil
}

func (db *MockDatabase) SetError(shouldError bool) {
	db.shouldError = shouldError
}

// TestUserService 演示如何使用DI进行单元测试
func TestUserService_GetUserProfile(t *testing.T) {
	// 创建测试用的DI容器
	container := di.New()

	// 注册Mock数据库而不是真实数据库
	if err := container.Provide(NewMockDatabase); err != nil {
		t.Fatal(err)
	}
	if err := container.Provide(NewUserRepository); err != nil {
		t.Fatal(err)
	}
	if err := container.Provide(NewUserService); err != nil {
		t.Fatal(err)
	}

	// 测试成功场景
	err := container.Invoke(func(service *UserService, db Database) {
		// 准备测试数据
		mockDB := db.(*MockDatabase)
		testUser := &User{ID: 1, Name: "测试用户", Email: "test@example.com"}
		mockDB.SaveUser(testUser)

		// 执行测试
		user, err := service.GetUserProfile(1)
		if err != nil {
			t.Errorf("期望成功，但得到错误: %v", err)
		}

		if user.Name != "测试用户" {
			t.Errorf("期望用户名为'测试用户'，但得到: %s", user.Name)
		}
	})

	if err != nil {
		t.Fatal(err)
	}
}

func TestUserService_GetUserProfile_Error(t *testing.T) {
	// 测试错误场景
	container := di.New()

	if err := container.Provide(NewMockDatabase); err != nil {
		t.Fatal(err)
	}
	if err := container.Provide(NewUserRepository); err != nil {
		t.Fatal(err)
	}
	if err := container.Provide(NewUserService); err != nil {
		t.Fatal(err)
	}

	err := container.Invoke(func(service *UserService, db Database) {
		// 设置Mock数据库返回错误
		mockDB := db.(*MockDatabase)
		mockDB.SetError(true)

		// 执行测试
		_, err := service.GetUserProfile(1)
		if err == nil {
			t.Error("期望得到错误，但操作成功了")
		}
	})

	if err != nil {
		t.Fatal(err)
	}
}

func TestUserController_HandleGetUser(t *testing.T) {
	// 测试控制器层
	container := di.New()

	if err := container.Provide(NewMockDatabase); err != nil {
		t.Fatal(err)
	}
	if err := container.Provide(NewUserRepository); err != nil {
		t.Fatal(err)
	}
	if err := container.Provide(NewUserService); err != nil {
		t.Fatal(err)
	}
	if err := container.Provide(NewUserController); err != nil {
		t.Fatal(err)
	}

	err := container.Invoke(func(controller *UserController, db Database) {
		// 准备测试数据
		mockDB := db.(*MockDatabase)
		testUser := &User{ID: 2, Name: "控制器测试用户", Email: "controller@example.com"}
		mockDB.SaveUser(testUser)

		// 测试控制器方法（在实际项目中，这里会验证HTTP响应）
		controller.HandleGetUser(2)
		// 在真实测试中，这里会检查响应状态码、响应体等
	})

	if err != nil {
		t.Fatal(err)
	}
}

// 演示如何测试整个应用程序
func TestApp_Integration(t *testing.T) {
	// 集成测试：使用Mock数据库测试整个应用流程
	container := di.New()

	// 注册所有依赖，但使用Mock数据库
	if err := container.Provide(NewMockDatabase); err != nil {
		t.Fatal(err)
	}
	if err := container.Provide(NewUserRepository); err != nil {
		t.Fatal(err)
	}
	if err := container.Provide(NewUserService); err != nil {
		t.Fatal(err)
	}
	if err := container.Provide(NewUserController); err != nil {
		t.Fatal(err)
	}
	if err := container.Provide(NewLogger); err != nil {
		t.Fatal(err)
	}
	if err := container.Provide(NewApp); err != nil {
		t.Fatal(err)
	}

	err := container.Invoke(func(app *App) {
		// 运行应用程序的核心逻辑
		// 在真实的集成测试中，这里会验证整个业务流程
		app.Start()
	})

	if err != nil {
		t.Fatal(err)
	}
}
