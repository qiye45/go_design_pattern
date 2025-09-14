package singleton

import "sync"

// EagerSingleton 饿汉式单例
type EagerSingleton struct {
	data string
}

var eagerInstance = &EagerSingleton{data: "eager"}

func GetEagerInstance() *EagerSingleton {
	return eagerInstance
}

// LazySingleton 懒汉式单例
type LazySingleton struct {
	data string
}

var (
	lazyInstance *LazySingleton
	once         sync.Once
)

func GetLazyInstance() *LazySingleton {
	once.Do(func() {
		lazyInstance = &LazySingleton{data: "lazy"}
	})
	return lazyInstance
}

func (s *EagerSingleton) GetData() string {
	return s.data
}

func (s *LazySingleton) GetData() string {
	return s.data
}
