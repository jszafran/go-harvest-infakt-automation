package main

type HarvestHTTP struct {
	Config HarvestConfig
}

func (h HarvestHTTP) NewHarvestClient(configPath string) (HarvestHTTP, error) {
	var hv HarvestHTTP
	hvConf, err := HarvestConfigFromJSON(configPath)
	if err != nil {
		return hv, err
	}

	return HarvestHTTP{
		Config: hvConf,
	}, nil
}
