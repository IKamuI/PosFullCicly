package main

import (
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"time"
)

type Cotacao []struct {
	Code       string `json:"code"`
	Codein     string `json:"codein"`
	Name       string `json:"name"`
	High       string `json:"high"`
	Low        string `json:"low"`
	VarBid     string `json:"varBid"`
	PctChange  string `json:"pctChange"`
	Bid        string `json:"bid"`
	Ask        string `json:"ask"`
	Timestamp  string `json:"timestamp"`
	CreateDate string `json:"create_date"`
}

func main() {

	http.HandleFunc("/", GetDollarHandler)
	http.ListenAndServe(":8080", nil)

}

func GetDollarHandler(w http.ResponseWriter, r *http.Request) {

	data, err := GetDollar()
	cotacao := *data
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(cotacao[0])
}

func GetDollar() (*Cotacao, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 200*time.Millisecond)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "GET", "https://economia.awesomeapi.com.br/json/USD-BRL", nil)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	resp, err := http.DefaultClient.Do(req)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)

	if err != nil {
		log.Println(err)
		return nil, err

	}

	var cotacao Cotacao

	json.Unmarshal(data, &cotacao)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &cotacao, nil

}
