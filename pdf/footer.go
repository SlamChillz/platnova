package pdf

import (
	"github.com/johnfercher/maroto/v2/pkg/components/code"
	"github.com/johnfercher/maroto/v2/pkg/components/col"
	"github.com/johnfercher/maroto/v2/pkg/components/row"
	"github.com/johnfercher/maroto/v2/pkg/components/text"
	"github.com/johnfercher/maroto/v2/pkg/consts/align"
	"github.com/johnfercher/maroto/v2/pkg/consts/fontstyle"
	"github.com/johnfercher/maroto/v2/pkg/core"
	"github.com/johnfercher/maroto/v2/pkg/props"
)

func generateFooter() []core.Row {
	var rows []core.Row
	hereLink := "https://google.com"
	topRow := row.New(6).Add(
		code.NewQrCol(1, "qrcode", props.Rect{
			Left:    1,
			Percent: 100,
		}),
		//col.New(0.25),
		col.New(2).Add(
			text.New("Report lost or stolen card", props.Text{
				Style: fontstyle.Bold,
				Size:  3.6,
				Align: align.Left,
				Color: &props.Color{Red: 100, Green: 100, Blue: 100},
			}),
			text.New("+370 5 214 3608", props.Text{
				Style: fontstyle.Normal,
				Size:  3,
				Align: align.Left,
				Top:   1.5,
				Color: &props.Color{Red: 100, Green: 100, Blue: 100},
			}),
			text.New("Get help directly in app", props.Text{
				Style: fontstyle.Bold,
				Size:  3.6,
				Align: align.Left,
				Top:   3,
				Color: &props.Color{Red: 100, Green: 100, Blue: 100},
			}),
			text.New("Scan the QR code", props.Text{
				Style: fontstyle.Normal,
				Size:  3,
				Align: align.Left,
				Top:   4.5,
				Color: &props.Color{Red: 100, Green: 100, Blue: 100},
			}),
		),
		col.New(9).Add(
			text.New("Revolut Bank UAB is a credit institution licensed in the Republic of Lithuania with company number 304580906 and authorisation code LB002119, and whose registered office is at Konstituocijos 21B LT-08130 Vihius, the Republic of Lithuania. We are licensed and regulated by the Bank of Lithuania and the European Central Bank. The deposits are protected by Lithuanian Deposit Insurance Systems but some exceptions may apply. Please refer to our Deposit Insurance Information document ", props.Text{
				Style: fontstyle.Normal,
				Size:  2.6,
				Align: align.Left,
			}),
			text.New("here", props.Text{
				Style: fontstyle.Normal,
				Size:  2.6,
				Top:   1.85,
				Left:  53.9,
				//Align:     align.Left,
				Hyperlink: &hereLink,
			}),
			text.New(". More information on deposit insurance of the ", props.Text{
				Style: fontstyle.Normal,
				Size:  2.6,
				Top:   1.85,
				Left:  55.75,
			}),
			text.New("State Enterprise Deposit and investment insurance (VI investment insurance draudimas) is available at ", props.Text{
				Style: fontstyle.Normal,
				Size:  2.6,
				Top:   2.85,
			}),
			text.New("www.idruadimas.it", props.Text{
				Style: fontstyle.Normal,
				Size:  2.6,
				Top:   2.85,
				Left:  42,
				//Align:     align.Left,
				Hyperlink: &hereLink,
			}),
			text.New(". If you have any questions, please reach out to us via the ", props.Text{
				Style: fontstyle.Normal,
				Size:  2.6,
				Top:   2.85,
				Left:  49.5,
			}),
			text.New("in-app chat in the Revolt app.", props.Text{
				Style: fontstyle.Normal,
				Size:  2.6,
				Top:   4,
			}),
		),
	)
	rows = append(rows, topRow)
	rows = append(rows, row.New(2).Add(
		col.New(3).Add(
			text.New("© 2023 Revolut Bank UAB", props.Text{
				Style: fontstyle.Bold,
				Size:  4,
				Align: align.Left,
				Top:   2.5,
			}),
		),
	))
	return rows
}
