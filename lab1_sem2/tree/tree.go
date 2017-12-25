package tree

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type Abstract interface {
	GetName() string
	GetAuthor() string
	Less(abstract Abstract) bool
	Equals(abstract Abstract) bool
}

type Node struct {
	left, right, parent *Node
	Data                Abstract
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

func (ptr *Node) excludeChildrenTo(what *Node) {
	if ptr == ptr.parent.right {
		ptr.parent.right = what
	} else {
		ptr.parent.left = what
	}
	if what != nil {
		what.parent = ptr.parent
	}
}
func (ptr *Node) excludeParent() {
	if ptr == ptr.parent.right {
		ptr.parent.right = nil
	} else {
		ptr.parent.left = nil
	}
	ptr.parent = nil
}

func (t *Tree) Erase(p *Node) {
	if p.left != nil {
		ptr := t.max(p.left)
		p.Data = ptr.Data
		ptr.excludeChildrenTo(ptr.left)
	} else {
		if p.right == nil {
			p.excludeParent()
		} else {
			ptr := t.min(p.right)
			p.Data = ptr.Data
			ptr.excludeChildrenTo(ptr.right)
		}
	}
}

func (t *Tree) Find(abstract Abstract) *Node {
	ptr := t.root
	for !ptr.Data.Equals(abstract) {
		if ptr.Data.Less(abstract) {
			ptr = ptr.left
		} else {
			ptr = ptr.right
		}
		if ptr == nil {
			return ptr
		}
	}
	return ptr
}

func (ptr *Node) findRec(abstract Abstract) *Node {
	if ptr.Data.Equals(abstract) {
		return ptr
	}
	if nil != ptr.left && nil != ptr.left.findRec(abstract) {
		return ptr.left.findRec(abstract)
	}
	if nil != ptr.right && nil != ptr.right.findRec(abstract) {
		return ptr.right.findRec(abstract)
	}
	return nil
}

func (t *Tree) FindRec(abstract Abstract) *Node {
	return t.root.findRec(abstract)
}

func (t *Tree) min(p *Node) *Node {
	if p.left != nil {
		ptr := p.left
		for ; ptr.left != nil; ptr = ptr.left {
		}
		return ptr
	}
	return p
}

func (t *Tree) max(p *Node) *Node {
	if p.right != nil {
		ptr := p.right
		for ; ptr.right != nil; ptr = ptr.right {
		}
		return ptr
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
