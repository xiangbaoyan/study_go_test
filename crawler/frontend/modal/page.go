package modal

type SearchResult struct {
	Hits  int64
	Start int
	Query string
	Items []interface{}
}
