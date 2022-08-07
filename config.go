package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
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

func AppConfigFromEnv() (AppConfig, error) {
	// TODO: research how similar thing would be done with Viper

	var (
		config                  AppConfig
		harvestApiUrl           string
		harvestAuthToken        string
		harvestAccountId        string
		infaktApiUrl            string
		infaktApiKey            string
		infaktHourlyRateInGrosz uint
		infaktClientId          uint
	)

	missingEnvVarError := func(varName string) string {
		return fmt.Sprintf("missing %s env variable", varName)
	}

	harvestApiUrl = os.Getenv("HARVEST_API_URL")
	if harvestApiUrl == "" {
		return config, errors.New(missingEnvVarError("HARVEST_API_URL"))
	}

	harvestAuthToken = os.Getenv("HARVEST_AUTH_TOKEN")
	if harvestAuthToken == "" {
		return config, errors.New(missingEnvVarError("HARVEST_AUTH_TOKEN"))
	}

	harvestAccountId = os.Getenv("HARVEST_ACCOUNT_ID")
	if harvestAccountId == "" {
		return config, errors.New(missingEnvVarError("HARVEST_ACCOUNT_ID"))
	}

	infaktApiUrl = os.Getenv("INFAKT_API_URL")
	if infaktApiUrl == "" {
		return config, errors.New(missingEnvVarError("INFAKT_API_URL"))
	}

	infaktApiKey = os.Getenv("INFAKT_API_KEY")
	if infaktApiKey == "" {
		return config, errors.New(missingEnvVarError("INFAKT_API_KEY"))
	}

	rateParsed, err := strconv.ParseUint(os.Getenv("INFAKT_HOURLY_RATE_IN_GROSZ"), 0, 64)
	if err != nil {
		return config, err
	}
	infaktHourlyRateInGrosz = uint(rateParsed)
	if infaktHourlyRateInGrosz == 0 {
		return config, fmt.Errorf("infakt hourly rate in grosz: %w", err)
	}

	clientIdParsed, err := strconv.ParseUint(os.Getenv("INFAKT_CLIENT_ID"), 0, 64)
	infaktClientId = uint(clientIdParsed)
	if err != nil {
		return config, err
	}
	if infaktClientId == 0 {
		return config, fmt.Errorf("infakt client id cannot be 0: %w", err)
	}

	config = AppConfig{
		Harvest: HarvestConfig{
			ApiUrl:    harvestApiUrl,
			AuthToken: harvestAuthToken,
			AccountId: harvestAccountId,
		},
		Infakt: InfaktConfig{
			ApiUrl:            infaktApiUrl,
			ApiKey:            infaktApiKey,
			HourlyRateInGrosz: infaktHourlyRateInGrosz,
			ClientId:          infaktClientId,
		},
	}
	return config, nil
}
