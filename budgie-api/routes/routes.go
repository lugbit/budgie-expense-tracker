package routes

import (
	"github.com/gorilla/mux"
	"github.com/lugbit/budgie-expense-tracker/controllers"
	"github.com/lugbit/budgie-expense-tracker/middlewares"
)

// defined routes and the route handler
func GetRoutes(r *mux.Router) {
	r.HandleFunc("/api/register", controllers.RegisterController).Methods("POST")
	r.HandleFunc("/api/authenticate", controllers.AuthController).Methods("POST")
	r.HandleFunc("/api/user", middlewares.ValidateToken(controllers.UserController)).Methods("GET") // protected
	r.HandleFunc("/api/logout", controllers.LogoutController).Methods("POST")
	r.HandleFunc("/api/expense", middlewares.ValidateToken(controllers.AddExpenseController)).Methods("POST")                         // protected
	r.HandleFunc("/api/expense/{id}", middlewares.ValidateToken(controllers.DeleteExpenseController)).Methods("DELETE")               // protected
	r.HandleFunc("/api/expense/{id}", middlewares.ValidateToken(controllers.UpdateExpenseController)).Methods("PUT")                  // protected
	r.HandleFunc("/api/expense/{id}", middlewares.ValidateToken(controllers.GetExpenseController)).Methods("GET")                     // protected
	r.HandleFunc("/api/expenses", middlewares.ValidateToken(controllers.GetExpensesController)).Methods("POST")                       // protected
	r.HandleFunc("/api/expenses/count-by-category", middlewares.ValidateToken(controllers.GetExpenseCountByCategory)).Methods("POST") // protected
	r.HandleFunc("/api/expense-categories", controllers.GetExpenseCategoryController).Methods("GET")
}
