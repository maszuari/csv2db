package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
)

type MyDb struct {
	*sqlx.DB
}

func processCsv(filePath string) {

	f, err := os.Open(filePath)

	if err != nil {
		log.Fatal("Unable to open CSV file", err)
	}
	defer f.Close()

	r := csv.NewReader(f)
	db := connectToDb()
	myDb := MyDb{db}
	for {

		row, err := r.Read()
		// Stop at EOF.
		if err == io.EOF {
			break
		}

		if err != nil {
			panic(err)
		}

		fmt.Println(row[0])
		myDb.saveRow(row[0])
	}
}

func connectToDb() *sqlx.DB {
	//TODO Set username, password and db name in .env
	db, err := sqlx.Connect("mysql", "dbuser:dbpassword@(localhost:3306)/dbname")
	if err != nil {
		log.Fatalln(err)
	}

	return db
}

func (db *MyDb) saveRow(row string) {

	tx := db.MustBegin()
	db.MustExec("INSERT INTO test_one (code) VALUES (?)", row)
	tx.Commit()

}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	filePath := os.Getenv("file_path")
	fmt.Println(filePath)
	processCsv(filePath)
}
