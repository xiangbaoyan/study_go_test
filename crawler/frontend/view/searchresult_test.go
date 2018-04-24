package view

import (
	"github.com/xiangbaoyan/study_go_test/crawler/engine"
	"github.com/xiangbaoyan/study_go_test/crawler/frontend/modal"
	common "github.com/xiangbaoyan/study_go_test/crawler/modal"
	"os"
	"testing"
)

func TestSearchResultView_Render(t *testing.T) {
	//temp := template.Must(template.ParseFiles("template.html"))
	view := CreateSearchResultView("template.html")
	out, err := os.Create("template.test.html")
	page := modal.SearchResult{}
	page.Hits = 123
	item := engine.Item{
		Url:  "http://album.zhenai.com/u/109879487",
		Type: "zhenai",
		Id:   "109879487",
		Payload: common.Profile{
			Name:       "sugar",
			Gender:     "女",
			Age:        25,
			Height:     170,
			Weight:     51,
			Income:     "8001-12000元",
			Marriage:   "未婚",
			Education:  "大学本科",
			Occupation: "总裁助理",
			Hokou:      "和家人同住",
			Xinzuo:     "金牛座",
			House:      "上海闵行区",
			Car:        "已购车",
		},
	}
	for i := 0; i < 10; i++ {
		page.Items = append(page.Items, item)
	}

	err = view.Render(out, page)
	//err = temp.Execute(out, page)
	if err != nil {
		panic(err)
	}

}
