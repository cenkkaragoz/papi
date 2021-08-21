package generator

import (
	"bytes"
	"log"
	"strings"

	p "github.com/SebastiaanKlippert/go-wkhtmltopdf"
)

func NewPdfGenerator(htmlBytesBuffer *bytes.Buffer) {

	pageFile := htmlBytesBuffer.String()
	page := p.NewPageReader(strings.NewReader(pageFile))

	page.PageOptions.EnableLocalFileAccess.Set(true)

	pdfg, err := p.NewPDFGenerator()
	if err != nil {
		log.Fatal(err)
	}

	// htmlFile := `<html>
	// <Body>
	// 	<h2 style={color:'red'}> Hello my pdf</h2>
	// 	<div style={alignText:'center'}>this is the test of pdf generator.</div>
	// </Body>
	// </html>`

	pdfg.Dpi.Set(600)
	pdfg.Orientation.Set(p.OrientationPortrait)
	pdfg.PageSize.Set("A4")

	log.Println("args:", pdfg.ArgString())

	// pdfg.AddPage(p.NewPageReader(strings.NewReader(htmlFile)))

	pdfg.AddPage(page)

	err = pdfg.Create()
	if err != nil {
		log.Fatal(err)
	}

	err = pdfg.WriteFile("./simplesample2.pdf")
	if err != nil {
		log.Fatal(err)
	}

	log.Println("done.")
}
