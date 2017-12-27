package book

import (
	"../tree"
	"os"
	"log"
	"encoding/csv"
	"bufio"
	"io"
)

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

func AddFromCsvFile(t *tree.Tree, fileName string) {
	csvFile, err := os.Open(fileName)

	if err != nil {
		log.Fatal(err)
	}
	defer csvFile.Close()

	reader := csv.NewReader(bufio.NewReader(csvFile))

	for {
		line, error := reader.Read()
		if error == io.EOF {
			break
		} else if error != nil {
			log.Fatal(error)
		}

		t.Add(Book{string(line[0]), string(line[1])})
	}
}


func AddFromCsvFileRecBalance(t *tree.Tree, fileName string) {
	csvFile, err := os.Open(fileName)

	if err != nil {
		log.Fatal(err)
	}
	defer csvFile.Close()

	reader := csv.NewReader(bufio.NewReader(csvFile))

	for {
		line, error := reader.Read()
		if error == io.EOF {
			break
		} else if error != nil {
			log.Fatal(error)
		}

		t.AddRecBalance(Book{string(line[0]), string(line[1])})
	}
}

func AddFromCsvFileRec(t *tree.Tree, fileName string) {
	csvFile, err := os.Open(fileName)

	if err != nil {
		log.Fatal(err)
	}
	defer csvFile.Close()

	reader := csv.NewReader(bufio.NewReader(csvFile))

	for {
		line, error := reader.Read()
		if error == io.EOF {
			break
		} else if error != nil {
			log.Fatal(error)
		}

		t.AddRec(Book{string(line[0]), string(line[1])})
	}
}