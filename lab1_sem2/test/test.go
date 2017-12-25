package main

import (
	"../book"
	"../tree"
	"fmt"
)

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
	auto := mytree.FindRec(boook)
	mytree.Erase(auto)

	fmt.Println("Tree after calling Erase fuction:")
	mytree.Print()
}
