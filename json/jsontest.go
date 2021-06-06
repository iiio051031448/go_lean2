package main

import (
	json "encoding/json"
	"fmt"
)

type Order struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price,omitempty"`
}

func main() {
	o := Order{ID: "778899", Name: "liudehua", Price: 100}
	b, _ := json.Marshal(o)
	fmt.Printf("%s\n", b)

	os := `{"id":"778899","name":"liudehua","price":100}`
	var uo Order
	err := json.Unmarshal([]byte(os), &uo)
	if err != nil {
		return
	}
	fmt.Printf("%+v\n", uo)
}
