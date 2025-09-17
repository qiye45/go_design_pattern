package main

import "fmt"

// 产品接口
type Parser interface{ Parse(data string) }

// 工厂接口
type Factory interface{ Create() Parser }

// json实现
type JsonParser struct{}

func (JsonParser) Parse(data string) { fmt.Println("json解析", data) }

type JsonFactory struct{}

func (JsonFactory) Create() Parser { return JsonParser{} }

// yaml实现
type YamlParser struct{}

func (YamlParser) Parse(data string) { fmt.Println("yaml解析", data) }

type YamlFactory struct{}

func (YamlFactory) Create() Parser { return YamlParser{} }

func main() {
	var f Factory = JsonFactory{}
	p := f.Create()
	p.Parse(`{"a":1}`)
}
