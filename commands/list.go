package commands

import (
	"fmt"

	"github.com/xssxx/expense-tracker/services"
)

func ListExpense() {
	datasource := services.NewFileDatasource("data", "expenses.csv")
	expenses := datasource.ReadFile()
	if len(expenses) == 0 {
		fmt.Println("No expenses found.")
		return
	}

	println("Expenses:")
	for _, expense := range expenses {
		fmt.Println("ID:", expense.ID)
		fmt.Printf("Amount: %.2f\n", expense.Amount)
		fmt.Println("Description:", expense.Description)
		fmt.Println("Category:", expense.Category)
	}
}
