package main

import "fmt"

// Parser 产品接口
type Parser interface{ Parse(data string) }

// JsonParser 具体产品
type JsonParser struct{}

func (JsonParser) Parse(data string) { fmt.Println("json解析", data) }

type YamlParser struct{}

func (YamlParser) Parse(data string) { fmt.Println("yaml解析", data) }

// NewParser 简单工厂
func NewParser(t string) Parser {
	switch t {
	case "json":
		return JsonParser{}
	case "yaml":
		return YamlParser{}
	}
	return nil
}

func main() {
	p := NewParser("json")
	p.Parse(`{"a":1}`)
}
