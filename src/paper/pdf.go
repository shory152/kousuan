package main

import (
	"fmt"

	"github.com/signintech/gopdf"
)

var _ = gopdf.GoPdf{}

type color struct{ r, g, b uint8 }

func (c color) RGB() (uint8, uint8, uint8) { return c.r, c.g, c.b }

var (
	textLineBkColor1 color = color{r: 255, g: 250, b: 205}
	textLineBkColor2 color = color{191, 239, 255} // LightBlue1
	textLineBkColor3 color = color{240, 240, 240}
)

const (
	A4PageWidth    float64 = 595.28
	A4PageHight    float64 = 841.89
	lineWidth      float64 = 510
	textAreaWidth  float64 = 510
	textAreaHight  float64 = 700
	textOffX       float64 = 40
	textOffY       float64 = headerLineOffY + textInterval
	textOverFlow   float64 = tailLineOffY - 20
	textFontSize   int     = 14
	textInterval   float64 = 25
	headerOffX     float64 = 40
	headerOffY     float64 = 40
	headerFontSize int     = 24
	headerLineOffX float64 = 40
	headerLineOffY float64 = 70
	tailLineOffX   float64 = 40
	tailLineOffY   float64 = 780
	tailPageNoOffY float64 = 785
)

func maxString(tm []string) string {
	max := ""
	for _, v := range tm {
		if len(max) < len(v) {
			max = v
		}
	}
	return max
}

func nCols(pdf *gopdf.GoPdf, tm []string, textWidth int) (nCol int, colWidth int) {
	maxStr := maxString(tm)
	textLen, _ := pdf.MeasureTextWidth(maxStr)
	nCol = textWidth / int(textLen)
	colWidth = textWidth / nCol
	return
}

func A4Rect() gopdf.Rect {
	return gopdf.Rect{W: A4PageWidth, H: A4PageHight}
}

func printHeader(pdf *gopdf.GoPdf, hdrText string) error {
	pdf.SetX(headerOffX)
	pdf.SetY(headerOffY)
	if err := pdf.SetFont("simkai", "", headerFontSize); err != nil {
		fmt.Printf("set font: %v\n", err)
		return err
	}
	pdf.Cell(nil, hdrText)

	x := headerLineOffX
	y := headerLineOffY
	pdf.SetX(x)
	pdf.SetY(y)
	pdf.Line(x, y, x+lineWidth, y)

	pdf.SetFont("simkai", "", textFontSize)
	pdf.SetX(textOffX)
	pdf.SetY(textOffY)

	return nil
}

func printTail(pdf *gopdf.GoPdf, pgNo int) error {
	x := tailLineOffX
	y := tailLineOffY
	pdf.Line(x, y, x+lineWidth, y)
	if pgNo > 0 {
		pnStr := fmt.Sprintf("%d", pgNo)
		pnStrLen, _ := pdf.MeasureTextWidth(pnStr)
		pdf.SetX(textOffX + textAreaWidth/2 - pnStrLen/2)
		pdf.SetY(tailPageNoOffY)
		pdf.Cell(nil, pnStr)
	}
	return nil
}

func getFontPath(fname string) string {
	return "../font/" + fname
}

func printpdf(tm []string, hdrText string, fname string) {
	pdfConf := gopdf.Config{}
	pdfConf.PageSize = A4Rect()
	pdf := &gopdf.GoPdf{}
	pdf.Start(pdfConf)

	if err := pdf.AddTTFFont("simkai", getFontPath("simkai.ttf")); err != nil {
		fmt.Printf("add font: %v\n", err)
		return
	}

	pdf.AddPage()

	printHeader(pdf, hdrText)

	//x := textOffX
	y := pdf.GetY()
	nPage := 1
	nCol, colWidth := nCols(pdf, tm, int(textAreaWidth))
	for i, val := range tm {

		// draw bgcolor
		if i%nCol == 0 {
			r, g, b := textLineBkColor3.RGB()
			if i%2 == 0 {
				pdf.SetLineWidth(textInterval)
				pdf.SetStrokeColor(r, g, b)
				pdf.Line(headerLineOffX, y+5, headerLineOffX+lineWidth, y+5)
				// restore line width and stroke color
				pdf.SetLineWidth(1)
				pdf.SetStrokeColor(0, 0, 0)
			}
		}

		// draw text
		pdf.SetX(textOffX + float64((i%nCol)*colWidth))
		pdf.SetY(y)

		pdf.Cell(nil, val)

		if i+1 == len(tm) {
			break
		}

		if (i+1)%nCol == 0 {
			y += textInterval
			if y > textOverFlow {
				printTail(pdf, nPage)
				nPage++
				pdf.AddPage()
				printHeader(pdf, hdrText)
				y = pdf.GetY()
			}
		}
	}

	if y < textOverFlow {
		printTail(pdf, nPage)
	}

	pdf.WritePdf(fname)
}
