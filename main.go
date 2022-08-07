package main

import (
	"log"
	"os"
	"strconv"
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
	argsLen := len(os.Args)
	if !(argsLen == 1 || argsLen == 3) {
		msg := `Incorrect number of args.
Either provide no args (current month & year will be used) or 2 args (month & year).
Example: geninvoice 7 2022
`
		log.Fatal(msg)
	}

	if argsLen == 1 {
		y, m, _ := time.Now().Date()
		return Args{
			Month: int(m),
			Year:  y,
		}
	}

	month, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	year, err = strconv.Atoi(os.Args[2])
	if err != nil {
		log.Fatal(err)
	}

	return Args{
		Month: month,
		Year:  year,
	}
}

func main() {
	config, err := AppConfigFromEnv()
	if err != nil {
		log.Fatal(err)
	}

	harvest := NewHarvestClient(config)
	if err != nil {
		log.Fatal(err)
	}

	args := GetArgs()

	log.Printf("Will fetch data for month: %d year: %d\n", args.Month, args.Year)
	entries, err := harvest.GetTimeEntries(args.Month, args.Year)

	if err != nil {
		log.Fatal(err)
	}

	if len(entries) == 0 {
		log.Println("No hours fetched from Harvest API. Exiting.")
		return
	}

	ms := MakeMonthlySummary(entries)

	res := ProceedWithInfaktInvoice(ms)

	if !res {
		log.Println("Exiting.")
		os.Exit(0)
	}

	infakt := NewInfaktClient(config)

	err = infakt.CreateDraftInvoice(args.Month, args.Year, ms)
	if err != nil {
		log.Fatal(err)
	}
}
