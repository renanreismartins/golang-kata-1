package main

import (
	"encoding/csv"
	"fmt"
	"os"
)
import "time"

func main() {
	fmt.Println(welcomeMessage())
}

func readAuthors(path string) []author {
	file, _ := os.Open(path)
	reader := csv.NewReader(file)
	reader.Comma = ';'

	csvData, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Error reading file")
		fmt.Println(err) // TODO: Investigate if reading a file that does not exist ReadAll return an error 'invalid argument', no stacktrace; Research about better error handling
		os.Exit(1) // TODO: Does not sound good exiting the process; maybe return an empty array?!
	}

	var authors []author
	for _, column := range csvData {
		authors = append(authors, author{email: column[0], firstName: column[1], lastName: column[2]})
	}

	return authors[1:] // Remove first line; CSV header
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
