package logs

import (
	"bytes"
	"encoding/base64"
	"os"
	"text/template"

	"github.com/SebastiaanKlippert/go-wkhtmltopdf"
)

func ExecuteTemplateHTML(data interface{}, generatedPath, templatePath string) error {
	err := MkDirByFilePath(generatedPath)
	if err != nil {
		return err
	}

	tmplt, err := template.ParseFiles(templatePath)
	if err != nil {
		return err
	}

	f, err := os.Create(generatedPath)
	if err != nil {
		return err
	}

	err = tmplt.Execute(f, data)
	if err != nil {
		return err
	}

	defer f.Close()

	return nil
}

func HtmlToPdf(filePath string, data []byte, dpi, height, width uint) error {
	err := MkDirByFilePath(filePath)
	if err != nil {
		return err
	}

	pdfg, err := wkhtmltopdf.NewPDFGenerator()
	if err != nil {
		return err
	}
	page := wkhtmltopdf.NewPageReader(bytes.NewReader(data))
	pdfg.AddPage(page)

	pdfg.Orientation.Set(wkhtmltopdf.OrientationPortrait)
	pdfg.Dpi.Set(dpi)
	pdfg.PageHeight.Set(height)
	pdfg.PageWidth.Set(width)

	err = pdfg.Create()
	if err != nil {
		return err
	}

	err = pdfg.WriteFile(filePath)
	if err != nil {
		return err
	}

	return nil
}

func SaveFile(filePath string, content []byte) error {
	err := MkDirByFilePath(filePath)
	if err != nil {
		return err
	}

	f, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer f.Close()

	if _, err := f.Write(content); err != nil {
		return err
	}

	if err := f.Sync(); err != nil {
		return err
	}

	return err
}

func RemoveFile(filePath string) error {
	err := os.Remove(filePath)

	return err
}

func ByteToFile(filePath string, byteArray []byte) error {
	enc := base64.StdEncoding.EncodeToString(byteArray)

	dec, err := base64.StdEncoding.DecodeString(enc)
	if err != nil {
		panic(err)
	}

	err = SaveFile(filePath, dec)
	if err != nil {
		return err
	}

	return err
}
