package main

import (
	elasticsearch8 "github.com/elastic/go-elasticsearch/v8"
	"log"
	"os"
	"fmt"
)


func main() {
	elasticPwd, ok := os.LookupEnv("ELASTIC_PASSWORD")
	if !ok {
		log.Fatalln("Missing elasticPwd var env")
	}

	cfg := elasticsearch8.Config{
		Addresses: []string{"https://localhost:9200"},
		Username: "elastic",
		Password: elasticPwd,
	}
	clientEs8, err := elasticsearch8.NewClient(cfg)
	if err != nil {
		log.Fatalf("Error creating the client: %s", err)
	}
}
