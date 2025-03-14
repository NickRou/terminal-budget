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
	case "Chase":
		displayChaseCSV(records)
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
	// there is NO HEADER in the csv file
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

func displayChaseCSV(records [][]string) {
	// records for chase csv will come in as
	// [Transaction Date Post Date Description Category Type Amount Memo]
	columns := []table.Column{
		{Title: "Transaction Date", Width: 15},
		{Title: "Post Date", Width: 10},
		{Title: "Description", Width: 20},
		{Title: "Category", Width: 20},
		{Title: "Type", Width: 10},
		{Title: "Amount", Width: 15},
	}
	rows := []table.Row{}

	// exlcuding memo (column 6)
	for _, row := range records[1:] {
		date, postDate, description, category, transactionType, amount := row[0], row[1], row[2], row[3], row[4], row[5]
		rows = append(rows, table.Row{date, postDate, description, category, transactionType, amount})
	}
	DisplayTable(columns, rows)
}
