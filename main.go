package main

import (
	"encoding/csv"
	"fmt"
	"os"
)
import "time"

func main() {
	fmt.Println(welcomeMessage())

	var authors []author
	csvfile, _ := os.Open("resources/authors.csv")
	r := csv.NewReader(csvfile)
	r.Comma = ';'

	csvData, err := r.ReadAll()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	for _, column := range csvData {
		authors = append(authors, author{email: column[0], firstName: column[1], lastName: column[2]})
	}

	fmt.Println(authors)
}

type author struct {
	email string
	firstName string
	lastName string
}

type book struct {
	title string
	isbn string
	authors []string
	description string
}

type magazine struct {
	title string
	isbn string
	authors []string
	publishedAt time.Time
}

func welcomeMessage() string {
	return "Hi"
}
