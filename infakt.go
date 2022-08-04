package main

type InfaktHTTP struct {
	Config InfaktConfig
}

func (i InfaktHTTP) NewInfaktClient(configPath string) (InfaktHTTP, error) {
	var infakt InfaktHTTP
	infaktConf, err := InfaktConfigFromJSON(configPath)
	if err != nil {
		return infakt, err
	}

	return InfaktHTTP{
		Config: infaktConf,
	}, nil
}
