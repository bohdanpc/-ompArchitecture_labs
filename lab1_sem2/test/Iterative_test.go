package main

import (
"testing"
"../book"
"../tree"
)

const findLoopCount2 = 15000000

var mytreeIterative = new(tree.Tree)

func TestAddIterative(t *testing.T) {

	book.AddFromCsvFile(mytreeIterative, "/home/thereptile/" +
		"GoglandProjects/CompArchitecture_labs/lab1_sem2/books.csv")
	mytreeIterative.Add(book.Book{"MANKIND REVOLUTION", "CRUEL MAMBY"})
}


func TestFindIterative(t *testing.T) {
	book1 := book.Book{"MANKIND REVOLUTION", "CRUEL MAMBY"}
	for i := 0; i < findLoopCount2; i++ {
		bookNode := mytreeIterative.Find(book1)
		if book1.GetName() != bookNode.Data.GetName() {
			t.Errorf("Found is incorrect, got: %s, want: %s.", bookNode.Data.GetName(), book1.GetName())
		}
	}
}

func TestEraseIterative(t *testing.T) {
	book1 := book.Book{"MANKIND REVOLUTION", "CRUEL MAMBY"}
	bookNode := mytreeIterative.Find(book1)
	mytreeIterative.Erase(bookNode)
}
