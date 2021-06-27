package repo

import (
	"database/sql"
	"errors"
	"log"
	"strconv"

	"github.com/lugbit/budgie-expense-tracker/database"
	"github.com/lugbit/budgie-expense-tracker/models"
)

// add a new expense
func AddExpense(expense models.Expense) (int, error) {
	// sql query
	query := `
			INSERT INTO tblExpense(fldTitle, fldFKCategoryID, fldDate, fldAmount, fldNote, fldFKUserID)
			VALUES(?, ?, ?, ?, ?, ?);
			`

	// prepare statement
	stmt, err := database.DB.Prepare(query)
	if err != nil {
		log.Fatalln(err)
	}
	defer stmt.Close()

	// execute prepared statement and return result object
	result, err := stmt.Exec(expense.Title, expense.CategoryID, expense.Date, expense.Amount, expense.Note, expense.UserID)
	if err != nil {
		log.Fatalln(err)
	}

	// check no. of rows affected
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Fatalln(err)
	}
	// if no rows are affected, return with -1 and the error
	if rowsAffected == 0 {
		return -1, errors.New("no rows affected")
	}

	// retrieve last insert ID from the insert
	lastInsertID, err := result.LastInsertId()
	if err != nil {
		log.Fatalln(err)
	}

	return int(lastInsertID), nil
}

// for pagination purposes. Retrieves the total expenses for a user
func GetExpensesCount(startDate string, endDate string, userID int) (int, error) {
	var countStr string

	query := `
			SELECT 
				COUNT(*)
			FROM 
				tblExpense
			WHERE (tblExpense.fldDate BETWEEN ? AND ?)
			AND tblExpense.fldFKUserID = ?
			`

	stmt, err := database.DB.Prepare(query)
	if err != nil {
		log.Fatalln(err)
	}
	defer stmt.Close()

	err = stmt.QueryRow(startDate, endDate, userID).Scan(&countStr)
	if err != nil {
		if err == sql.ErrNoRows {
			return 0, nil
		}
		log.Fatalln(err)
	}
	count, _ := strconv.Atoi(countStr)
	return count, nil
}

// get expenses total
func GetExpensesTotal(startDate string, endDate string, userID int) (float64, error) {
	var totalStr string

	query := `
			SELECT 
				SUM(tblExpense.fldAmount) 
			FROM 
				tblExpense
			WHERE (tblExpense.fldDate BETWEEN ? AND ?)
			AND tblExpense.fldFKUserID = ?
			`

	stmt, err := database.DB.Prepare(query)
	if err != nil {
		log.Fatalln(err)
	}
	defer stmt.Close()

	err = stmt.QueryRow(startDate, endDate, userID).Scan(&totalStr)
	if err != nil {
		if err == sql.ErrNoRows {
			return -1, nil
		}
		log.Fatalln(err)
	}
	if totalStr == "" {
		totalStr = "0"
	}

	total, _ := strconv.ParseFloat(totalStr, 8)
	return total, nil
}

// get all expense from a user within the date range. This also accepts a limit and offset for pagination
func GetExpensesWithRange(userID int, startDate string, endDate string, limit int, offset int) ([]models.Expense, error) {
	expense := models.Expense{}
	expenses := []models.Expense{}

	query := `
			SELECT
				tblExpense.fldID, 
				tblExpense.fldTitle,
				tblExpenseCat.fldCategoryName,
				tblExpense.fldDate, 
				tblExpense.fldAmount, 
				tblExpense.fldNote
			FROM 
				tblExpense
			INNER JOIN tblExpenseCat
				ON tblExpense.fldFKCategoryID = tblExpenseCat.fldID
			WHERE (tblExpense.fldDate BETWEEN ? AND ?)
			AND tblExpense.fldFKUserID = ?
			ORDER BY tblExpense.fldDate DESC
			LIMIT ?
			OFFSET ?;
			`

	stmt, err := database.DB.Prepare(query)
	if err != nil {
		log.Fatalln(err)
	}
	defer stmt.Close()

	rows, err := stmt.Query(startDate, endDate, userID, limit, offset)
	if err != nil {
		if err == sql.ErrNoRows {
			return expenses, nil
		}
	}
	defer rows.Close()

	rowCount := 0

	for rows.Next() {
		rowCount++
		err := rows.Scan(&expense.ID, &expense.Title, &expense.CategoryName, &expense.Date, &expense.Amount, &expense.Note)
		if err != nil {
			log.Fatalln(err)
		}

		expenses = append(expenses, expense)
	}

	if rowCount == 0 {
		return []models.Expense{}, sql.ErrNoRows
	}

	err = rows.Err()
	if err != nil {
		log.Fatalln(err)
	}

	return expenses, nil
}

// get expense count by category
func GetExpenseCountByCategory(userID int, startDate string, endDate string) ([]models.ExpenseCategoryCount, error) {
	expenseCatCount := models.ExpenseCategoryCount{}
	expenseCatCounts := []models.ExpenseCategoryCount{}

	query := `
			SELECT
				tblExpenseCat.fldCategoryName, 
				COUNT(*)
			FROM 
				tblExpense
			INNER JOIN tblExpenseCat
			ON tblExpense.fldFKCategoryID = tblExpenseCat.fldID
			WHERE (tblExpense.fldDate BETWEEN ? AND ?)
			AND tblExpense.fldFKUserID = ?
			GROUP BY tblExpenseCat.fldCategoryName;
			`

	stmt, err := database.DB.Prepare(query)
	if err != nil {
		log.Fatalln(err)
	}
	defer stmt.Close()

	rows, err := stmt.Query(startDate, endDate, userID)
	if err != nil {
		log.Fatalln(err)
	}
	defer rows.Close()

	rowCount := 0

	for rows.Next() {
		rowCount++
		err := rows.Scan(&expenseCatCount.CategoryName, &expenseCatCount.CategoryCount)
		if err != nil {
			log.Fatalln(err)
		}

		expenseCatCounts = append(expenseCatCounts, expenseCatCount)
	}

	if rowCount == 0 {
		return []models.ExpenseCategoryCount{}, sql.ErrNoRows
	}

	err = rows.Err()
	if err != nil {
		log.Fatalln(err)
	}

	return expenseCatCounts, nil
}

// get expense categories
func GetExpenseCategories() ([]models.ExpenseCategory, error) {
	expenseCategory := models.ExpenseCategory{}     // individual expense category
	expenseCategories := []models.ExpenseCategory{} // slice of expense catagorys

	query := `
			SELECT
				fldID, 
				fldCategoryName
			FROM 
				tblExpenseCat;
			`

	stmt, err := database.DB.Prepare(query)
	if err != nil {
		log.Fatalln(err)
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		log.Fatalln(err)
	}
	defer rows.Close()

	rowCount := 0

	for rows.Next() {
		rowCount++
		err := rows.Scan(&expenseCategory.ID, &expenseCategory.CategoryName)
		if err != nil {
			log.Fatalln(err)
		}

		expenseCategories = append(expenseCategories, expenseCategory)
	}

	if rowCount == 0 {
		return []models.ExpenseCategory{}, errors.New("no rows found")
	}

	err = rows.Err()
	if err != nil {
		log.Fatalln(err)
	}

	return expenseCategories, nil

}

// get an expense
func GetExpense(expenseID, userID int) (models.Expense, error) {
	expense := models.Expense{}

	query := `
			SELECT
				fldID, 
				fldTitle, 
				fldDate, 
        		fldAmount, 
        		fldNote, 
        		fldFKCategoryID
			FROM tblExpense
				WHERE fldID = ?
				AND fldFKUserID = ?;
			`

	stmt, err := database.DB.Prepare(query)
	if err != nil {
		log.Fatalln(err)
	}
	defer stmt.Close()

	err = stmt.QueryRow(expenseID, userID).Scan(&expense.ID, &expense.Title, &expense.Date, &expense.Amount, &expense.Note, &expense.CategoryID)
	if err != nil {
		if err == sql.ErrNoRows {
			return models.Expense{}, errors.New("row not found")
		}
		log.Fatalln(err)
	}

	return expense, nil
}

// delete an expense
func DeleteExpense(expenseID, userID int) (int, error) {
	query := `
			DELETE FROM tblExpense
			WHERE fldID = ?
    		AND fldFKUserID = ?;
			`

	stmt, err := database.DB.Prepare(query)
	if err != nil {
		log.Fatalln(err)
	}
	defer stmt.Close()

	res, err := stmt.Exec(expenseID, userID)
	if err != nil {
		log.Fatalln(err)
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		log.Fatalln(err)
	}

	if rowsAffected == 0 {
		return -1, errors.New("no rows affected")
	}

	return int(rowsAffected), nil
}

// update an existing
func UpdateExpense(updatedNote models.Expense, expenseID, userID int) (int, error) {
	query := `
			UPDATE tblExpense
			SET 
				fldTitle = ?, 
				fldDate = ?, 
        		fldAmount = ?, 
        		fldNote = ?, 
        		fldFKCategoryID = ?
			WHERE fldID = ? 
			AND fldFKUserID = ?;
			`

	stmt, err := database.DB.Prepare(query)
	if err != nil {
		log.Fatalln(err)
	}
	defer stmt.Close()

	res, err := stmt.Exec(updatedNote.Title, updatedNote.Date, updatedNote.Amount, updatedNote.Note, updatedNote.CategoryID, expenseID, userID)
	if err != nil {
		log.Fatalln(err)
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		log.Fatalln(err)
	}

	if rowsAffected == 0 {
		return -1, errors.New("no rows affected")
	}

	return int(rowsAffected), nil
}
