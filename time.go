package main

import (
	"errors"
	"fmt"
	"time"
)

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
