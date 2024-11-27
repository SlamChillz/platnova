package main

import (
	"log"
	"time"
)

// Custom Time with unique behaviours
type PlatnovaTime time.Time

// UnmarshalJSON Custom UnmarshalJSON for PlatnovaTime
func (p *PlatnovaTime) UnmarshalJSON(data []byte) error {
	// Remove quotes from the JSON string
	str := string(data)
	str = str[1 : len(str)-1]

	// Define the layout of your date string
	layout := "2006-01-02T15:04:05Z"

	// Parse the string into a time.Time
	parsedTime, err := time.Parse(layout, str)
	if err != nil {
		log.Fatalf("Error parsing PlatnovaTime: %w", err)
	}

	// Assign parsed time to the custom type
	*p = PlatnovaTime(parsedTime)
	return nil
}

// Generates date format used for Account Transaction Table sub-heading
func (p *PlatnovaTime) HeaderFormat() string {
	t := time.Time(*p)
	// Format the date to "23 March 2012"
	output := t.Format("2 January 2006")
	// Return the formatted string
	return output
}

// Generates date format used for Account Transaction Table `Date` column data
func (p *PlatnovaTime) TableFormat() string {
	t := time.Time(*p)
	// Format the date to "23 March 2012"
	output := t.Format("2 Jan 2006")
	// Return the formatted string
	return output
}

// BillingAddress Billing Address data
type BillingAddress struct {
	Name       string `json:"name"`
	Address    string `json:"address"`
	City       string `json:"city"`
	State      string `json:"state"`
	Country    string `json:"country"`
	PostalCode string `json:"postal_code"`
}

// Account customer account details
type Account struct {
	Iban string `json:"iban"`
	Bic  string `json:"bic"`
	Note string `json:"note"`
}

// BalanceData customer account info
type BalanceData struct {
	Product        string  `json:"product"`
	OpenBalance    float64 `json:"open_balance"`
	MoneyOut       float64 `json:"money_out"`
	MoneyIn        float64 `json:"money_in"`
	ClosingBalance float64 `json:"closing_balance"`
}

// BalanceSummary overview of customer balance
type BalanceSummary struct {
	Data []BalanceData `json:"data"`
	Note string        `json:"note"`
}

// TransactionRecord represents a single transaction record
type TransactionRecord struct {
	Date            PlatnovaTime `json:"date"`
	Description     string       `json:"description"`
	DescriptionNote string       `json:"description_note"`
	MoneyOut        float64      `json:"money_out"`
	MoneyIn         float64      `json:"money_in"`
	Balance         float64      `json:"balance"`
}

// TransactionData houses all transaction information
type TransactionData struct {
	StartDate PlatnovaTime        `json:"start_date"`
	EndDate   PlatnovaTime        `json:"end_date"`
	Records   []TransactionRecord `json:"records"`
}

// AccountStatement is the golang equivalent of the json data structure
type AccountStatement struct {
	BillingAddress BillingAddress  `json:"billing_address"`
	Accounts       []Account       `json:"accounts"`
	CurrencySymbol string          `json:"currency_symbol"`
	BalanceSummary BalanceSummary  `json:"balance_summary"`
	Transactions   TransactionData `json:"transactions"`
}
