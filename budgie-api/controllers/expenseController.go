package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/lugbit/budgie-expense-tracker/models"
	"github.com/lugbit/budgie-expense-tracker/repo"
	"github.com/lugbit/budgie-expense-tracker/util"
)

// add expense
func AddExpenseController(w http.ResponseWriter, r *http.Request) {
	// get userID from validate token middleware context
	userID, _ := r.Context().Value("userID").(int)

	errors := []models.Error{}

	// new expense object
	newExpense := models.Expense{}
	json.NewDecoder(r.Body).Decode(&newExpense)

	// check for empty fields
	if newExpense.Title == "" {
		errors = append(errors, util.NewError("", "Invalid Attribute", "Expense title is required"))
	}
	if newExpense.CategoryID == "" {
		errors = append(errors, util.NewError("", "Invalid Attribute", "Expense category ID is required"))
	}
	if newExpense.Date == "" {
		errors = append(errors, util.NewError("", "Invalid Attribute", "Expense date is required"))
	}
	if newExpense.Amount == "" {
		errors = append(errors, util.NewError("", "Invalid Attribute", "Expense amount is required"))
	}

	if len(errors) > 0 {
		// send errors as response
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)

		json.NewEncoder(w).Encode(models.Errors{Errors: errors})
		return
	}

	newExpense.UserID = userID

	// insert expense into db
	newExpenseID, err := repo.AddExpense(newExpense)
	if err != nil {
		log.Fatalln(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(map[string]int{"expenseID: ": newExpenseID})
}

// delete expense
func DeleteExpenseController(w http.ResponseWriter, r *http.Request) {
	// get userID from validate token middleware context
	userID, _ := r.Context().Value("userID").(int)

	// get ID param
	params := mux.Vars(r)
	expenseIDString := params["id"]
	expenseID, _ := strconv.Atoi(expenseIDString)

	_, err := repo.DeleteExpense(expenseID, userID)
	if err != nil {
		// no rows affected (nothing was deleted)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(400)

		json.NewEncoder(w).Encode(map[string]string{"Error": "Nothing was deleted"})
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(map[string]int{"expenseID deleted: ": expenseID})
}

// get expense
func GetExpenseController(w http.ResponseWriter, r *http.Request) {
	// get userID from validate token middleware context
	userID, _ := r.Context().Value("userID").(int)

	// get ID param
	params := mux.Vars(r)
	expenseIDString := params["id"]
	expenseID, _ := strconv.Atoi(expenseIDString)

	expense, err := repo.GetExpense(expenseID, userID)
	if err != nil {
		// no expense found with that ID
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(400)

		json.NewEncoder(w).Encode(map[string]string{"Error": "Expense does not exist"})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(expense)
}

// get expenses
func GetExpensesController(w http.ResponseWriter, r *http.Request) {
	// get userID from validate token middleware context
	userID, _ := r.Context().Value("userID").(int)

	dateRange := models.DateRangeLimit{}

	json.NewDecoder(r.Body).Decode(&dateRange)

	startDateString := dateRange.StartDate
	endDateString := dateRange.EndDate
	limit := dateRange.Limit
	offset := dateRange.Offset

	// get expenses count and total
	expensesCount, _ := repo.GetExpensesCount(startDateString, endDateString, userID)
	var expensesTotal float64
	if expensesCount > 0 {
		expensesTotal, _ = repo.GetExpensesTotal(startDateString, endDateString, userID)
	}

	// get all user expense
	expenses, _ := repo.GetExpensesWithRange(userID, startDateString, endDateString, limit, offset)

	expensesPayload := models.ExpensesPayload{Count: expensesCount, Total: expensesTotal, Expenses: expenses}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(expensesPayload)
}

// get expense category
func GetExpenseCategoryController(w http.ResponseWriter, r *http.Request) {
	// get all user expense
	expenseCategories, err := repo.GetExpenseCategories()
	if err != nil {
		log.Fatalln(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(expenseCategories)
}

// get expense count by category
func GetExpenseCountByCategory(w http.ResponseWriter, r *http.Request) {
	// get userID from validate token middleware context
	userID, _ := r.Context().Value("userID").(int)

	dateRange := models.DateRangeLimit{}

	json.NewDecoder(r.Body).Decode(&dateRange)

	startDateString := dateRange.StartDate
	endDateString := dateRange.EndDate

	// get all user expense
	expenses, _ := repo.GetExpenseCountByCategory(userID, startDateString, endDateString)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(expenses)
}

// update expense
func UpdateExpenseController(w http.ResponseWriter, r *http.Request) {
	// get userID from validate token middleware context
	userID, _ := r.Context().Value("userID").(int)

	// get ID param
	params := mux.Vars(r)
	expenseIDString := params["id"]
	expenseID, _ := strconv.Atoi(expenseIDString)

	// should check if ID is legit here

	// destructure JSON response
	updatedExpense := models.Expense{}
	json.NewDecoder(r.Body).Decode(&updatedExpense)

	_, err := repo.UpdateExpense(updatedExpense, expenseID, userID)
	if err != nil {
		// update failed
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(400)

		json.NewEncoder(w).Encode(map[string]string{"Error": "Update failed"})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(updatedExpense) // return updatedExpense
}
