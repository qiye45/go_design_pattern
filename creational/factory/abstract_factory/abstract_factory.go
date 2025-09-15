package abstract_factory

type IRuleConfigParser interface {
	Parse(data []byte)
}

type jsonRuleConfigParser struct{}

func (j jsonRuleConfigParser) Parse(data []byte) {
	panic("implement me")
}

type ISystemConfigParser interface {
	ParseSystem(data []byte)
}

type jsonSystemConfigParser struct{}

func (j jsonSystemConfigParser) ParseSystem(data []byte) {
	panic("implement me")
}

// IConfigParserFactory 工厂方法接口
type IConfigParserFactory interface {
	CreateRuleParser() IRuleConfigParser
	CreateSystemParser() ISystemConfigParser
}

type jsonConfigParserFactory struct{}

func (j jsonConfigParserFactory) CreateRuleParser() IRuleConfigParser {
	return jsonRuleConfigParser{}
}

func (j jsonConfigParserFactory) CreateSystemParser() ISystemConfigParser {
	return jsonSystemConfigParser{}
}

type yamlConfigParserFactory struct{}

func (y yamlConfigParserFactory) CreateRuleParser() IRuleConfigParser {
	return yamlRuleConfigParser{}
}

func (y yamlConfigParserFactory) CreateSystemParser() ISystemConfigParser {
	return yamlSystemConfigParser{}
}

type yamlRuleConfigParser struct{}

func (y yamlRuleConfigParser) Parse(data []byte) {
	panic("implement me")
}

type yamlSystemConfigParser struct{}

func (y yamlSystemConfigParser) ParseSystem(data []byte) {
	panic("implement me")
}

// NewIConfigParserFactory ...
// 抽象工厂基于工厂方法，用一个简单工厂封装工厂方法
// 只不过一个工厂方法可创建多个相关的类
func NewIConfigParserFactory(t string) IConfigParserFactory {
	switch t {
	case "json":
		return jsonConfigParserFactory{}
	case "yaml":
		return yamlConfigParserFactory{}
	}
	return nil
}
