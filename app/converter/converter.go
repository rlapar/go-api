package converter

import (
	"bytes"

	wkhtmltopdf "github.com/SebastiaanKlippert/go-wkhtmltopdf"
)

func HtmlToPDF(htmlContent string) ([]byte, error) {
	outputBuffer := new(bytes.Buffer)
	pdfGenerator, err := wkhtmltopdf.NewPDFGenerator()
	if err != nil {
		// TODO LOG
		return nil, err
	}

	pdfGenerator.SetOutput(outputBuffer)

	pdfGenerator.AddPage(&wkhtmltopdf.PageReader{
		Input: bytes.NewBufferString(htmlContent),
	})

	err = pdfGenerator.Create()
	if err != nil {
		// TODO LOG
		return nil, err
	}

	return outputBuffer.Bytes(), nil
}

//TODO HtmlToPDF with passed writer
