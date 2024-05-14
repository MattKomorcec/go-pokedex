package util

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

const locationsUrl = "https://pokeapi.co/api/v2/location-area?offset=0&limit=20"

type Response struct {
	Count    int            `json:"count"`
	Next     string         `json:"next"`
	Previous *string        `json:"previous"`
	Results  []LocationArea `json:"results"`
}

type LocationArea struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

func LoadLocationAreas() {
	res, err := http.Get(locationsUrl)
	if err != nil {
		log.Fatal(err)
	}

	body, err := io.ReadAll(res.Body)
	res.Body.Close()

	if res.StatusCode > 299 {
		log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
	}
	if err != nil {
		log.Fatal(err)
	}

	parsedRes := Response{}
	err = json.Unmarshal(body, &parsedRes)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("previous: ", parsedRes.Previous)
	fmt.Println("next: ", parsedRes.Next)
	for _, v := range parsedRes.Results {
		fmt.Println(v.Name)
	}
}
