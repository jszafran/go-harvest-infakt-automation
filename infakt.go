package main

type InfaktHTTP struct {
	Config InfaktConfig
}

func (i InfaktHTTP) NewInfaktClient(configPath string) (InfaktHTTP, error) {
	var ifkt InfaktHTTP
	ifConf, err := InfaktConfigFromJSON(configPath)
	if err != nil {
		return ifkt, err
	}

	return InfaktHTTP{
		Config: ifConf,
	}, nil
}
