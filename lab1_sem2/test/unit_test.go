package main

import (
	"testing"
	"../book"
    "../tree"
)


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


func TestFind(t *testing.T) {
	mytree := new(tree.Tree)
	mytree.Add(book.Book{"Robinsone Crusoe", "Daniel Defoe"})
	mytree.Add(book.Book{"Nine Princes in Amber", "Roger Zelyazny"})
	mytree.Add(book.Book{"The Chronicles of Amber", "Roger Zelyazny"})

	book.AddFromCsvFile(mytree, "/home/thereptile/" +
		"GoglandProjects/CompArchitecture_labs/lab1_sem2/books.csv")

	book1 := book.Book{"Robinsone Crusoe", "Daniel Defoe"}
	bookNode := mytree.Find(book1)
	//*bookNode.data.GetName()
	if book1.GetName() != bookNode.Data.GetName() {
		t.Errorf("Found is incorrect, got: %s, want: %s.", bookNode.Data.GetName(), book1.GetName())
	}
}

func TestErase(t *testing.T) {
	mytree := new(tree.Tree)
	mytree.Add(book.Book{"Robinsone Crusoe", "Daniel Defoe"})
	mytree.Add(book.Book{"Nine Princes in Amber", "Roger Zelyazny"})
	book1 := book.Book{"The Song of Ice and Fire", "Gourge Martine"}
	mytree.Add(book1);
	book.AddFromCsvFile(mytree, "/home/thereptile/"+
		"GoglandProjects/CompArchitecture_labs/lab1_sem2/books.csv")

	book2 := book.Book{"Robinsone Crusoe", "Daniel Defoe"}
	bookNode := mytree.Find(book2)
	mytree.Erase(bookNode)
}

func TestAdd(t *testing.T) {
	mytree := new(tree.Tree)
	mytree.Add(book.Book{"Robinsone Crusoe", "Daniel Defoe"})
	mytree.Add(book.Book{"Nine Princes in Amber", "Roger Zelyazny"})

	book.AddFromCsvFile(mytree, "/home/thereptile/" +
		"GoglandProjects/CompArchitecture_labs/lab1_sem2/books.csv")
}