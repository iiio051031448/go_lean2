package parser

import (
	"io/ioutil"
	"testing"
)

func TestParseCityList(t *testing.T) {
	//contents, err := fetcher.Fetch("https://www.zhenai.com/zhenghun")
	contents, err := ioutil.ReadFile("citylist_test_data.html")
	if err != nil {
		return
	}

	pResult := ParseCityList(contents)

	const resultSize = 470
	expectedUrls := []string{
		"http://www.zhenai.com/zhenghun/aba",
		"http://www.zhenai.com/zhenghun/akesu",
		"http://www.zhenai.com/zhenghun/alashanmeng",
	}

	for i, url := range expectedUrls {
		if pResult.Requests[i].Url != url {
			t.Errorf("[%d] Url expect :%s, but :%s", i, expectedUrls[i], pResult.Requests[i].Url)
		}
	}

	if len(pResult.Requests) != 470 {
		t.Errorf("Requests expected: %d, but :%d", resultSize, len(pResult.Requests))
	}
}
