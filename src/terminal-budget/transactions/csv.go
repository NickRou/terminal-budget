package transactions

import (
	"encoding/csv"
	"fmt"
	"os"

	"time"

	"github.com/charmbracelet/bubbles/table"
	"github.com/charmbracelet/huh/spinner"
)

// public methods -------------------------------------------------------------------------------------

func DisplayTransactionsFromCSV(fileName string, institutionName string) {
	records := [][]string{}

	action := func() {
		time.Sleep(2 * time.Second) // add some delay for fun
		records = readCSV(fileName)
	}
	_ = spinner.New().Title("Loading your transactions...").Action(action).Run()

	displayCSVRecords(institutionName, records)
}

// private methods ------------------------------------------------------------------------------------

func readCSV(fileName string) [][]string {
	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)

	records, err := reader.ReadAll()
	if err != nil {
		panic(err)
	}

	return records
}

func displayCSVRecords(institutionName string, records [][]string) {
	switch institutionName {
	case "American Express":
		displayAmericanExpressCSV(records)
	default:
		fmt.Println("Institution not supported. Defaulting to regular print.")
		defaultDisplayCSV(records)
	}
}

func defaultDisplayCSV(records [][]string) {
	for _, row := range records {
		fmt.Println(row)
	}
}

func displayAmericanExpressCSV(records [][]string) {
	// records for american express csv will come in as
	// [date transaction amount]
	columns := []table.Column{
		{Title: "Date", Width: 10},
		{Title: "Transaction", Width: 30},
		{Title: "Amount", Width: 20},
	}
	rows := []table.Row{}

	for _, row := range records {
		date, transaction, amount := row[0], row[1], row[2]
		rows = append(rows, table.Row{date, transaction, amount})
	}
	DisplayTable(columns, rows)
}
