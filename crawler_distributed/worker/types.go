package worker

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/xiangbaoyan/study_go_test/crawler/engine"
	"github.com/xiangbaoyan/study_go_test/crawler/zhenai/parser"
	"github.com/xiangbaoyan/study_go_test/crawler_distributed/config"
	"log"
)

type SerializedParser struct {
	Name string
	Args interface{}
}
type Request struct {
	Url    string
	Parser SerializedParser
}

type ParseResult struct {
	Items    []engine.Item
	Requests []Request
}

func SerializeRequest(r engine.Request) Request {
	name, args := r.Parser.Serialize()
	return Request{
		Url: r.Url,
		Parser: SerializedParser{
			Name: name,
			Args: args,
		},
	}
}

func SerializeResult(
	r engine.ParseResult) ParseResult {
	result := ParseResult{
		Items: r.Items,
	}
	for _, req := range r.Requests {
		result.Requests = append(result.Requests, SerializeRequest(req))
	}
	return result
}

func DeserializeResult(r ParseResult) engine.ParseResult {
	result := engine.ParseResult{
		Items: r.Items,
	}
	for _, req := range r.Requests {
		engineReq, err := DeserializeRequest(req)
		if err != nil {
			log.Printf("error deserializing request: %v", engineReq)
			continue
		}
		result.Requests = append(result.Requests, engineReq)
	}
	return result
}

func DeserializeRequest(
	r Request) (engine.Request, error) {
	res, err := deserializeParser(r.Parser)
	if err != nil {
		return engine.Request{}, err
	}
	return engine.Request{
		Url:    r.Url,
		Parser: res,
	}, nil
}

func deserializeParser(
	p SerializedParser) (engine.Parser, error) {
	switch p.Name {
	case config.ParseCityList:
		return engine.NewFuncParser(
			parser.ParseCityList,
			config.ParseCityList), nil
	case config.ParseCity:
		return engine.NewFuncParser(
			parser.ParseCity,
			config.ParseCity), nil
	case config.NilParser:
		return engine.NilParser{}, nil
	case config.ParseProfile:
		if username, ok := p.Args.(string); ok {
			return parser.NewProfileParser(username), nil
		} else {
			return nil, fmt.Errorf("invalid "+
				"arg: %v", p.Args)
		}
	default:
		return nil, errors.New("unKnown parser name ")
	}

}
