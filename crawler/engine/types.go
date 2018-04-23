package engine

import "log"

type ParseFunc func([]byte) ParseResult
type Request struct {
	Url       string
	ParseFunc func([]byte) ParseResult
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

func NilParser([]byte) ParseResult {
	log.Println("得到新数据，继续处理")
	return ParseResult{}
}
