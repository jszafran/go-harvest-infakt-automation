package main

import (
	"errors"
	"fmt"
	"time"
)

type MonthRange struct {
	Start string
	End   string
}

// MonthsLastDay returns a last day (integer) of given month & year.
func MonthsLastDay(month int, year int) (int, error) {
	var lastDay int

	if month < 1 || month > 12 {
		return lastDay, errors.New(fmt.Sprintf("invalid month: %d", month))
	}

	if year < 2022 {
		return lastDay, errors.New(fmt.Sprintf("invalid year: %d", year))
	}

	for day := 28; day <= 31; day++ {
		dt := time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
		if _, m, _ := dt.Add(time.Hour * 24).Date(); !(m == time.Month(month)) {
			lastDay = dt.Day()
			break
		}
	}
	return lastDay, nil
}

func AsMonthRange(month int, year int) (MonthRange, error) {
	ld, err := MonthsLastDay(month, year)
	if err != nil {
		return MonthRange{}, err
	}
	start := IsoDate(time.Date(year, time.Month(month), 1, 0, 0, 0, 0, time.UTC))
	end := IsoDate(time.Date(year, time.Month(month), ld, 0, 0, 0, 0, time.UTC))
	return MonthRange{
		Start: start,
		End:   end,
	}, nil
}

func IsoDate(d time.Time) string {
	return d.Format(time.RFC3339)[:10]
}
