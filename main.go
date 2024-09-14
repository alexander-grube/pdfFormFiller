package main

import (
	"os"

	"github.com/benoitkugler/pdf/model"
	"github.com/benoitkugler/pdf/reader"
)

func main() {
	file, err := os.Open("form.pdf")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	doc, _, err := reader.ParsePDFReader(file, reader.Options{})
	if err != nil {
		panic(err)
	}

	for _, field := range doc.Catalog.AcroForm.Flatten() {
		switch field.Field.FullFieldName() {
		case "field_1":
			field.Field.FT = model.FormFieldText{
				V: "Hello",
			}
		case "field_2":
			field.Field.FT = model.FormFieldText{
				V: "World",
			}
		}
		field.Field.Ff |= model.ReadOnly
	}

	// write the modified document to a new file
	out, err := os.Create("output.pdf")
	if err != nil {
		panic(err)
	}

	err = doc.Write(out, nil)
	if err != nil {
		panic(err)
	}
}
