package main

import (
	"log"
	"net/http"
	"os"
	"stu/http/crawler/frontend/controller"
)

func main() {
	pwd, _ := os.Getwd()
	log.Printf("PWD:%s\n", pwd)

	http.Handle("/", http.FileServer(
		http.Dir("http/crawler/frontend/view")))
	http.Handle("/search",
		controller.CreateSearchResultHandler(
			"http/crawler/frontend/view/template.html"))
	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		panic(err)
	}
}
