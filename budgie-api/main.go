package main

import (
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/lugbit/budgie-expense-tracker/database"
	"github.com/lugbit/budgie-expense-tracker/routes"
	"github.com/subosito/gotenv"
)

// load env variables on startup
func init() {
	err := gotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	// configure CORS (whitelisted headers, methods and origins)
	headers := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"})
	origins := handlers.AllowedOrigins([]string{"http://localhost:3000"})

	// connect to database
	database.ConnectDB()
	// test database connection
	err := database.DB.Ping()
	if err != nil {
		log.Fatalln("error: unable to connect to the database.")
	}

	// create new gorilla mux router
	r := mux.NewRouter()
	// call the routes passing in the mux router instance
	routes.GetRoutes(r)

	log.Fatal(http.ListenAndServe(os.Getenv("SERVER_PORT"), handlers.CORS(headers, methods, origins, handlers.AllowCredentials())(r)))
}
