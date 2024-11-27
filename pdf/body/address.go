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

func GenerateAddressRow(address types.BillingAddress) []core.Row {
	var rows []core.Row
	name := row.New(1).Add(
		col.New(5).Add(
			text.New(address.Name, props.Text{
				Style: fontstyle.Bold,
				Size:  8,
				Align: align.Left,
			}),
		),
	)
	rows = append(rows, name)
	rows = append(rows, row.New(20).Add(
		col.New(4).Add(
			text.New(address.Address, props.Text{
				Style: fontstyle.Bold,
				Size:  6,
				Align: align.Left,
				Top:   5,
			}),
			text.New(address.City, props.Text{
				Style: fontstyle.Bold,
				Size:  6,
				Align: align.Left,
				Top:   8,
			}),
			text.New(address.State, props.Text{
				Style: fontstyle.Bold,
				Size:  6,
				Align: align.Left,
				Top:   11,
			}),
			text.New(address.Country, props.Text{
				Style: fontstyle.Bold,
				Size:  6,
				Align: align.Left,
				Top:   14,
			}),
			text.New(address.PostalCode, props.Text{
				Style: fontstyle.Bold,
				Size:  6,
				Align: align.Left,
				Top:   17,
			}),
		)),
	)
	return rows
}
