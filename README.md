# Budgie Expense Tracker
A simple expense tracker built using Golang, React JS and MySQL. The application supports user registration and authentication, CRUD operations for expenses, date range searching and pagination.

## Features
* Registration and Authentication
* CRUD (Create, Read, Update and Delete) expenses.
* Date range search with 7 or 30 day presets.
* Pagination
* Expense category pie chart

![](BudgieGif1.gif)
![](BudgieGif2.gif)

## Project Structure
     .
    ├── budgie-api              # Backend Golang API Folder
    │   ├── controllers         # Endpoint handlers     
    │   ├── database            # Database connection library
    |   ├── middlewares         # Middlewares
    │   ├── models              # Custom structs
    │   ├── repo                # Database calls/methods
    │   ├── routes              # API routes
    |   ├── static/mysql        # MySQL DB schema
    |   ├── util                # Utility functions
    │   ├── .env                # Environment variables
    │   ├── go.mod
    |   ├── go.sum           
    │   └── main.go             # Go API entry point
    |      
    ├── budgie-frontend         # Frontend React app
    │   ├── public
    │   ├── src                 # Components and Pages
    │   ├── package.json
    │   └── package-lock.json
    └── README.md    

## Installation

### Clone Repo

```sh
git clone https://github.com/budgie-expense-tracker
```
   
### Initialize MySQL Database

* Load the database schema located in budgie-api/static/mysql/budgieDB.sql into MySQL.
* Start your MySQL server and update environment variables located in budgie-api/.env file with your MySQL username, password, host and port.

```sh
## HTTP Server Port ##
SERVER_PORT=":8080"

## MySQL Config ##
DB_USER="<DB_USER>" 
DB_PASSWORD="<DB_PASS>"
DB_HOST="<DB_HOST>"
DB_PORT="<DB_PORT>"
DB_NAME="budgieDB"

## JWT Secret
JWT_SECRET="secret"
```

### Start Go API Server

Ensure Golang version 1.16.2 or newer is installed.

Change directory to the go API folder
```sh
cd budgie-expense-tracker/budgie-api
```

Download the Go API dependancies (list of dependancies are located in budgie-expense-tracker/budgie-api/go.mod)
```sh
go get
```

Start the Go server
```sh
go run .
```

The Go server will start on localhost:8080 by default (port can be changed in the .env file)

### Start React Server
Download the front end react dependancies and start the server.

Change directory to the react-auth folder
```sh
cd budgie-expense-tracker/budgie-frontend
```

Download react app dependancies
```sh
npm install
```

Start react server
```sh
npm start
```

The react server will start on localhost:3000
