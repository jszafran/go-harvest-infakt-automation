package main

type InfaktHTTP struct {
	Config InfaktConfig
}

type InvoiceDetails struct {
	ClientId       string `json:"client_id"`
	PaymentMethods string `json:"payment_methods"`
	SalesDate      string `json:"sales_date"`
	InvoiceDate    string `json:"invoice_date"`
}

type DraftInvoiceRequest struct {
	Invoice InvoiceDetails `json:"invoice"`
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
