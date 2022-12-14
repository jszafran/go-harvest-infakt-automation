package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/shopspring/decimal"
	"log"
	"net/http"
	"time"
)

type InfaktHTTP struct {
	Config InfaktConfig
	client http.Client
}

type ServiceLine struct {
	Name              string          `json:"name"`
	TaxSymbol         string          `json:"tax_symbol"`
	Quantity          decimal.Decimal `json:"quantity"`
	UnitNetPrice      uint            `json:"unit_net_price"`
	FlatRateTaxSymbol string          `json:"flat_rate_tax_symbol"`
}

type InvoiceDetails struct {
	ClientId       uint          `json:"client_id"`
	PaymentMethods string        `json:"payment_method"`
	SalesDate      string        `json:"sales_date"`
	InvoiceDate    string        `json:"invoice_date"`
	Services       []ServiceLine `json:"services"`
}

type DraftInvoiceRequest struct {
	Invoice InvoiceDetails `json:"invoice"`
}

type InvoiceDraftResponse struct {
	Id     int    `json:"id"`
	Number string `json:"number,"`
}

func NewInfaktClient(config AppConfig) InfaktHTTP {
	infaktConf := config.Infakt
	httpClient := http.Client{Timeout: time.Second * 5}
	return InfaktHTTP{
		Config: infaktConf,
		client: httpClient,
	}
}

func (i InfaktHTTP) postRequest(path string, data DraftInvoiceRequest) (*http.Response, error) {
	if string(path[0]) == "/" {
		path = path[1:]
	}

	url := fmt.Sprintf("%s/%s", i.Config.ApiUrl, path)

	jsonBody, err := json.Marshal(data)
	if err != nil {
		log.Fatal(err)
	}
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(jsonBody))
	if err != nil {
		log.Fatalf("Failed to create request: %v", err)
	}

	req.Header.Set("X-inFakt-ApiKey", i.Config.ApiKey)
	req.Header.Set("Content-Type", "application/json")

	return i.client.Do(req)
}
func (i InfaktHTTP) generateServicesFromMonthlySummary(ms MonthlySummary) []ServiceLine {
	svs := make([]ServiceLine, 0)
	for client, hrs := range ms {
		if client != TimeOffClient {
			sl := ServiceLine{
				Name:              fmt.Sprintf("Usługi programistyczne - %s", client),
				TaxSymbol:         "23",
				Quantity:          hrs,
				UnitNetPrice:      i.Config.HourlyRateInGrosz,
				FlatRateTaxSymbol: "12",
			}
			svs = append(svs, sl)
		}
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
	resp, err := i.postRequest("invoices.json", reqData)
	if err != nil {
		return err
	}

	var response InvoiceDraftResponse
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return err
	}

	log.Printf("Successfully created an invoice draft. Details: %+v\n", response)
	log.Printf("Link: https://app.infakt.pl/app/faktury/%d\n", response.Id)
	return nil
}
