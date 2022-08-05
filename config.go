package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type AppConfig struct {
	Harvest HarvestConfig `json:"harvest"`
	Infakt  InfaktConfig  `json:"infakt"`
}

type HarvestConfig struct {
	ApiUrl    string `json:"harvestApiUrl"`
	AuthToken string `json:"harvestAuthToken"`
	AccountId string `json:"harvestAccountId"`
}

type InfaktConfig struct {
	ApiUrl            string `json:"infaktApiUrl"`
	ApiKey            string `json:"infaktApiKey"`
	HourlyRateInGrosz uint   `json:"infaktHourlyRateInGrosz"`
	ClientId          uint   `json:"infaktClientId"`
}

func appConfigFromJSON(jsonPath string) (AppConfig, error) {
	var ac AppConfig
	file, err := os.Open(jsonPath)
	if err != nil {
		return ac, err
	}

	b, err := ioutil.ReadAll(file)

	if err != nil {
		return ac, err
	}

	err = json.Unmarshal(b, &ac)

	if err != nil {
		return ac, err
	}

	return ac, nil
}

func HarvestConfigFromJSON(jsonPath string) (HarvestConfig, error) {
	var hvConf HarvestConfig
	ac, err := appConfigFromJSON(jsonPath)

	if err != nil {
		return hvConf, err
	}

	return HarvestConfig{
		ApiUrl:    ac.Harvest.ApiUrl,
		AuthToken: ac.Harvest.AuthToken,
		AccountId: ac.Harvest.AccountId,
	}, nil
}

func InfaktConfigFromJSON(jsonPath string) (InfaktConfig, error) {
	var ifConf InfaktConfig
	ac, err := appConfigFromJSON(jsonPath)

	if err != nil {
		return ifConf, err
	}

	return InfaktConfig{
		ApiUrl:            ac.Infakt.ApiUrl,
		ApiKey:            ac.Infakt.ApiKey,
		HourlyRateInGrosz: ac.Infakt.HourlyRateInGrosz,
		ClientId:          ac.Infakt.ClientId,
	}, nil
}
