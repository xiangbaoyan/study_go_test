package worker

import "github.com/xiangbaoyan/study_go_test/crawler/engine"

//开始注册了
type CrawService struct{}

func (CrawService) Process(
	req Request,
	result *ParseResult) error {

	engineReq, err := DeserializeRequest(req)
	if err != nil {
		return err
	}
	parseResult, err := engine.Worker(engineReq)
	if err != nil {
		return err
	}
	*result = SerializeResult(parseResult)
	return nil

}
