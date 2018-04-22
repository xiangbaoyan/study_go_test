package main

import (
	"fmt"
	"github.com/xiangbaoyan/study_go_test/crawler/fetcher"
)

func main() {
	//http://album.zhenai.com/u/107157060
	//http://album.zhenai.com/u/108816494
	//http://album.zhenai.com/u/1753109395
	bytes, err := fetcher.Fetch("http://album.zhenai.com/u/108340848")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(bytes))
}
