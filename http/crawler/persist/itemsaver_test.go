package persist

import (
	"context"
	"reflect"
	"stu/http/crawler/engine"
	"testing"

	"github.com/olivere/elastic/v7"
)

func TestSave(t *testing.T) {
	client, err := elastic.NewClient(
		elastic.SetURL("http://192.168.136.128:9200"),
		elastic.SetSniff(false))
	if err != nil {
		panic(err)
	}

	resp, err := client.Get().
		Index("data_profile").
		Id("5203248443662889113").
		Do(context.Background())
	if err != nil {
		panic(err)
	}

	t.Logf("%s", resp.Source)

	searchResp, err := client.Search("data_profile").
		Query(elastic.NewQueryStringQuery("博士 Payload.Age:(<30)")).
		From(0).
		Do(context.Background())
	if err != nil {
		return
	}

	t.Logf("Hits %d", searchResp.TotalHits())
	items := searchResp.Each(reflect.TypeOf(engine.Item{}))

	for _, it := range items {
		t.Logf("%+v", it)
	}
}
