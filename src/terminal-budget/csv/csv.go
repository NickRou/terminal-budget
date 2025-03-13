package csv

import (
	"encoding/csv"
	"fmt"
	"os"
)

func ReadCSVFile() {
	file, err := os.Open("data.csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		panic(err)
	}

	for _, record := range records {
		fmt.Println(record)
	}
}
