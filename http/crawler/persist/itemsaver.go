package persist

import (
	"context"
	"encoding/json"
	"log"
	"stu/http/crawler/engine"

	"github.com/olivere/elastic/v7"
)

func ItemSaver() chan interface{} {
	out := make(chan interface{})

	go func() {
		for {
			it := <-out
			item := it.(engine.Item)
			log.Printf("ItemSaver Got it:%+v", item)

			data_json, err := json.Marshal(item)
			if err != nil {
				log.Printf("Item to JSON string failed.\n")
				continue
			}
			save(item.Id, string(data_json))
		}
	}()

	return out
}

func save(id string, data string) {
	client, err := elastic.NewClient(
		elastic.SetURL("http://192.168.162.129:9200"),
		elastic.SetSniff(false))
	if err != nil {
		log.Printf("ElasticSearch Client create failed.\n")
		return
	}

	indexResp, err := client.Index().Index("data_profile").
		Id(id).BodyString(data).Do(context.Background())
	if err != nil {
		log.Printf("save data to ElasticSearch Failed.\n")
		return
	}

	log.Printf("indexResp: %+v\n", indexResp)
}
