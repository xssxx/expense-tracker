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

	fmt.Println("Expenses:")
	for _, expense := range expenses {
		fmt.Println("ID:", expense.ID)
		fmt.Printf("Amount: %.2f\n", expense.Amount)
		fmt.Println("Description:", expense.Description)
		fmt.Println("Category:", expense.Category)
		fmt.Println("Created At:", expense.CreatedAt.Format("02/01/2006 15:04"))
	}
}
