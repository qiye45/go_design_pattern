package proxy

import "fmt"

//go:generate go run generate_proxy.go

// 主题接口
type Subject interface {
	Request() string
}

// 真实主题
type RealSubject struct {
	data string
}

func (r *RealSubject) Request() string {
	return fmt.Sprintf("RealSubject: %s", r.data)
}

// 静态代理
type StaticProxy struct {
	realSubject *RealSubject
}

func NewStaticProxy(data string) *StaticProxy {
	return &StaticProxy{
		realSubject: &RealSubject{data: data},
	}
}

func (p *StaticProxy) Request() string {
	fmt.Println("Proxy: Pre-processing")
	result := p.realSubject.Request()
	fmt.Println("Proxy: Post-processing")
	return result
}

// 缓存代理
type CacheProxy struct {
	realSubject *RealSubject
	cache       string
	cached      bool
}

func NewCacheProxy(data string) *CacheProxy {
	return &CacheProxy{
		realSubject: &RealSubject{data: data},
	}
}

func (p *CacheProxy) Request() string {
	if !p.cached {
		p.cache = p.realSubject.Request()
		p.cached = true
	}
	return p.cache
}
