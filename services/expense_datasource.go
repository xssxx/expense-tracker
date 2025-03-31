package services

import (
	"bufio"
	"os"
	"strconv"
	"strings"

	"github.com/xssxx/expense-tracker/models"
)

type FileDataSource struct {
	FilePath string
	FileName string
}

func NewFileDatasource(filePath, fileName string) *FileDataSource {
	return &FileDataSource{
		FilePath: filePath,
		FileName: fileName,
	}
}

func (f *FileDataSource) ReadFile() []models.Expense {
	file, err := os.Open(f.FilePath + "/" + f.FileName)
	if err != nil {
		panic(err)
	}

	defer file.Close()

	expenses := make([]models.Expense, 0)
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ",")
		if len(parts) < 4 {
			continue
		}

		id, err := strconv.Atoi(parts[0])
		if err != nil {
			continue
		}

		amount, err := strconv.ParseFloat(parts[1], 64)
		if err != nil {
			continue
		}

		expense := models.Expense{
			ID:          id,
			Amount:      amount,
			Description: parts[2],
			Category:    parts[3],
		}

		expenses = append(expenses, expense)
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return expenses
}

func (f *FileDataSource) WriteFile(expenses []models.Expense) error {
	file, err := os.OpenFile(f.FilePath+"/"+f.FileName, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)

	for _, expense := range expenses {
		line := strconv.Itoa(expense.ID) + "," +
			strconv.FormatFloat(expense.Amount, 'f', 2, 64) + "," +
			expense.Description + "," +
			expense.Category + "\n"
		_, err := writer.WriteString(line)
		if err != nil {
			return err
		}
	}

	return writer.Flush()
}
