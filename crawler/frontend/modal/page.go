package modal

type SearchResult struct {
	Hits     int64
	Start    int
	Query    string
	NextFrom int
	PrevFrom int
	Items    []interface{}
}
