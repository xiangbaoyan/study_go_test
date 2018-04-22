package parser

import (
	"github.com/xiangbaoyan/study_go_test/crawler/engine"
	"regexp"
)

const cityRe = `<a href="(http://album.zhenai.com/u/[0-9]+)" target="_blank">([^<]+)</a>`

func ParseCity(contents []byte) engine.ParseResult {
	re := regexp.MustCompile(cityRe)
	matches := re.FindAllSubmatch(contents, -1)
	result := engine.ParseResult{}
	limit := 1
	for _, m := range matches {
		result.Items = append(result.Items, "User "+string(m[2]))
		result.Requests = append(result.Requests, engine.Request{
			Url: string(m[1]),
			ParseFunc: func(contents []byte) engine.ParseResult {
				return ParseProfile(contents, string(m[2]))
			},
		})
		limit--
		if limit == 0 {
			break
		}
	}

	return result

}
