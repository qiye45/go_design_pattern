package main

import "fmt"

// 中间件函数类型：接收请求参数和一个“继续执行”的函数
type HandlerFunc func(req int, next func())

func main() {
	// 定义两个中间件
	m1 := func(req int, next func()) {
		fmt.Println("进入 m1")
		next() // 调用下一个中间件
		fmt.Println("离开 m1")
	}

	m2 := func(req int, next func()) {
		fmt.Println("进入 m2")
		next()
		// 这里没有“离开 m2”的打印，模拟只在前面处理
	}

	// 执行中间件链
	runChain := func(req int, middlewares []HandlerFunc) {
		// 定义一个递归函数，用来依次执行中间件
		var call func(index int)
		call = func(index int) {
			if index < len(middlewares) {
				// 执行第 index 个中间件，并传入一个 next 函数
				middlewares[index](req, func() {
					call(index + 1) // 调用下一个
				})
			}
		}
		call(0) // 从第一个中间件开始
	}

	// 模拟请求
	runChain(123, []HandlerFunc{m1, m2})
}
