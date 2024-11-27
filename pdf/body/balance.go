package body

import (
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

func GenerateBalanceSummary(balanceSummary types.BalanceSummary, currencySymbol string) []core.Row {
	var rows []core.Row
	sudHeader := row.New(2).Add(
		col.New(12).Add(
			text.New("Balance Summary", props.Text{
				Style: fontstyle.Bold,
				Size:  5,
				Align: align.Left,
			}),
		),
	)
	rows = append(rows, sudHeader, row.New(3))
	rows = append(rows, generateBalanceSummaryTable(balanceSummary.Data, currencySymbol)...)
	rows = append(rows, row.New(0.5), row.New(1).Add(
		text.NewCol(12, balanceSummary.Note, props.Text{
			Style: fontstyle.Normal,
			Size:  2.4,
			Align: align.Left,
		}),
	))
	return rows
}

func generateBalanceSummaryTable(balanceSummaryData []types.BalanceData, currencySymbol string) []core.Row {
	var rows []core.Row
	var totalOpenBalance float64 = 0
	var totalMoneyOut float64 = 0
	var totalMoneyIn float64 = 0
	var totalClosingBalance float64 = 0
	tableHeader := row.New(1.5).Add(
		col.New(4).Add(
			text.New("Product", props.Text{
				Style: fontstyle.Bold,
				Size:  3,
				Align: align.Left,
			}),
		),
		col.New(2).Add(
			text.New("Opening balance", props.Text{
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
			text.New("Closing balance", props.Text{
				Style: fontstyle.Bold,
				Size:  3,
				Align: align.Right,
			}),
		),
	)
	rows = append(rows, tableHeader)
	rows = append(rows, row.New(0.75).Add(
		line.NewCol(12, props.Line{
			Thickness: 0.15,
			//Color:       &props.Color{Red: 200, Green: 200, Blue: 200},
			SizePercent: 100,
			Style:       linestyle.Solid,
		}),
	))

	for _, data := range balanceSummaryData {
		totalOpenBalance += data.OpenBalance
		totalMoneyOut += data.MoneyOut
		totalMoneyIn += data.MoneyIn
		totalClosingBalance += data.ClosingBalance

		tmpRow := row.New(0.75).Add(
			col.New(4).Add(
				text.New(data.Product, props.Text{
					Style: fontstyle.Normal,
					Align: align.Left,
					Size:  3,
				}),
			),
			col.New(2).Add(
				text.New(util.AddCurrencySymbol(currencySymbol, util.FormatNumber(data.OpenBalance)), props.Text{
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
				text.New(util.AddCurrencySymbol(currencySymbol, util.FormatNumber(data.ClosingBalance)), props.Text{
					Style: fontstyle.Normal,
					Size:  3,
					Align: align.Right,
				}),
			),
		)
		rows = append(rows, tmpRow)
		rows = append(rows, row.New(1))
		rows = append(rows, row.New(1).Add(
			line.NewCol(12, props.Line{
				Thickness: 0.15,
				//Color:       &props.Color{Red: 200, Green: 200, Blue: 200},
				SizePercent: 100,
				Style:       linestyle.Solid,
			}),
		))
	}
	tmpRow := row.New(2).Add(
		col.New(4).Add(
			text.New("Total", props.Text{
				Style: fontstyle.Normal,
				Align: align.Left,
				Size:  3,
			}),
		),
		col.New(2).Add(
			text.New(util.AddCurrencySymbol(currencySymbol, util.FormatNumber(totalOpenBalance)), props.Text{
				Style: fontstyle.Normal,
				Size:  3,
				Align: align.Left,
			}),
		),
		col.New(2).Add(
			text.New(util.AddCurrencySymbol(currencySymbol, util.FormatNumber(totalMoneyOut)), props.Text{
				Style: fontstyle.Normal,
				Size:  3,
				Align: align.Left,
			}),
		),
		col.New(2).Add(
			text.New(util.AddCurrencySymbol(currencySymbol, util.FormatNumber(totalMoneyIn)), props.Text{
				Style: fontstyle.Normal,
				Size:  3,
				Align: align.Left,
			}),
		),
		col.New(2).Add(
			text.New(util.AddCurrencySymbol(currencySymbol, util.FormatNumber(totalClosingBalance)), props.Text{
				Style: fontstyle.Normal,
				Size:  3,
				Align: align.Right,
			}),
		),
	)
	rows = append(rows, tmpRow)
	return rows
}
