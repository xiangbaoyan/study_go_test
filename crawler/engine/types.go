package engine

type Parser interface {
	Parse(contents []byte, url string) ParseResult
	Serialize() (name string, args interface{})
}

type ParseFunc func(contents []byte, url string) ParseResult

//放一个parser 进去就可以调用方法了
//将Parser 转成一个接口 ，这个接口有两个实现
type Request struct {
	Url    string
	Parser Parser
}

type ParseResult struct {
	Requests []Request
	Items    []Item
}

type Item struct {
	Url     string
	Id      string
	Type    string
	Payload interface{}
}

type NilParser struct {
}

func (NilParser) Parse(_ []byte, _ string) ParseResult {
	return ParseResult{}
}

func (NilParser) Serialize() (name string, args interface{}) {
	return "NilParser", args
}

//func NilParser([]byte) ParseResult {
//	log.Println("得到新数据，继续处理")
//	return ParseResult{}
//}
//这里奇怪的是下面New的时候，方法就会自动执行（不一定，可能没有执行，在errhandling那一课可以看看）
type FuncParser struct {
	parser ParseFunc
	name   string
}

func (f *FuncParser) Parse(contents []byte, url string) ParseResult {
	return f.parser(contents, url)
}

func (f *FuncParser) Serialize() (name string, args interface{}) {
	return f.name, nil
}

//看到名字就会调用，
func NewFuncParser(p ParseFunc, name string) *FuncParser {
	//专门为rpc 序列化准备的
	return &FuncParser{
		parser: p,
		name:   name,
	}

}
