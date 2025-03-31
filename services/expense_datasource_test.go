package services

import (
	"os"
	"testing"

	"github.com/xssxx/expense-tracker/models"
)

func TestFileDataSource_NegativeAmount(t *testing.T) {
	// Setup: Create a temporary file for testing
	tempFilePath := "test_expenses_negative.csv"
	defer os.Remove(tempFilePath) // Clean up after the test

	dataSource := NewFileDatasource(".", tempFilePath)

	// Test data with a negative amount
	expenses := []models.Expense{
		{ID: 1, Amount: -50.00, Description: "Refund", Category: "Misc"},
	}

	// Test WriteFile
	err := dataSource.WriteFile(expenses)
	if err != nil {
		t.Fatalf("WriteFile failed: %v", err)
	}

	// Test ReadFile
	readExpenses := dataSource.ReadFile()
	if len(readExpenses) != 1 {
		t.Fatalf("Expected 1 expense, got %d", len(readExpenses))
	}

	// Verify the content
	if readExpenses[0].Amount != expenses[0].Amount {
		t.Errorf("Expected amount %.2f, got %.2f", expenses[0].Amount, readExpenses[0].Amount)
	}
}

func TestFileDataSource_DeleteNonExistentExpense(t *testing.T) {
	// Setup: Create a temporary file for testing
	tempFilePath := "test_expenses_delete.csv"
	defer os.Remove(tempFilePath) // Clean up after the test

	// Create a FileDataSource instance
	dataSource := NewFileDatasource(".", tempFilePath)

	// Test data
	expenses := []models.Expense{
		{ID: 1, Amount: 100.00, Description: "Groceries", Category: "Food"},
		{ID: 2, Amount: 200.00, Description: "Transport", Category: "Travel"},
	}

	// Write initial data to the file
	err := dataSource.WriteFile(expenses)
	if err != nil {
		t.Fatalf("WriteFile failed: %v", err)
	}

	// Attempt to delete a non-existent expense
	nonExistentID := 99
	updatedExpenses := make([]models.Expense, 0)
	found := false
	for _, expense := range expenses {
		if expense.ID == nonExistentID {
			found = true
			continue
		}
		updatedExpenses = append(updatedExpenses, expense)
	}

	// Verify that the expense was not found
	if found {
		t.Fatalf("Unexpectedly found expense with ID %d", nonExistentID)
	}

	// Write the updated list back to the file
	err = dataSource.WriteFile(updatedExpenses)
	if err != nil {
		t.Fatalf("WriteFile failed: %v", err)
	}

	// Read the file and verify the content is unchanged
	readExpenses := dataSource.ReadFile()
	if len(readExpenses) != len(expenses) {
		t.Fatalf("Expected %d expenses, got %d", len(expenses), len(readExpenses))
	}

	for i, expense := range readExpenses {
		if expense.ID != expenses[i].ID ||
			expense.Amount != expenses[i].Amount ||
			expense.Description != expenses[i].Description ||
			expense.Category != expenses[i].Category {
			t.Errorf("Mismatch at index %d: expected %+v, got %+v", i, expenses[i], expense)
		}
	}
}
