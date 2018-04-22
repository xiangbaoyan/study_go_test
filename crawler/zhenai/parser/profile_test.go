package parser

import (
	"fmt"
	"github.com/xiangbaoyan/study_go_test/crawler/fetcher"
	"testing"
)

func TestParseProfile(t *testing.T) {

	bytes, e := fetcher.Fetch("http://album.zhenai.com/u/1476304799")

	if e != nil {
		panic(e)
	}
	result := ParseProfile(bytes, "小红")
	fmt.Println(result)
}
