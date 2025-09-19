package main

import "fmt"

// Iterator 迭代器接口
type Iterator interface {
	HasNext() bool
	Next() any
}

// IntIterator 具体迭代器
type IntIterator struct {
	data []int
	idx  int
}

func (it *IntIterator) HasNext() bool { return it.idx < len(it.data) }
func (it *IntIterator) Next() any {
	val := it.data[it.idx]
	it.idx++
	return val
}

// IntCollection 集合
type IntCollection struct{ data []int }

func (c *IntCollection) Iterator() Iterator {
	return &IntIterator{data: c.data}
}

func main() {
	coll := &IntCollection{data: []int{10, 20, 30}}
	it := coll.Iterator()

	for it.HasNext() {
		fmt.Println(it.Next())
	}
}
