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

type node struct {
	left, right, parent *node
	data                Abstract
}

type Tree struct {
	root *node
}

func (t *Tree) Add(field Abstract) {

	if t.root == nil {
		t.root = new(node)
		t.root.data = field
		t.root.parent = t.root
		return
	}

	for p := t.root; p != nil; {
		p_parent := p
		if p.data.Less(field) {
			p = p.left
		} else {
			p = p.right
		}
		if p == nil {
			if p_parent.data.Less(field) {
				p_parent.left = new(node)
				p_parent.left.data = field
				p_parent.left.parent = p_parent
			} else {
				p_parent.right = new(node)
				p_parent.right.data = field
				p_parent.right.parent = p_parent
			}
		}
	}
}

func (t *Tree) Erase(p *node) {
	if p.left != nil {
		tmp := t.max(p.left)
		p.data = tmp.data
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
			p.data = tmp.data
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

func (t *Tree) Find(abstract Abstract) *node {
	tmp := t.root
	for !tmp.data.Equals(abstract) {
		if tmp.data.Less(abstract) {
			tmp = tmp.left
		} else {
			tmp = tmp.right
		}
	}
	return tmp
}

func (t *Tree) min(p *node) *node {
	if p.left != nil {
		tmp := p.left
		for ; tmp.left != nil; tmp = tmp.left {
		}
		return tmp
	}
	return p
}

func (t *Tree) max(p *node) *node {
	if p.right != nil {
		tmp := p.right
		for ; tmp.right != nil; tmp = tmp.right {
		}
		return tmp
	}
	return p
}

func (t *Tree) printTree(p *node, depth int) {
	if p != nil {
		for i := 0; i < depth; i++ {
			fmt.Print("	")
		}
		fmt.Println(p.data)
		t.printTree(p.left, depth+1)
		t.printTree(p.right, depth+1)
	}
}

func (t *Tree) Print() {
	t.printTree(t.root, 0)
}
