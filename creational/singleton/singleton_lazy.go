package singleton

import "sync"

var (
	lazySingleton *ConfigSingleton
	once          = &sync.Once{}
)

// GetLazyInstance 懒汉式
func GetLazyInstance() *ConfigSingleton {
	// once 内的方法只会执行一次，所以不需要再次判断
	once.Do(func() {
		lazySingleton = &ConfigSingleton{
			appName:    "MyApp",
			appVersion: "1.0.0",
		}
	})

	return lazySingleton
}
