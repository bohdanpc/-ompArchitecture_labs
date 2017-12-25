package tree

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"log"
)

type Abstract interface {
	GetName() string
	GetAuthor() string
	Less(abstract Abstract) bool
	Equals(abstract Abstract) bool
}

type Node struct {
	left, right, parent *Node
	Data Abstract
}

type Tree struct {
	root *Node
}

func (t *Tree) Add(field Abstract) {

	if t.root == nil {
		t.root = new(Node)
		t.root.Data = field
		t.root.parent = t.root
		return
	}

	for p := t.root; p != nil; {
		p_parent := p
		if p.Data.Less(field) {
			p = p.left
		} else {
			p = p.right
		}
		if p == nil {
			if p_parent.Data.Less(field) {
				p_parent.left = new(Node)
				p_parent.left.Data = field
				p_parent.left.parent = p_parent
			} else {
				p_parent.right = new(Node)
				p_parent.right.Data = field
				p_parent.right.parent = p_parent
			}
		}
	}
}

func (t *Tree) Erase(p *Node) {
	if p.left != nil {
		tmp := t.max(p.left)
		p.Data = tmp.Data
		if tmp == tmp.parent.right {
			tmp.parent.right = tmp.left
		} else {
			tmp.parent.left = tmp.left
		}
		if tmp.left != nil {
			tmp.left.parent = tmp.parent
		}
	} else {
		if p.right == nil {
			if p.parent.left == p {
				p.parent.left = nil
			} else {
				p.parent.right = nil
			}
			p.parent = nil
		} else {
			tmp := t.min(p.right)
			p.Data = tmp.Data
			if tmp == tmp.parent.left {
				tmp.parent.left = tmp.right
			} else {
				tmp.parent.right = tmp.right
			}
			if tmp.right != nil {
				tmp.right.parent = tmp.parent
			}
		}
	}
}

func (t *Tree) Find(abstract Abstract) *Node {
	tmp := t.root
	for !tmp.Data.Equals(abstract) {
		if tmp.Data.Less(abstract) {
			tmp = tmp.left
		} else {
			tmp = tmp.right
		}
		if tmp == nil {
			return tmp
		}
	}
	return tmp
}

func (t *Tree) min(p *Node) *Node {
	if p.left != nil {
		tmp := p.left
		for ; tmp.left != nil; tmp = tmp.left {
		}
		return tmp
	}
	return p
}

func (t *Tree) max(p *Node) *Node {
	if p.right != nil {
		tmp := p.right
		for ; tmp.right != nil; tmp = tmp.right {
		}
		return tmp
	}
	return p
}

func (t *Tree) printTree(p *Node, depth int) {
	if p != nil {
		for i := 0; i < depth; i++ {
			fmt.Print("	")
		}
		fmt.Println(p.Data)
		t.printTree(p.left, depth+1)
		t.printTree(p.right, depth+1)
	}
}

func (t *Tree) Print() {
	t.printTree(t.root, 0)
}

func (t *Tree) AddFromFile() {
	file, err := os.Open("/home/thereptile/" +
		"GoglandProjects/CompArchitecture_labs/lab1_sem2/info.dat")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {

		ucl := strings.ToUpper(scanner.Text())
			fmt.Println(ucl)
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}
}
