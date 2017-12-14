package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"

	"log"
)

func main() {
	// sql.Open does not open connection, you should check it through db.Ping
	db, err := sql.Open("mysql", "dbq:123@tcp(127.0.0.1:3306)/dbq")
	if err != nil {
		log.Fatal(err)
	}
	// check if we can work
	pingerror := db.Ping()
	if pingerror != nil {
		log.Fatal(pingerror)
	}
	fmt.Println("Connected")
	defer db.Close() // its wise to do closing
	QueryforTables(db)
	QueryforSingleResult(db)
}

// some query returns result set of sql.Rows
func QueryforTables(db *sql.DB) {
	// placeholder val to store result
	var res string
	// Query results rows set (sql.Rows), if you want only single row result,
	// use QueryRow
	rows, err := db.Query("show tables")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	// iterate over rows
	for rows.Next() {
		// read the column in each rows to var with rows.Scan
		err := rows.Scan(&res)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(res) // display the results
	}
	// its good to do check the error after we're done iterating over the rows
	rowerr := rows.Err()
	if err != nil {
		log.Fatal(rowerr)
	}

}

// function returns single row result
func QueryforSingleResult(db *sql.DB) {
	var name string
	// Query for single row, contains the string
	err := db.QueryRow("select nama from pns_aktif where nama like ?", "%FATKHUL MUS%").Scan(&name)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(name, "terdaftar")
}
