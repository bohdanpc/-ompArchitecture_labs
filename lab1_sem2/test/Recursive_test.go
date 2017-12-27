package main

import (
	"testing"
	"../book"
	"../tree"
)

const findLoopCount3 = 15000000

var treeRecursive = new(tree.Tree)

func TestAddRec(t *testing.T) {
	book.AddFromCsvFileRec(treeRecursive, "/home/thereptile/" +
		"GoglandProjects/CompArchitecture_labs/lab1_sem2/books.csv")
	treeRecursive.AddRec(book.Book{"MANKIND REVOLUTION", "CRUEL MAMBY"})
}


func TestFindRecursive(t *testing.T) {
	book1 := book.Book{"MANKIND REVOLUTION", "CRUEL MAMBY"}
	for i := 0; i < findLoopCount3; i++ {
		bookNode := treeRecursive.FindRec(book1)
		if book1.GetName() != bookNode.Data.GetName() {
			t.Errorf("Found is incorrect, got: %s, want: %s.", bookNode.Data.GetName(), book1.GetName())
		}
	}
}
