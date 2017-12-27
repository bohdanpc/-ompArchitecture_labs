package main

import (
	"testing"
	"../book"
    "../tree"
)

const findLoopCount = 15000000

var mytreeBalanced = new(tree.Tree)

func TestEquals(t *testing.T) {
	book1 := book.Book{"Robinsone Crusoe", "Daniel Defoe"}
	book2 := book.Book{"Robinsone Crusoe", "Daniel Defoe"}
	if !book1.Equals(book2) {
		t.Errorf("Expection equality of two books")
	}
}

func TestLess(t *testing.T) {
	book1 := book.Book{"Nine Princes in Amber", "Roger Zelyazny"}
	book2 := book.Book{"Robinsone Crusoe", "Daniel Defoe"}
	if !book1.Less(book2) {
		t.Errorf("Expected book1 < book2")
	}
}

func TestGetName(t *testing.T) {
	const author = "Daniel Defoe"
	book := book.Book{"", author}
	if book.GetAuthor() != author {
		t.Errorf("Name is incorrect, got: %s, want: %s.", book.GetAuthor(), author)
	}
}

func TestGetAuthor(t *testing.T) {
	const name = "Robinson Crusoe"
	book := book.Book{name, ""}
	if book.GetName() != name {
		t.Errorf("Name is incorrect, got: %s, want: %s.", book.GetAuthor(), name)
	}
}


func TestAddRecBalance(t *testing.T) {
	mytreeBalanced = new(tree.Tree)

	book.AddFromCsvFileRecBalance(mytreeBalanced, "/home/thereptile/" +
		"GoglandProjects/CompArchitecture_labs/lab1_sem2/books.csv")

	mytreeBalanced.AddRecBalance(book.Book{"MANKIND REVOLUTION", "CRUEL MAMBY"})
	}



func TestFindBalance(t *testing.T) {
	book1 := book.Book{"MANKIND REVOLUTION", "CRUEL MAMBY"}
	for i := 0; i < findLoopCount; i++ {
		bookNode := mytreeBalanced.Find(book1)
		if bookNode == nil || (book1.GetName() != bookNode.Data.GetName()) {
			t.Errorf("Found is incorrect, got: %s, want: %s.", bookNode.Data.GetName(), book1.GetName())
		}
	}
}


func TestEraseBalanced(t *testing.T) {
	book1 := book.Book{"MANKIND REVOLUTION", "CRUEL MAMBY"}
	bookNode := mytreeBalanced.Find(book1)
	mytreeBalanced.Erase(bookNode)
}
