package main

import (
	"../book"
	"../tree"
	//"fmt"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	mytree := new(tree.Tree)
	mytree.AddRec(book.Book{"Robinsone Crusoe", "Daniel Defoe"})
	mytree.AddRec(book.Book{"Tine Princes in Amber", "Roger Zelyazny"})
	mytree.AddRec(book.Book{"The Chronicles of Amber", "Roger Zelyazny"})
	boook := book.Book{"The Song of Ice and Fire", "Gourge Martine"}
	mytree.AddRec(boook)
	mytree.AddRec(book.Book{"Traph Monte-Kristo", "Daniel Defoe"})
	mytree.AddRec(book.Book{"The Dark Tower", "Stephen King"})
	mytree.Print()

	boook = book.Book{"Robinsone Crusoe", "Daniel Defoe"}
	auto := mytree.Find(boook)
	mytree.Erase(auto)

	//fmt.Println("\nTree after calling Erase fuction:\n")

	//mytree.AddFromFile()
	//book.AddFromCsvFile(mytree, "/home/thereptile/" +
	//	"GoglandProjects/CompArchitecture_labs/lab1_sem2/foreign_names (copy).csv")
	//mytree.Print()
}
