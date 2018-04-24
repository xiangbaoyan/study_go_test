package view

import (
	"github.com/xiangbaoyan/study_go_test/crawler/frontend/modal"
	"html/template"
	"io"
)

type SearchResultView struct {
	template *template.Template
}

func CreateSearchResultView(fileName string) SearchResultView {
	return SearchResultView{
		template: template.Must(template.ParseFiles(fileName)),
	}
}

func (s *SearchResultView) Render(w io.Writer, data modal.SearchResult) error {
	return s.template.Execute(w, data)
}
