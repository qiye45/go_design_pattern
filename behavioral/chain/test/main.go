package main

import (
	"fmt"
)

type Handler func(int, func(req int))

func main() {
	f1 := func(req int, next func(req int)) {
		fmt.Println("start f1", req)
		req += 1
		next(req)
		fmt.Println("stop f1", req)
	}
	f2 := func(req int, next func(req int)) {
		fmt.Println("start f2", req)
		req += 1
		next(req)
		fmt.Println("stop f2", req)
	}

	chain := func(req int, handlers []Handler) {
		var call func(index int, r int)
		call = func(index int, r int) {
			if index < len(handlers) {
				handlers[index](r, func(nextReq int) {
					call(index+1, nextReq)
				})
			}
		}
		call(0, req)
	}
	chain(0, []Handler{f1, f2})
}
