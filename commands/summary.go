package commands

import (
	"fmt"

	"github.com/xssxx/expense-tracker/services"
)

func SummaryExpense() {
	datasource := services.NewFileDatasource("data", "expenses.csv")
	expenses := datasource.ReadFile()
	if len(expenses) == 0 {
		fmt.Println("No expenses found.")
		return
	}

	total := 0.0
	for _, expense := range expenses {
		total += expense.Amount
	}

	fmt.Printf("Total Expenses: %.2f\n", total)
}
