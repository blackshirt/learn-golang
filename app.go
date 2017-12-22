// app.go
package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path"

	"github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	// sample my own package
	"github.com/blackshirt/learn-golang/token"
)

var database *sql.DB
var err error
var myKey = []byte("rahasia")

func main() {

	db, err := sql.Open("mysql", "dbq:123@tcp(127.0.0.1:3306)/dbq")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		panic(err)
	}

	database = db
	// StrictSlash berpengaruh pada default behaviour untuk trailing slash.
	// defaultnya false. Saat true, jika route path adalah "/path/", maka mengakses
	// "/path" akan redirect ke "/path/" begitu pula sebaliknya.
	// jika false, "/path/" dan "/path" adalah dua route yang tidak match.
	// Pengecualian jika route diset path prefix dengan PathPrefix(), strict slash
	// akan diignore untuk route tersebut.
	router := mux.NewRouter().StrictSlash(true)

	// This is for static dir, for examples, js and css or ther related files
	staticFileDir := http.Dir("./static/")
	// serve static file over /static/ path
	staticFileHandler := http.StripPrefix("/static/", http.FileServer(staticFileDir))
	router.PathPrefix("/static/").Handler(staticFileHandler).Methods("GET")

	// Get the index handler using static resources
	router.HandleFunc("/", IndexHandler).Methods("GET")
	router.HandleFunc("/login", LoginHandler).Methods("POST")
	router.HandleFunc("/token", TokenHandler).Methods("GET")
	router.HandleFunc("/trainings", TrainingsHandler).Methods("GET")
	router.HandleFunc("/protected", ProtectedHandler).Methods("GET")
	http.ListenAndServe(":8080", handlers.LoggingHandler(os.Stdout, router))
}

type training struct {
	// elemen pakai CapitalCase, karena kalau private, tidak dikenal
	// untuk menghandle NULL data dari database, gunakan related NullType dari sql
	// karena rows.Scan akan error, sql: Scan error on column index 2:
	// unsupported Scan, storing driver.Value type <nil> into type *string

	Id            int            `json:"id"`
	Penyelenggara sql.NullString `json:"penyelenggara"`
	Nama_diklat   string         `json:"nama_diklat"`
	Thn_diklat    string         `json:"thn_diklat"`
}

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

func TrainingsHandler(w http.ResponseWriter, r *http.Request) {
	stmt := "select id, penyelenggara, nama_diklat, thn_diklat from riwayat_diklat limit 20"
	rows, err := database.Query(stmt)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	trainings := []training{}
	for rows.Next() {
		var res training
		err := rows.Scan(&res.Id, &res.Penyelenggara, &res.Nama_diklat, &res.Thn_diklat)
		if err != nil {
			log.Fatal(err)
		}
		trainings = append(trainings, res)
	}
	// set content-type header to application/json
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	// writes result to response writer
	json.NewEncoder(w).Encode(trainings)
}

func TokenHandler(w http.ResponseWriter, r *http.Request) {
	/*
		// create new token
		token := jwt.New(jwt.SigningMethodHS256)
		// create map claims
		claims := token.Claims.(jwt.MapClaims)

		// Set token claims
		claims["admin"] = true
		claims["name"] = "black hajime"
		claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

		// Sign the token with secret
		tokenString, _ := token.SignedString(myKey)

		// Finally, write the token to the response writer
		w.Write([]byte(tokenString))
	*/

	// test using own package, token
	var username string = "test user"
	tokenString, err := token.GenerateToken(username)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	}
	w.Write([]byte(tokenString))

}

func ProtectedHandler(w http.ResponseWriter, r *http.Request) {
	// extract jwt token, from query arguments, with key "token"
	extractor := request.ArgumentExtractor{"token"}
	keyFunc := func(token *jwt.Token) (interface{}, error) {
		// check signing method
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		publickeybytes, err := ioutil.ReadFile("keys/public.key")
		if err != nil {
			return nil, fmt.Errorf("could not open jwt key, %v", err)
		}
		// jwt.ParseRSAPublicKeyFromPEM return *rsa.PrivateKey struct, represents an RSA key
		pubkey, err := jwt.ParseRSAPublicKeyFromPEM(publickeybytes)
		if err != nil {
			fmt.Println("error parse public key:", err)
		}
		return pubkey, nil
	}
	token, err := request.ParseFromRequest(r, extractor, keyFunc)

	if err != nil {
		//http.Error(w, err.Error(), http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		// no token present in request, or
		// or token contains an invalid number of segments
		// illegal base64 data at input byte 1
		return
	}
	json.NewEncoder(w).Encode(token)
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.Write([]byte("gak boleh"))
	}
	w.Write([]byte(http.StatusText(http.StatusOK)))
}
