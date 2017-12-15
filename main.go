package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/blackshirt/learn-golang/cleanarch"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	// create mux Router
	router := mux.NewRouter()
	// create subrouter for 'trains' path
	trRouter := router.PathPrefix("/trains").Subrouter()

	// open database connections to mysql server. maybe its should be moved to infrastructures layer.
	db, err := sql.Open("mysql", "test:123@tcp(127.0.0.1:3306)/test")
	if err != nil {
		log.Fatal(err.Error())
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		log.Fatal(err.Error())
	}

	// Setup new SQL db handler from opened connection.
	handle := cleanarch.NewSQLHandler(db)
	// create use case
	truc := cleanarch.NewTrainUCase(handle)
	// create new http services handler
	handler := cleanarch.NewHTTPTrainHandler(truc)

	// Set route path
	trRouter.HandleFunc("/", handler.Fetch).Methods("GET")
	trRouter.HandleFunc("/{id:[0-9]+}", handler.GetById).Methods("GET")

	// its ready to start http services
	http.ListenAndServe(":8000", handlers.LoggingHandler(os.Stdout, router))
}
