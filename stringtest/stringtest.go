package main

import (
	"log"
	"strings"
)

func main() {
	text := "http://s.opencalais.com/1/pred/BusinessRelationType"
	ts := strings.Split(text, "/")
	log.Printf("ts:%+v\n", ts)
	log.Printf("last:%+v\n", ts[len(ts)-1])
}
