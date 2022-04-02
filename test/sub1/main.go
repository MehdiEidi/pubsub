package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/MehdiEidi/pubsub/service"
)

type subscriber struct {
	Topics []string `json:"topics"`
	Addr   string   `json:"addr"`
}

func main() {
	ss := subscriber{Topics: []string{"football", "volleyball", "waterpolo"}, Addr: "http://localhost:8081/"}

	j, err := json.Marshal(ss)
	if err != nil {
		log.Fatal(err)
	}

	go service.Listen(":8081")

	_, err = http.Post("http://localhost:8080/subscribe", "application/json", bytes.NewBuffer(j))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Scanln()
}
