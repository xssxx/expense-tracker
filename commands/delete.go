package commands

import (
	"flag"
	"fmt"
	"os"

	"github.com/xssxx/expense-tracker/models"
	"github.com/xssxx/expense-tracker/services"
)

func DeleteExpense(args []string) {
	deleteCmd := flag.NewFlagSet("delete", flag.ExitOnError)

	id := deleteCmd.Int("id", 0, "ID of the expense to delete")

	err := deleteCmd.Parse(args)
	if err != nil {
		fmt.Println("Error parsing arguments:", err)
		os.Exit(1)
	}

	// Validate input
	if *id <= 0 {
		fmt.Println("Please provide a valid ID.")
		os.Exit(1)
	}

	// Read and update expenses
	datasource := services.NewFileDatasource("data", "expenses.csv")
	expenses := datasource.ReadFile()

	found := false
	updatedExpenses := make([]models.Expense, 0)
	for _, expense := range expenses {
		if expense.ID == *id {
			found = true
			continue // Skip the expense to be deleted
		}
		updatedExpenses = append(updatedExpenses, expense)
	}

	if !found {
		fmt.Printf("Error: No expense found with ID %d\n", *id)
		os.Exit(1)
	}

	datasource.WriteFile(updatedExpenses)

	fmt.Printf("Deleted Expense with ID: %d\n", *id)
}
