package main

import (
	"database/sql"
	"html/template"
	"log"
	"net/http"
	"os"
	"path"

	"github.com/blackshirt/learn-golang/cleanarch"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	// example using html/template features
	fp := path.Join("static", "index.html")
	tmpl, err := template.ParseFiles(fp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	// you can pass data in template.Execute(w, data)
	if tmpl.Execute(w, nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}

func main() {
	// create mux Router
	router := mux.NewRouter().StrictSlash(true)

	// This is for static dir, for examples, js and css or other related files
	staticFileDir := http.Dir("./static/")
	// serve static file over /static/ path
	staticFileHandler := http.StripPrefix("/static/", http.FileServer(staticFileDir))
	router.PathPrefix("/static/").Handler(staticFileHandler).Methods("GET")

	// Get the index handler using static resources
	router.HandleFunc("/", IndexHandler).Methods("GET")

	// create subrouter for 'trains' path
	trRouter := router.PathPrefix("/trains").Subrouter()

	// open database connections to mysql server. maybe its should be moved to infrastructures layer.
	db, err := sql.Open("mysql", "test:123@tcp(127.0.0.1:3306)/test")
	if err != nil {
		log.Fatal(err.Error())
	}
	// lebih bijaksana selalu panggil Close
	defer db.Close()
	// cek jika aktif dengan Ping
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
	trRouter.HandleFunc("/", handler.Posts).Methods("POST")
	trRouter.HandleFunc("/{id:[0-9]+}", handler.GetById).Methods("GET")

	// its ready to start http services
	http.ListenAndServe(":5000", handlers.LoggingHandler(os.Stdout, router))
}
