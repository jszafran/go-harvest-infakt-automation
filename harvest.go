package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type Config struct {
	ApiUrl    string `json:"apiUrl"`
	AuthToken string `json:"authToken"`
	AccountId int    `json:"accountId"`
}

func (c Config) FromJSON(jsonPath string) (Config, error) {
	var cfg Config
	file, err := os.Open(jsonPath)
	if err != nil {
		return cfg, err
	}

	b, err := ioutil.ReadAll(file)

	if err != nil {
		return cfg, err
	}

	err = json.Unmarshal(b, &cfg)

	if err != nil {
		return cfg, err
	}

	return cfg, nil
}

type HTTPClient struct {
	Config Config
}
