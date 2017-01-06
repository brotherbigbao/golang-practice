package main

import (
	"fmt"
	"encoding/json"
	"log"
)

type Movie struct {
	Title string
	Year int `json:"released"`
}

func main() {
	s := `[{"Title":"Casablanca","released":1942,"Actors":["Humphrey Bogart","Ingrid Bergman"]},{"Title":"Cool Hand Luke","released":1967,"color":true,"Actors":["Paul Newman"]},{"Title":"Bullitt","released":1968,"color":true,"Actors":["Steve McQueen","Jacqueline Bisset"]}]
`
	var m []Movie
	if err := json.Unmarshal([]byte(s), &m); err != nil {
		log.Fatalf("JSON unmarshaling failed: %s", err)
	}
	fmt.Println(m)
}
