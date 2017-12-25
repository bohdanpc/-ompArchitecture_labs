package main

import (
	"../book"
	"../tree"
	"fmt"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	mytree := new(tree.Tree)
	mytree.Add(book.Book{"Robinsone Crusoe", "Daniel Defoe"})
	mytree.Add(book.Book{"Nine Princes in Amber", "Roger Zelyazny"})
	mytree.Add(book.Book{"The Chronicles of Amber", "Roger Zelyazny"})
	boook := book.Book{"The Song of Ice and Fire", "Gourge Martine"}
	mytree.Add(boook)
	mytree.Add(book.Book{"Graph Monte-Kristo", "Daniel Defoe"})
	mytree.Add(book.Book{"The Dark Tower", "Stephen King"})
	mytree.Print()

	boook = book.Book{"Robinsone Crusoe", "Daniel Defoe"}
	auto := mytree.Find(boook)
	mytree.Erase(auto)

	fmt.Println("\nTree after calling Erase fuction:\n")

	//mytree.AddFromFile()
	book.AddFromCsvFile(mytree, "/home/thereptile/" +
		"GoglandProjects/CompArchitecture_labs/lab1_sem2/foreign_names (copy).csv")
	mytree.Print()
}
