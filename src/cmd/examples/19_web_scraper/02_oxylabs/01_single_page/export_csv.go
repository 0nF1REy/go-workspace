package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

func ExportCSV(products []Product, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	records := [][]string{{"Nome", "Preço", "URL"}}

	for _, p := range products {
		records = append(records, []string{p.Name, p.Price, p.URL})
	}

	if err := writer.WriteAll(records); err != nil {
		return err
	}

	if err := writer.Error(); err != nil {
		return err
	}

	fmt.Printf("Exportação dos dados para '%s' foi um sucesso!\n", filename)
	return nil
}
