package book

import "../tree"

type Book struct {
	Name, Author string
}

func (b Book) Equals(book tree.Abstract) bool {
	return b.Name == book.GetName() && b.Author == book.GetAuthor()
}

func (b Book) Less(book tree.Abstract) bool {
	return b.Name < book.GetName()
}

func (b Book) GetAuthor() string {
	return b.Author
}

func (b Book) GetName() string {
	return b.Name
}
