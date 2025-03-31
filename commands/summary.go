package commands

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/xssxx/expense-tracker/models"
	"github.com/xssxx/expense-tracker/services"
)

func SummaryExpense(args []string) {
	summaryCmd := flag.NewFlagSet("summary", flag.ExitOnError)
	month := summaryCmd.Int("month", 0, "Month to summarize (1-12, optional)")
	summaryCmd.Parse(args)

	datasource := services.NewFileDatasource("data", "expenses.csv")
	expenses := datasource.ReadFile()
	if len(expenses) == 0 {
		fmt.Println("No expenses found.")
		return
	}

	if *month >= 1 && *month <= 12 {
		filteredExpenses := make([]models.Expense, 0)
		for _, expense := range expenses {
			if expense.CreatedAt.Month() == time.Month(*month) {
				filteredExpenses = append(filteredExpenses, expense)
			}
		}
		expenses = filteredExpenses
	} else if *month != 0 {
		fmt.Println("Please provide a valid month (1-12).")
		os.Exit(1)
	}

	total := 0.0
	for _, expense := range expenses {
		total += expense.Amount
	}

	if *month >= 1 && *month <= 12 {
		fmt.Printf("Total Expenses for Month %d: %.2f\n", *month, total)
	} else {
		fmt.Printf("Total Expenses (All Months): %.2f\n", total)
	}
}
