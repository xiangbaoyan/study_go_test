package parser

import (
	"github.com/xiangbaoyan/study_go_test/crawler/engine"
	"log"
	"regexp"
)

const cityListRe = `<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)"[^>]*>([^<]+)</a>`

//处理得出n多request
func ParseCityList(contents []byte) engine.ParseResult {
	compile := regexp.MustCompile(cityListRe)
	matches := compile.FindAllSubmatch(contents, -1)
	result := engine.ParseResult{}
	for _, m := range matches {
		//result.Items = append(result.Items, string(m[2]))
		log.Printf("Got One Url,%s", string(m[1]))

		result.Requests = append(
			result.Requests, engine.Request{
				Url:       string(m[1]),
				ParseFunc: ParseCity,
			})
	}

	return result

}
