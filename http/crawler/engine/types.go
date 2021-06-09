package engine

type Request struct {
	Url       string
	ParseFunc func([]byte) ParseResult
}

type ParseResult struct {
	Requests       []Request
	Items          []interface{}
	ItemHandleFunc func(item interface{})
}

func NilParse([]byte) ParseResult {
	return ParseResult{}
}
