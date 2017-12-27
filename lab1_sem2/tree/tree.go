package tree

import (
	"fmt"
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
	height int
}

type Tree struct {
	root *Node
}

func getHeight(p *Node) int {
	if p == nil {
		return 0
	}
	return p.height
}

func getBalanceFactor(p *Node) int {
	return getHeight(p.right) - getHeight(p.left)
}

func fixHeight(p *Node) {
	leftInt := getHeight(p.right)
	rightInt := getHeight(p.left)

	if leftInt > rightInt {
		p.height = leftInt + 1
	} else {
		p.height = rightInt + 1
	}
}

func rotateRight(p *Node) *Node {
	q := p.left
	p.left = q.right
	q.right = p
	fixHeight(p)
	fixHeight(q)
	return q
}

func rotateLeft(q *Node) *Node {
	p := q.right
	q.right = p.left
	p.left = q
	fixHeight(p)
	fixHeight(q)
	return p
}

func balance(p *Node) *Node {
    fixHeight(p)
    if getBalanceFactor(p) == 2 {
    	if getBalanceFactor(p.right) < 0 {
    		p.right = rotateRight(p.right)
		}
		return rotateLeft(p)
	}
	if getBalanceFactor(p) == -2 {
		if getBalanceFactor(p.left) > 0 {
			p.left = rotateLeft(p.left)
		}
		return rotateRight(p)
	}
	return p
}

func addToNodeRec(p *Node, field Abstract) *Node{
	if p == nil {
		p = new(Node)
		p.Data = field
		p.height = 1
		p.left = nil
		p.right = nil
		p.parent = p
		return p
	}
	if p.Data.Less(field) {
		p.left = addToNodeRec(p.left, field)
	} else {
		p.right = addToNodeRec(p.right, field)
	}
	return p
}

func addToNodeRecBalance(p *Node, field Abstract) *Node{
    if p == nil {
    	p = new(Node)
    	p.Data = field
    	p.height = 1
    	p.left = nil
    	p.right = nil
    	p.parent = p
    	return p
	}
	if p.Data.Less(field) {
		p.left = addToNodeRecBalance(p.left, field)
	} else {
		p.right = addToNodeRecBalance(p.right, field)
	}
	return balance(p)
}

func (t *Tree) AddRecBalance(field Abstract) {
	t.root = addToNodeRecBalance(t.root, field)
}

func (t *Tree) AddRec(field Abstract) {
	t.root = addToNodeRec(t.root, field)
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


func (t *Tree) ErasePrev(p *Node) {
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


func findRec(abstract Abstract, p *Node) *Node {
	if p == nil || p.Data.Equals(abstract) {
		return p
	}
	if p.Data.Less(abstract) {
		return findRec(abstract, p.left)
	} else {
		return findRec(abstract, p.right)
	}
}


func (t *Tree) FindRec(abstract Abstract) *Node {
	return findRec(abstract, t.root)
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
