package proxy

import "testing"

func TestStaticProxy(t *testing.T) {
	proxy := NewStaticProxy("test data")
	result := proxy.Request()

	expected := "RealSubject: test data"
	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}

func TestCacheProxy(t *testing.T) {
	proxy := NewCacheProxy("cached data")

	// 第一次调用
	result1 := proxy.Request()
	// 第二次调用（应该使用缓存）
	result2 := proxy.Request()

	if result1 != result2 {
		t.Error("Cache proxy should return same result")
	}
}
