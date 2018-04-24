package controller

import (
	"context"
	"github.com/xiangbaoyan/study_go_test/crawler/engine"
	"github.com/xiangbaoyan/study_go_test/crawler/frontend/modal"
	"github.com/xiangbaoyan/study_go_test/crawler/frontend/view"
	"gopkg.in/olivere/elastic.v5"
	"net/http"
	"reflect"
	"strconv"
	"strings"
)

type SearchResultHandler struct {
	view   view.SearchResultView
	client *elastic.Client
}

func CreateSearchResultHandler(template string) SearchResultHandler {

	client, err := elastic.NewClient(elastic.SetSniff(false))
	if err != nil {
		panic(err)
	}

	return SearchResultHandler{
		view:   view.CreateSearchResultView(template),
		client: client,
	}
}

func (h SearchResultHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	q := strings.TrimSpace(req.FormValue("q"))
	from, err := strconv.Atoi(req.FormValue("from"))
	if err != nil {
		from = 0
	}
	var page modal.SearchResult
	page, err = h.getSearchResult(q, from)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	err = h.view.Render(w, page)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

}

func (h SearchResultHandler) getSearchResult(q string, from int) (modal.SearchResult, error) {
	var result modal.SearchResult
	resp, err := h.client.Search("dating_profile").
		Query(elastic.NewQueryStringQuery(q)).
		From(from).
		Do(context.Background())

	if err != nil {
		return result, err
	}

	result.Hits = resp.TotalHits()
	result.Start = from
	//for _,v := range resp.Each(reflect.TypeOf(engine.Item{})){
	//	item := v.(engine.Item)
	//}
	result.Items = resp.Each(
		reflect.TypeOf(engine.Item{}))
	return result, nil

}
