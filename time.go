package main

import (
	"errors"
	"fmt"
	"time"
)

func MonthsLastDay(month int, year int) (int, error) {
	var lastDay int
	if month < 1 && month > 12 {
		return lastDay, errors.New(fmt.Sprintf("invalid month: %d", month))
	}

	if year < 2022 {
		return lastDay, errors.New(fmt.Sprintf("invalid year: %d", year))
	}

	day := time.Date(year, time.Month(month), 28, 0, 0, 0, 0, time.UTC)

	for i := 28; i <= 31; i++ {
		nextDay := day.Add(time.Hour * 24)
		if d, m, _ := nextDay.Date(); !(m == time.Month(month)) {
			lastDay = d
		}
	}
	return lastDay, nil
}
