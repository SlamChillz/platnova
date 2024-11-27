package page

import (
	"github.com/johnfercher/maroto/v2/pkg/components/col"
	"github.com/johnfercher/maroto/v2/pkg/components/row"
	"github.com/johnfercher/maroto/v2/pkg/components/text"
	"github.com/johnfercher/maroto/v2/pkg/consts/align"
	"github.com/johnfercher/maroto/v2/pkg/consts/fontstyle"
	"github.com/johnfercher/maroto/v2/pkg/core"
	"github.com/johnfercher/maroto/v2/pkg/props"
)

func generatePageHeader() core.Row {
	return row.New(16).Add(
		col.New(4).Add(
			text.New("Revolut", props.Text{
				Style: fontstyle.Bold,
				Size:  12,
				Align: align.Left,
			}),
		),
		col.New(4),
		col.New(4).Add(
			text.New("EUR Statement", props.Text{
				Style: fontstyle.Bold,
				Size:  12,
				Align: align.Right,
			}),
			text.New("Generated on the 20 May 2023", props.Text{
				Top:   5,
				Size:  5,
				Align: align.Right,
			}),
			text.New("Revolut Bank UAB", props.Text{
				Top:   7,
				Size:  5,
				Align: align.Right,
			}),
		),
	)
}
