package main

import (
	"errors"
	"testing"
)

func TestMonthsLastDay(t *testing.T) {
	type TestCase struct {
		year  int
		month int
		want  int
		err   error
	}

	testCases := []TestCase{
		{2022, 0, 0, errors.New("invalid month: 0")},
		{2022, 13, 0, errors.New("invalid month: 13")},
		{2021, 1, 0, errors.New("invalid year: 2021")},
		{2022, 7, 31, nil},
		{2022, 8, 31, nil},
		{2022, 9, 30, nil},
		{2022, 12, 31, nil},
		{2022, 2, 28, nil},
		// test leap year
		{2024, 2, 29, nil},
	}

	for _, tc := range testCases {
		got, err := MonthsLastDay(tc.month, tc.year)
		if got != tc.want {
			t.Fatalf("Expected %v but got %v", tc.want, got)
		}

		if tc.err != nil {
			if err.Error() != tc.err.Error() {
				t.Fatalf("Expected err %v but got err %v", tc.err, err)
			}
		}
	}
}
