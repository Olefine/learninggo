package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Movie struct {
	Title  string
	Year   int  `json:"released"`
	Color  bool `json:"color,omitempty"`
	Actors []string
}

func main() {
	movies := []Movie{
		{Title: "Casablanca", Year: 1942, Color: false, Actors: []string{"Egor", "Ramsy"}},
		{Title: "Birth pf Egor", Year: 1993, Color: true, Actors: []string{"Unnamed hero"}},
	}

	data, err := json.Marshal(movies)

	if err != nil {
		log.Fatal("%v", err)
	}

	fmt.Println(string(data))
}
