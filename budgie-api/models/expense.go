package models

// expense model
type Expense struct {
	ID           string `json:"expenseID,omitempty"`
	Title        string `json:"title,omitempty"`
	CategoryID   string `json:"categoryID,omitempty"`
	CategoryName string `json:"categoryName,omitempty"`
	Date         string `json:"date,omitempty"`
	Amount       string `json:"amount,omitempty"`
	Note         string `json:"note,omitempty"`
	UserID       int    `json:"userID,omitempty"`
}

// expense payload model includes the expenses plus the count and total
type ExpensesPayload struct {
	Count    int       `json:"count,omitempty"`
	Total    float64   `json:"total,omitempty"`
	Expenses []Expense `json:"expenses,omitempty"`
}

// expense category model
type ExpenseCategory struct {
	ID           string `json:"categoryID"`
	CategoryName string `json:"categoryName"`
}

// expense category count model
type ExpenseCategoryCount struct {
	CategoryName  string `json:"categoryName,omitempty"`
	CategoryCount string `json:"categoryCount,omitempty"`
}
