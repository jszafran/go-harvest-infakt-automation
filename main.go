package main

import (
	"fmt"
	"log"
)

func main() {
	hvst, err := NewHarvestClient("config.json")
	if err != nil {
		log.Fatal(err)
	}
	entr, err := hvst.GetTimeEntries(7, 2022)

	if err != nil {
		log.Fatal(err)
	}

	for _, te := range entr {
		fmt.Printf("%+v", te)
	}
}
