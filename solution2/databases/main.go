package databases

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"os"
	"solution2/omdb-movie-search/models"

	_ "github.com/mattn/go-sqlite3"
)

func InitDB() {
	var db *sql.DB

	if _, err := os.Stat("./logcall.db"); errors.Is(err, os.ErrNotExist) {
		log.Println("Creating logcall.db")
		file, err := os.Create("./logcall.db") // Create SQLite db file
		if err != nil {
			log.Fatal(err.Error())
		}
		file.Close()
		log.Println("logcall.db created")

		db, err = sql.Open("sqlite3", "./logcall.db")
		if err != nil {
			log.Panic(fmt.Printf("Error connecting to database: error = %v", err))
		}
		defer db.Close()

		createTable(db)
	}
}

func GetDB() *sql.DB {
	db, err := sql.Open("sqlite3", "./logcall.db")
	if err != nil {
		log.Panic(fmt.Printf("Error connecting to database: error = %v", err))
	}

	return db
}

func createTable(db *sql.DB) {
	createLogCallTable := `CREATE TABLE logcall (
		"id" integer NOT NULL PRIMARY KEY AUTOINCREMENT,		
		"call_type" TEXT,
		"param_query" TEXT,
		"status" integer,		
		"created_at" DATETIME		
	  );` // SQL Statement for Create Table

	log.Println("Create logcall table...")
	statement, err := db.Prepare(createLogCallTable) // Prepare SQL Statement
	if err != nil {
		log.Fatal(err.Error())
	}
	statement.Exec() // Execute SQL Statements
	defer statement.Close()

	log.Println("logcall table created")
}

func AddLogCall(db *sql.DB, newLogcall models.LogCall) {
	statement, _ := db.Prepare("INSERT INTO logcall (id, call_type, param_query, status, created_at) VALUES (?, ?, ?, ?, ?)")
	statement.Exec(nil, newLogcall.CallType, newLogcall.ParamQuery, newLogcall.Status, newLogcall.CreatedAt)
	defer statement.Close()
}
