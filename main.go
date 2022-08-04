package main

import (
	"log"
)

func main() {
	hvst, err := NewHarvestClient("config.json")
	if err != nil {
		log.Fatal(err)
	}
	entries, err := hvst.GetTimeEntries(8, 2022)

	if err != nil {
		log.Fatal(err)
	}

	ms := MakeMonthlySummary(entries)

	ms.Print()
}
