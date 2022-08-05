package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"time"
)

type Args struct {
	Month int
	Year  int
}

func GetArgs() Args {
	var (
		month,
		year int
	)

	currTime := time.Now()
	flag.IntVar(&month, "month", int(currTime.Month()), "Month to generate the data for.")
	flag.IntVar(&year, "year", currTime.Year(), "Year to generate the data for")
	flag.Parse()

	return Args{
		Month: month,
		Year:  year,
	}
}

func main() {
	cfgPath := "config.json"
	hvst, err := NewHarvestClient(cfgPath)
	if err != nil {
		log.Fatal(err)
	}
	args := GetArgs()

	fmt.Printf("Will fetch data for month: %d year: %d\n", args.Month, args.Year)
	entries, err := hvst.GetTimeEntries(args.Month, args.Year)

	if err != nil {
		log.Fatal(err)
	}

	ms := MakeMonthlySummary(entries)

	res := ProceedWithInfaktInvoice(ms)

	if !res {
		fmt.Println("Quitting")
		os.Exit(0)
	}

	infakt, err := NewInfaktClient(cfgPath)
	if err != nil {
		log.Fatal(err)
	}

	err = infakt.CreateDraftInvoice(args.Month, args.Year, ms)
	if err != nil {
		log.Fatal(err)
	}
}
