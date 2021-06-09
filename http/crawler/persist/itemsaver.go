package persist

import "log"

func ItemSaver() chan interface{} {
	out := make(chan interface{})

	go func() {
		for {
			item := <-out
			log.Printf("ItemSaver Got item:%+v", item)
		}
	}()

	return out
}
