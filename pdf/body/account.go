package body

import (
	"github.com/johnfercher/maroto/v2/pkg/components/col"
	"github.com/johnfercher/maroto/v2/pkg/components/row"
	"github.com/johnfercher/maroto/v2/pkg/components/text"
	"github.com/johnfercher/maroto/v2/pkg/consts/align"
	"github.com/johnfercher/maroto/v2/pkg/consts/fontstyle"
	"github.com/johnfercher/maroto/v2/pkg/core"
	"github.com/johnfercher/maroto/v2/pkg/props"

	"github.com/slamchillz/platnova/types"
)

func GenerateAccountInfo(accounts []types.Account) []core.Row {
	var rows []core.Row
	for _, account := range accounts {
		rows = append(rows, generateIbanRow(account.Iban), generateBicRow(account.Bic))
		if account.Note != "" {
			rows = append(rows, generateIbanNoteRow(account.Note))
		} else {
			rows = append(rows, row.New(1))
		}
	}
	return rows
}

func generateIbanRow(value string) core.Row {
	tmpRow := row.New(4).Add(
		col.New(7),
		col.New(1).Add(
			text.New("IBAN", props.Text{
				Style: fontstyle.Bold,
				Size:  4.5,
				Align: align.Left,
				Top:   1.5,
			}),
		),
		col.New(4).Add(
			text.New(value, props.Text{
				Style: fontstyle.Normal,
				Size:  4.5,
				Align: align.Left,
				Top:   1.5,
			}),
		),
	)
	return tmpRow
}

func generateBicRow(value string) core.Row {
	tmpRow := row.New(2).Add(
		col.New(7),
		col.New(1).Add(
			text.New("BIC", props.Text{
				Style: fontstyle.Bold,
				Size:  4.5,
				Align: align.Left,
				Top:   0,
			}),
		),
		col.New(4).Add(
			text.New(value, props.Text{
				Style: fontstyle.Normal,
				Size:  4.5,
				Align: align.Left,
				Top:   0,
			}),
		),
	)
	return tmpRow
}

func generateIbanNoteRow(value string) core.Row {
	tmpRow := row.New(5).Add(
		col.New(8),
		col.New(4).Add(
			text.New(value, props.Text{
				Style: fontstyle.Normal,
				Size:  4.5,
				Align: align.Left,
				//Bottom: 2,
			}),
		),
	)
	return tmpRow
}
