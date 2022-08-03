package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"
)

type HarvestHTTP struct {
	Config HarvestConfig
	client http.Client
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

	url := fmt.Sprintf("%s%s", h.Config.HarvestApiUrl, path)
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatalf("Failed to create requqest: %v", err)
	}

	accId := strconv.Itoa(h.Config.HarvestAccountId)
	req.Header.Set("Harvest-Account-ID", accId)
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", h.Config.HarvestAuthToken))

	return h.client.Do(req)
}
