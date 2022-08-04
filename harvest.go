package main

import (
	"encoding/json"
	"fmt"
	"github.com/shopspring/decimal"
	"io"
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

func NewHarvestClient(configPath string) (HarvestHTTP, error) {
	var hv HarvestHTTP
	hvConf, err := HarvestConfigFromJSON(configPath)
	if err != nil {
		return hv, err
	}

	httpClient := http.Client{Timeout: time.Second * 5}
	return HarvestHTTP{
		Config: hvConf,
		client: httpClient,
	}, nil
}

func (h HarvestHTTP) getRequest(path string) (*http.Response, error) {
	if string(path[0]) == "/" {
		path = path[1:]
	}

	url := fmt.Sprintf("%s/%s", h.Config.HarvestApiUrl, path)
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatalf("Failed to create requqest: %v", err)
	}

	req.Header.Set("Harvest-Account-ID", h.Config.HarvestAccountId)
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", h.Config.HarvestAuthToken))

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

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return entries, err
	}

	err = json.Unmarshal(b, &apiResp)
	if err != nil {
		return entries, err
	}

	return apiResp.TimeEntries, nil
}
