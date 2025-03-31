package commands

import (
	"flag"
	"fmt"
	"os"

	"github.com/xssxx/expense-tracker/models"
	"github.com/xssxx/expense-tracker/services"
)

func AddExpense(args []string) {
	addCmd := flag.NewFlagSet("add", flag.ExitOnError)

	amount := addCmd.Float64("amount", 0.0, "Amount of the expense")
	description := addCmd.String("description", "", "Description of the expense")
	category := addCmd.String("category", "", "Category of the expense")

	addCmd.Parse(args)

	// Validate input
	if *amount <= 0 || *description == "" || *category == "" {
		fmt.Println("Please provide a valid amount, description, and category.")
		os.Exit(1)
	}

	// Read and update expenses
	datasource := services.NewFileDatasource("data", "expenses.csv")
	expenses := datasource.ReadFile()
	expense := models.Expense{
		ID:          len(expenses) + 1,
		Amount:      *amount,
		Description: *description,
		Category:    *category,
	}
	expenses = append(expenses, expense)

	// Write updated expenses to file
	err := datasource.WriteFile(expenses)
	if err != nil {
		fmt.Printf("Error adding expense: %v\n", err)
		os.Exit(1)
	}

	// Output the parsed values
	fmt.Printf("Added Expense:\n")
	fmt.Printf("  Amount: $%.2f\n", *amount)
	fmt.Printf("  Description: %s\n", *description)
	fmt.Printf("  Category: %s\n", *category)
}
