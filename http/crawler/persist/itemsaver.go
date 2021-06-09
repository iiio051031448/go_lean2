package persist

import (
	"context"
	"encoding/json"
	"github.com/olivere/elastic/v7"
	"log"
)

func ItemSaver() chan interface{} {
	out := make(chan interface{})

	go func() {
		for {
			item := <-out
			log.Printf("ItemSaver Got item:%+v", item)

			data_json, err := json.Marshal(item)
			if err != nil {
				log.Printf("Item to JSON string failed.\n")
				continue
			}
			save(string(data_json))
		}
	}()

	return out
}

func save(data string) {
	client, err := elastic.NewClient(
		elastic.SetURL("http://192.168.136.128:9200"),
		elastic.SetSniff(false))
	if err != nil {
		log.Printf("ElasticSearch Client create failed.\n")
		return
	}

	indexResp, err := client.Index().Index("data_profile").
		Id("").
		BodyString(data).Do(context.Background())
	if err != nil {
		log.Printf("save data to ElasticSearch Failed.\n")
		return
	}

	log.Printf("indexResp: %+v\n", indexResp)
}
