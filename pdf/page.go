package page

import (
	"fmt"
	"github.com/johnfercher/maroto/v2"
	"github.com/johnfercher/maroto/v2/pkg/config"
	"github.com/johnfercher/maroto/v2/pkg/core"
	"github.com/slamchillz/platnova/page/body"
	"github.com/slamchillz/platnova/types"
	"log"
	"os"
)

type Page struct {
	maroto core.Maroto
	types.AccountStatement
}

func New(config config.Builder, statement types.AccountStatement) *Page {
	return &Page{
		maroto:           maroto.New(config.Build()),
		AccountStatement: statement,
	}
}

func (p *Page) assemble() {
	err := p.maroto.RegisterHeader(generatePageHeader())

	if err != nil {
		log.Fatalln("Error registering header:", err)
	}
	p.maroto.AddRow(6)
	p.maroto.AddRows(body.GenerateAddressRow(p.AccountStatement.BillingAddress)...)
	p.maroto.AddRows(body.GenerateAccountInfo(p.AccountStatement.Accounts)...)
	p.maroto.AddRow(10)
	p.maroto.AddRows(body.GenerateBalanceSummary(p.AccountStatement.BalanceSummary, p.AccountStatement.CurrencySymbol)...)
	p.maroto.AddRow(10)
	p.maroto.AddRows(body.GenerateTransactions(p.AccountStatement.Transactions, p.AccountStatement.CurrencySymbol)...)
	p.maroto.AddRow(15)
	err = p.maroto.RegisterFooter(generateFooter()...)
	if err != nil {
		log.Fatalln("Error registering header:", err)
	}
}

func (p *Page) Create(outPutFile string) {
	document, err := p.maroto.Generate()
	if err != nil {
		log.Fatalln("Error generating document:", err)
	}

	filename := fmt.Sprintf("report.pdf")

	// Check if temp folder exists, if not create it
	if _, err := os.Stat("temp"); os.IsNotExist(err) {
		err = os.Mkdir("temp", 0755)
		if err != nil {
			log.Println("Error creating directory:", err)
		}
	}

	err = document.Save(outPutFile)
	if err != nil {
		log.Println("Unable to save file:", err)
	}
}
