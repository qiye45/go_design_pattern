package singleton

import "fmt"

// ConfigSingleton 饿汉式单例，持有配置信息
type ConfigSingleton struct {
	appName    string // 应用名称
	appVersion string // 应用版本
}

// instance 全局唯一单例实例
var instance *ConfigSingleton

// Init 在程序启动时初始化实例
func Init() {
	fmt.Println("init ConfigSingleton")
	instance = &ConfigSingleton{
		appName:    "MyApp",
		appVersion: "1.0.0",
	}
}

// GetInstance 获取全局唯一实例（饿汉式单例）
func GetInstance() *ConfigSingleton {
	return instance
}

// NewInstance 创建新的对象（非单例，用于测试）
func NewInstance() *ConfigSingleton {
	return &ConfigSingleton{
		appName:    "MyApp",
		appVersion: "1.0.0",
	}
}
