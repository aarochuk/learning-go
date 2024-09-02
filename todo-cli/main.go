// making a cli todo application using csv files to store the tasks

package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"text/tabwriter"
)

func add(file string) {
	fmt.Printf("add function %s", file)
}

func delete(file string) {
	fmt.Printf("delete function %s", file)
}

func complete(file string) {
	fmt.Printf("complete function %s", file)
}

func list(file string) {
	w := tabwriter.NewWriter(os.Stdout, 0, 2, 3, ' ', 0)

	filer, err := os.Open(file)
	if err != nil {
		panic(err)
	}
	defer filer.Close()

	reader := csv.NewReader(filer)
	reader.FieldsPerRecord = -1
	data, err := reader.ReadAll()
	if err != nil {
		panic(err)
	}

	for pos, row := range data {
		if pos > 0 && row[3] == "No" {
			var str string = ""
			for _, col := range row {
				str += col + "\t"
			}
			fmt.Fprintln(w, str)
		}
	}
	w.Flush()
}

func list_a(file string) {
	w := tabwriter.NewWriter(os.Stdout, 0, 2, 3, ' ', 0)

	filer, err := os.Open(file)
	if err != nil {
		panic(err)
	}
	defer filer.Close()

	reader := csv.NewReader(filer)
	reader.FieldsPerRecord = -1
	data, err := reader.ReadAll()
	if err != nil {
		panic(err)
	}

	for pos, row := range data {
		if pos > 0 {
			var str string = ""
			for _, col := range row {
				str += col + "\t"
			}
			fmt.Fprintln(w, str)
		}
	}
	w.Flush()
}

func main() {
	var option int
	var file string
	fmt.Println("Please enter your list file")
	fmt.Scan(&file)
	fmt.Println("Pick what you want to do")
	fmt.Println("1. list (lists none completed items)")
	fmt.Println("2. list -a (lists items on the list)")
	fmt.Println("3. add (<task details - add later in cobra> adds item to list)")
	fmt.Println("4. complete (<task id - add later in cobra> marks task as complete)")
	fmt.Println("5. delete (<task id - add later in cobra> deletes task from list)")
	fmt.Println("6. end")
	fmt.Scan(&option)
	for {
		if option == 6 {
			break
		}
		switch option {
		case 1:
			list(file)
		case 2:
			list_a(file)
		case 3:
			add(file)
		case 4:
			complete(file)
		case 5:
			delete(file)
		}
		fmt.Scan(&option)
	}
}
