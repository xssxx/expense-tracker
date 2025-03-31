package main

import (
	"fmt"
	"os"

	"github.com/xssxx/expense-tracker/commands"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please provide a valid subcommand (e.g., 'add').")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "add":
		commands.AddExpense(os.Args[2:])
	case "list":
		commands.ListExpense()
	case "delete":
		commands.DeleteExpense(os.Args[2:])
	case "summary":
		commands.SummaryExpense()
	default:
		fmt.Println("Invalid command. Use 'add' to add an expense.")
		os.Exit(1)
	}
}
