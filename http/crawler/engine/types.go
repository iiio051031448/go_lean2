package engine

type Request struct {
	Url       string
	ParseFunc func([]byte) ParseResult
}

type ParseResult struct {
	Requests []Request
	Items    []Item
}

type Item struct {
	Id         string
	Url        string
	Payload    interface{}
	HandleFunc func(item Item)
}

func NilParse([]byte) ParseResult {
	return ParseResult{}
}
