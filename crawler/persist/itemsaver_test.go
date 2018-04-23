package persist

import (
	"context"
	"encoding/json"
	"github.com/xiangbaoyan/study_go_test/crawler/engine"
	"github.com/xiangbaoyan/study_go_test/crawler/modal"
	"gopkg.in/olivere/elastic.v5"
	"testing"
)

func TestSaver(t *testing.T) {
	expected := engine.Item{
		Url:  "http://album.zhenai.com/u/107157060",
		Type: "zhenai",
		Id:   "107157060",
		Payload: modal.Profile{
			Name:       "小红",
			Gender:     "女",
			Age:        0,
			Height:     160,
			Weight:     0,
			Income:     "3000元以下",
			Marriage:   "离异",
			Education:  "高中及以下",
			Occupation: "--",
			Hokou:      "--",
			Xinzuo:     "",
			House:      "--",
			Car:        "未购车",
		},
	}
	client, err := elastic.NewClient(elastic.SetSniff(false))
	if err != nil {
		panic(err)
	}
	const index = "dating_test"
	err = save(client, index, expected)

	if err != nil {
		panic(err)
	}
	client, e := elastic.NewClient(elastic.SetSniff(false))
	if e != nil {
		panic(e)
	}
	result, err := client.Get().Index(index).
		Type(expected.Type).Id(expected.Id).Do(context.Background())
	if err != nil {
		panic(err)
	}
	//t.Logf("%+v", result)
	var actual engine.Item
	err = json.Unmarshal([]byte(*result.Source), &actual)
	if err != nil {
		t.Logf("%+v", actual)
		panic(err)
	}
	actualProfile, _ := modal.FromJsonObj(actual.Payload)
	actual.Payload = actualProfile
	if actual != expected {
		t.Errorf("got %+v,expected %+v", actual, expected)
	}

}
