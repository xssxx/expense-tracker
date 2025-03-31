package models

import "time"

type Expense struct {
	ID          int
	Amount      float64
	Description string
	Category    string
	CreatedAt   time.Time
}
