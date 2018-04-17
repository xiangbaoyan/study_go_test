package engine

import "log"

type Request struct {
	Url       string
	ParseFunc func([]byte) ParseResult
}

type ParseResult struct {
	Requests []Request
	Items    []interface{}
}

func NilParser([]byte) ParseResult {
	log.Println("得到新数据，继续处理")
	return ParseResult{}
}
