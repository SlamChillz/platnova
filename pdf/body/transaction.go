package body

import (
	"fmt"
	"github.com/johnfercher/maroto/v2/pkg/components/col"
	"github.com/johnfercher/maroto/v2/pkg/components/line"
	"github.com/johnfercher/maroto/v2/pkg/components/row"
	"github.com/johnfercher/maroto/v2/pkg/components/text"
	"github.com/johnfercher/maroto/v2/pkg/consts/align"
	"github.com/johnfercher/maroto/v2/pkg/consts/fontstyle"
	"github.com/johnfercher/maroto/v2/pkg/consts/linestyle"
	"github.com/johnfercher/maroto/v2/pkg/core"
	"github.com/johnfercher/maroto/v2/pkg/props"

	"github.com/slamchillz/platnova/types"
	"github.com/slamchillz/platnova/util"
)

func GenerateTransactions(td types.TransactionData, currencySymbol string) []core.Row {
	var rows []core.Row
	sudHeader := row.New(5).Add(
		col.New(12).Add(
			text.New(fmt.Sprintf("Account transactions from %s to %s", td.StartDate.HeaderFormat(), td.EndDate.HeaderFormat()), props.Text{
				Style: fontstyle.Bold,
				Size:  5,
				Align: align.Left,
			}),
		),
	)
	rows = append(rows, sudHeader, row.New(2))
	rows = append(rows, generateAccountTransactionTable(td.Records, currencySymbol)...)
	return rows
}

func generateAccountTransactionTable(transactionRecords []types.TransactionRecord, currencySymbol string) []core.Row {
	var rows []core.Row
	tableHeader := row.New(2).Add(
		col.New(2).Add(
			text.New("Date", props.Text{
				Style: fontstyle.Bold,
				Size:  3,
				Align: align.Left,
			}),
		),
		col.New(4).Add(
			text.New("Description", props.Text{
				Style: fontstyle.Bold,
				Size:  3,
				Align: align.Left,
			}),
		),
		col.New(2).Add(
			text.New("Money out", props.Text{
				Style: fontstyle.Bold,
				Size:  3,
				Align: align.Left,
			}),
		),
		col.New(2).Add(
			text.New("Money in", props.Text{
				Style: fontstyle.Bold,
				Size:  3,
				Align: align.Left,
			}),
		),
		col.New(2).Add(
			text.New("Balance", props.Text{
				Style: fontstyle.Bold,
				Size:  3,
				Align: align.Right,
			}),
		),
	)
	rows = append(rows, tableHeader)
	rows = append(rows, row.New(1).Add(
		line.NewCol(12, props.Line{
			Thickness: 0.15,
			//Color:       &props.Color{Red: 200, Green: 200, Blue: 200},
			SizePercent: 100,
			Style:       linestyle.Solid,
		}),
	))
	for _, data := range transactionRecords {
		fmt.Println(data.MoneyOut)
		tmpRow := row.New(2).Add(
			col.New(2).Add(
				text.New(data.Date.TableFormat(), props.Text{
					Style: fontstyle.Normal,
					Align: align.Left,
					Size:  3,
				}),
			),
			col.New(4).Add(
				text.New(data.Description, props.Text{
					Style: fontstyle.Normal,
					Size:  3,
					Align: align.Left,
				}),
			),
			col.New(2).Add(
				text.New(util.AddCurrencySymbol(currencySymbol, util.FormatNumber(data.MoneyOut)), props.Text{
					Style: fontstyle.Normal,
					Size:  3,
					Align: align.Left,
				}),
			),
			col.New(2).Add(
				text.New(util.AddCurrencySymbol(currencySymbol, util.FormatNumber(data.MoneyIn)), props.Text{
					Style: fontstyle.Normal,
					Size:  3,
					Align: align.Left,
				}),
			),
			col.New(2).Add(
				text.New(util.AddCurrencySymbol(currencySymbol, util.FormatNumber(data.Balance)), props.Text{
					Style: fontstyle.Normal,
					Size:  3,
					Align: align.Right,
				}),
			),
		)
		rows = append(rows, tmpRow)
		rows = append(rows, row.New(1))
		if data.DescriptionNote != "" {
			rows = append(rows, row.New(1).Add(
				col.New(2),
				col.New(10).Add(
					text.New(data.DescriptionNote, props.Text{
						Style: fontstyle.Normal,
						Size:  1.5,
						Align: align.Left,
					}),
				),
			),
			)
		}
		rows = append(rows, row.New(1).Add(
			line.NewCol(12, props.Line{
				Thickness: 0.15,
				//Color:       &props.Color{Red: 200, Green: 200, Blue: 200},
				SizePercent: 100,
				Style:       linestyle.Solid,
			}),
		))
	}
	return rows
}
