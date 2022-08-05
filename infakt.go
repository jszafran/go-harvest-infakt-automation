package main

import (
	"errors"
	"fmt"
	"github.com/shopspring/decimal"
)

type InfaktHTTP struct {
	Config InfaktConfig
}

type ServiceLine struct {
	Name              string          `json:"name"`
	TaxSymbol         string          `json:"tax_symbol"`
	Quantity          decimal.Decimal `json:"quantity"`
	NetPrice          uint            `json:"net_price"`
	FlatRateTaxSymbol string          `json:"flat_rate_tax_symbol"`
}

type InvoiceDetails struct {
	ClientId       uint   `json:"client_id"`
	PaymentMethods string `json:"payment_methods"`
	SalesDate      string `json:"sales_date"`
	InvoiceDate    string `json:"invoice_date"`
	Services       []ServiceLine
}

type DraftInvoiceRequest struct {
	Invoice InvoiceDetails `json:"invoice"`
}

func NewInfaktClient(configPath string) (InfaktHTTP, error) {
	var infakt InfaktHTTP
	infaktConf, err := InfaktConfigFromJSON(configPath)
	if err != nil {
		return infakt, err
	}

	return InfaktHTTP{
		Config: infaktConf,
	}, nil
}

func (i InfaktHTTP) generateServicesFromMonthlySummary(ms MonthlySummary) []ServiceLine {
	svs := make([]ServiceLine, 0)
	for client, hrs := range ms {
		sl := ServiceLine{
			Name:              fmt.Sprintf("Usługi programistyczne - %s", client),
			TaxSymbol:         "23",
			Quantity:          hrs,
			NetPrice:          i.Config.HourlyRateInGrosz,
			FlatRateTaxSymbol: "12",
		}
		svs = append(svs, sl)
	}
	return svs
}

func (i InfaktHTTP) CreateDraftInvoice(month int, year int, ms MonthlySummary) error {
	if len(ms) == 0 {
		return errors.New("empty monthly summary")
	}
	mthRng, err := AsMonthRange(month, year)
	if err != nil {
		return err
	}

	svs := i.generateServicesFromMonthlySummary(ms)
	reqData := DraftInvoiceRequest{
		Invoice: InvoiceDetails{
			ClientId:       i.Config.ClientId,
			PaymentMethods: "transfer",
			SalesDate:      mthRng.End,
			InvoiceDate:    mthRng.End,
			Services:       svs,
		},
	}
	fmt.Printf("%+v\n", reqData)
	return nil
}
