package main

import (
	"github.com/NickRou/terminal-budget/form"
	csv "github.com/NickRou/terminal-budget/transactions"
)

func main() {
	institution := form.DisplayInstitutionForm()
	csv.DisplayTransactionsFromCSV(institution.File, institution.Name)
}
