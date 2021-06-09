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
	Id       string
	Url      string
	Payload  interface{}
	SaveFunc func(item Item, saver chan interface{})
}

func NilParse([]byte) ParseResult {
	return ParseResult{}
}
