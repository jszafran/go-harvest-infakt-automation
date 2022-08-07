package main

import (
	"encoding/json"
	"fmt"
	"github.com/shopspring/decimal"
	"log"
	"net/http"
	"time"
)

type HarvestHTTP struct {
	Config HarvestConfig
	client http.Client
}

type Client struct {
	Name string `json:name`
}

type TimeEntry struct {
	Client       Client
	RoundedHours decimal.Decimal `json:"rounded_hours"`
}

type TimeEntriesApiResponse struct {
	TimeEntries []TimeEntry `json:"time_entries"`
}

type MonthlySummary map[string]decimal.Decimal

func (m MonthlySummary) Print() {
	if len(m) == 0 {
		log.Println("No hours logged for given period.")
		return
	}

	for k, v := range m {
		log.Printf("%s: %s hours\n", k, v)
	}
}

func NewHarvestClient(config AppConfig) HarvestHTTP {
	hvConf := config.Harvest
	httpClient := http.Client{Timeout: time.Second * 5}
	return HarvestHTTP{
		Config: hvConf,
		client: httpClient,
	}
}

func (h HarvestHTTP) getRequest(path string) (*http.Response, error) {
	if string(path[0]) == "/" {
		path = path[1:]
	}

	url := fmt.Sprintf("%s/%s", h.Config.ApiUrl, path)
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatalf("Failed to create requqest: %v", err)
	}

	req.Header.Set("Harvest-Account-ID", h.Config.AccountId)
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", h.Config.AuthToken))

	return h.client.Do(req)
}

func (h HarvestHTTP) GetTimeEntries(month int, year int) ([]TimeEntry, error) {
	var apiResp TimeEntriesApiResponse
	var entries []TimeEntry
	mr, err := AsMonthRange(month, year)

	if err != nil {
		return entries, nil
	}

	reqPath := fmt.Sprintf("time_entries?per_page=100&from=%s&to=%s", mr.Start, mr.End)
	resp, err := h.getRequest(reqPath)

	if err != nil {
		return entries, err
	}

	err = json.NewDecoder(resp.Body).Decode(&apiResp)
	if err != nil {
		return entries, err
	}

	return apiResp.TimeEntries, nil
}

func MakeMonthlySummary(te []TimeEntry) MonthlySummary {
	ms := make(map[string]decimal.Decimal)

	for _, entry := range te {
		ms[entry.Client.Name] = ms[entry.Client.Name].Add(entry.RoundedHours)
	}
	return ms
}
