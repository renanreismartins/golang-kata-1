package main

import (
	"testing"
)

func TestAbs(t *testing.T) {
	got := welcomeMessage()
	if got != "Hi" {
		t.Errorf("welcomeMessage() = %s; want Hi", got)
	}
}

func TestReadAuthors(t *testing.T) {
	authors := readAuthors("resources/authors_test.csv")
	var author = authors[0]

	if len(authors) != 1 {
		t.Errorf("readAuthors is not removing the first line of the file (header)")
	}

	if author.firstName != "Paul" {
		t.Errorf("Read firstName incorretly")
	}

	if author.lastName != "Walter" {
		t.Errorf("Read lastName incorretly")
	}

	if author.email != "null-walter@echocat.org" {
		t.Errorf("Read email incorretly")
	}
}