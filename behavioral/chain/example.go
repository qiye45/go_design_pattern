package chain

import "fmt"

// 处理器接口
type Handler interface {
	SetNext(Handler) Handler
	Handle(request string) string
}

// 基础处理器
type BaseHandler struct {
	next Handler
}

func (h *BaseHandler) SetNext(handler Handler) Handler {
	h.next = handler
	return handler
}

func (h *BaseHandler) Handle(request string) string {
	if h.next != nil {
		return h.next.Handle(request)
	}
	return ""
}

// 具体处理器A
type ConcreteHandlerA struct {
	BaseHandler
}

func (h *ConcreteHandlerA) Handle(request string) string {
	if request == "A" {
		return "Handled by A"
	}
	return h.BaseHandler.Handle(request)
}

// 具体处理器B
type ConcreteHandlerB struct {
	BaseHandler
}

func (h *ConcreteHandlerB) Handle(request string) string {
	if request == "B" {
		return "Handled by B"
	}
	return h.BaseHandler.Handle(request)
}

// 中间件示例
type Middleware func(string) string

type MiddlewareChain struct {
	middlewares []Middleware
}

func (mc *MiddlewareChain) Use(middleware Middleware) {
	mc.middlewares = append(mc.middlewares, middleware)
}

func (mc *MiddlewareChain) Execute(request string) string {
	result := request
	for _, middleware := range mc.middlewares {
		result = middleware(result)
	}
	return result
}

func LoggingMiddleware(request string) string {
	fmt.Printf("Logging: %s\n", request)
	return request
}

func AuthMiddleware(request string) string {
	return fmt.Sprintf("Auth(%s)", request)
}
