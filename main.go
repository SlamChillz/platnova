package main

import (
	"encoding/json"
	"github.com/johnfercher/maroto/v2/pkg/config"
	"github.com/johnfercher/maroto/v2/pkg/consts/fontstyle"
	"github.com/johnfercher/maroto/v2/pkg/props"
	"log"
	"os"

	"github.com/slamchillz/platnova/constant"
	"github.com/slamchillz/platnova/pdf"
	"github.com/slamchillz/platnova/types"
)

func extractStatementFromJsonFile() types.AccountStatement {
	file, err := os.Open(constant.DefaultInputMockFile)
	if err != nil {
		log.Fatalln("Error opening file:", err)
	}
	defer func() {
		if err := file.Close(); err != nil {
			log.Fatalln("Error closing file:", err)
		}
		log.Println("File was successfully closed")
	}()
	var statement types.AccountStatement
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&statement)
	if err != nil {
		log.Fatalln("Error decoding file:", err)
	}
	log.Println("Account Statement has been extracted from JSON file")
	return statement
}

func generatePDF() {
	statement := extractStatementFromJsonFile()
	pageNumberSettings := props.PageNumber{
		Pattern: "Page {current} of {total}",
		Place:   props.RightBottom,
		Style:   fontstyle.Bold,
		Size:    4,
	}
	pdfConfig := config.NewBuilder().WithDimensions(120, 200).WithPageNumber(pageNumberSettings)
	page := pdf.New(pdfConfig, statement)
	page.Create(constant.DefaultOutPutFile)
}

func main() {
	generatePDF()
}
