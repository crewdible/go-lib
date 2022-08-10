package logs

import (
	"encoding/base64"
	"os"
	"text/template"
)

func ExecuteTemplateHTML(data interface{}, generatedPath, templatePath string) error {
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

func SaveFile(filePath string, content []byte) error {
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
