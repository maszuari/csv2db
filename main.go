package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func processCsv(filePath string) {

	f, err := os.Open(filePath)

	if err != nil {
		log.Fatal("Unable to open CSV file", err)
	}
	defer f.Close()

	r := csv.NewReader(f)
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
	}
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
