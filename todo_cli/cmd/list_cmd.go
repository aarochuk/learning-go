package cmd

import (
	"encoding/csv"
	"fmt"
	"os"
)

func list_cmd() (string, error) {
	// Open the CSV file
	file, err := os.Open("../data.csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	// Read the CSV data
	reader := csv.NewReader(file)
	reader.FieldsPerRecord = -1 // Allow variable number of fields
	data, err := reader.ReadAll()
	if err != nil {
		panic(err)
	}

	// Print the CSV data
	for _, row := range data {
		for _, col := range row {
			fmt.Printf("%s,", col)
		}
		fmt.Println()
	}
}
