package main

import (
	"errors"
	"log"
	"reflect"
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

func TestAsMonthRange(t *testing.T) {
	type TestCase struct {
		month int
		year  int
		want  MonthRange
		err   error
	}

	testCases := []TestCase{
		{1, 2022, MonthRange{Start: "2022-01-01", End: "2022-01-31"}, nil},
		{2, 2022, MonthRange{Start: "2022-02-01", End: "2022-02-28"}, nil},
		{3, 2022, MonthRange{Start: "2022-03-01", End: "2022-03-31"}, nil},
		{4, 2022, MonthRange{Start: "2022-04-01", End: "2022-04-30"}, nil},
		{2, 2024, MonthRange{Start: "2024-02-01", End: "2024-02-29"}, nil},
	}

	for _, tc := range testCases {
		got, err := AsMonthRange(tc.month, tc.year)
		if !reflect.DeepEqual(tc.want, got) {
			t.Fatalf("Expected %v but got %v", tc.want, got)
		}

		if tc.err != nil {
			if err == nil {
				log.Fatalf("Expected %v error but got none", tc.err)
			}
			if tc.err.Error() != err.Error() {
				t.Fatalf("Expected error %v but got %v", tc.err, err)
			}
		}
	}
}
