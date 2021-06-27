package util

import (
	"github.com/lugbit/budgie-expense-tracker/models"
)

// create and return a new error
func NewError(code, title, message string) models.Error {
	newError := models.Error{
		Code:    code,
		Title:   title,
		Message: message,
	}
	return newError
}
