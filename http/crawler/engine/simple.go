package engine

import (
	"log"
	"stu/http/crawler/fetcher"
)

type SimpleEngine struct{}

func (e SimpleEngine) Run(seeds ...Request) {
	var requests []Request
	for _, r := range seeds {
		requests = append(requests, r)
	}

	for len(requests) > 0 {
		r := requests[0]
		requests = requests[1:]

		parseResult, err := worker(r, 0)
		if err != nil {
			return
		}

		requests = append(requests, parseResult.Requests...)

		for _, i := range parseResult.Items {
			log.Printf("Got item %v", i)
		}
	}
}

func worker(r Request, workId int) (ParseResult, error) {
	log.Printf("[#%d]Fetching and parse : %s", workId, r.Url)
	body, err := fetcher.Fetch(r.Url)
	if err != nil {
		log.Printf("Fetcher error: fetching url %s : %v", r.Url, err)
		return ParseResult{}, err
	}

	return r.ParseFunc(body), nil
}
