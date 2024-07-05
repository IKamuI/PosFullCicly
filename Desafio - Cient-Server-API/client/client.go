package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"
)

type Cotacao struct {
	Bid string `json:"bid"`
}

func main() {

	ctx, cancel := context.WithTimeout(context.Background(), 300*time.Millisecond)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "GET", "http://localhost:8080", nil)

	if err != nil {
		panic(err)
	}

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		panic(err)
	}

	defer res.Body.Close()

	var cotacao Cotacao

	err = json.NewDecoder(res.Body).Decode(&cotacao)

	if err != nil {
		panic(err)
	}

	file, err := os.OpenFile("client/cotacao.text", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	defer file.Close()
	if err != nil {
		log.Fatal(err)
	}

	if err != nil {
		panic(err)
	}

	file.Write([]byte("Dólar: " + cotacao.Bid + ",\n"))

}
