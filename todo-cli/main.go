// making a cli todo application using csv files to store the tasks

package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

func main() {
	file, err := os.OpenFile("data.csv", os.O_APPEND, 0644)
	if err != nil {
		panic(err)
	}

	defer file.Close()
	writer := csv.NewWriter(file)
	defer writer.Flush()

	row := []string{"2", "finish go project", "2024-09-01", "No"}
	err = writer.Write(row)
	if err != nil {
		panic(err)
	}

	reader := csv.NewReader(file)
	reader.FieldsPerRecord = -1
	data, err := reader.ReadAll()
	if err != nil {
		panic(err)
	}

	for _, row := range data {
		for _, col := range row {
			fmt.Printf("%s", col)
		}
		fmt.Println()
	}

}
