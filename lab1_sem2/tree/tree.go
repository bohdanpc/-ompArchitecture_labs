package tree

import (
	"fmt"
	"container/list"
	"sync"
	"os"
)

const GoRoutinesCount = 4
var Wg sync.WaitGroup

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

func getHeight(pNode *Node) int {
	if pNode == nil {
		return 0
	}
	return pNode.height
}

func getBalanceFactor(pNode *Node) int {
	return getHeight(pNode.right) - getHeight(pNode.left)
}

func fixHeight(pNode *Node) {
	leftInt := getHeight(pNode.right)
	rightInt := getHeight(pNode.left)

	if leftInt > rightInt {
		pNode.height = leftInt + 1
	} else {
		pNode.height = rightInt + 1
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

func balance(pNode *Node) *Node {
    fixHeight(pNode)
    if getBalanceFactor(pNode) == 2 {
    	if getBalanceFactor(pNode.right) < 0 {
    		pNode.right = rotateRight(pNode.right)
		}
		return rotateLeft(pNode)
	}
	if getBalanceFactor(pNode) == -2 {
		if getBalanceFactor(pNode.left) > 0 {
			pNode.left = rotateLeft(pNode.left)
		}
		return rotateRight(pNode)
	}
	return pNode
}

func createNewNode(field Abstract) *Node {
	p := new(Node)
	p.Data = field
	p.height = 1
	p.left = nil
	p.right = nil
	p.parent = p
	return p
}


func addToNodeRec(p *Node, field Abstract) *Node{
	if p == nil {
		p = createNewNode(field)
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
		p = createNewNode(field)
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




func (t *Tree) Erase(pEraseNode *Node) {
	if pEraseNode.left != nil {
		pMaxNode := t.max(pEraseNode.left)
		pEraseNode.Data = pMaxNode.Data
		pMaxNode.excludeChildrenTo(pMaxNode.left)
	} else {
		if pEraseNode.right == nil {
			pEraseNode.excludeParent()
		} else {
			pMinNode := t.min(pEraseNode.right)
			pEraseNode.Data = pMinNode.Data
			pMinNode.excludeChildrenTo(pMinNode.right)
		}
	}
}


func (t *Tree) Find(abstract Abstract) *Node {
	pCurrNode := t.root
	for !pCurrNode.Data.Equals(abstract) {
		if pCurrNode.Data.Less(abstract) {
			pCurrNode = pCurrNode.left
		} else {
			pCurrNode = pCurrNode.right
		}
		if pCurrNode == nil {
			return pCurrNode
		}
	}
	return pCurrNode
}

/*
 * Additional recursive function
 * used in main interface of FindRec function
 */
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


func (t *Tree) min(pRoot *Node) *Node {
	if pRoot.left != nil {
		tmp := pRoot.left
		for ; tmp.left != nil; tmp = tmp.left {
		}
		return tmp
	}
	return pRoot
}

func (t *Tree) max(pRoot *Node) *Node {
	if pRoot.right != nil {
		tmp := pRoot.right
		for ; tmp.right != nil; tmp = tmp.right {
		}
		return tmp
	}
	return pRoot
}

func (t *Tree) printTree(pRoot *Node, depth int) {
	if pRoot != nil {
		for i := 0; i < depth; i++ {
			fmt.Print("	")
		}
		fmt.Println(pRoot.Data)
		t.printTree(pRoot.left, depth+1)
		t.printTree(pRoot.right, depth+1)
	}
}

func (t *Tree) Print() {
	t.printTree(t.root, 0)
}


func bfs(pRoot *Node, str string) {
	queue := list.New()
	queue.PushBack(pRoot)

	file, err := os.Create(str)
	if err != nil {
		return
	}
	defer file.Close()

	for ; queue.Len() > 0 ; {
		p_node := queue.Front()

		queue.Remove(p_node)
		file.WriteString(p_node.Value.(*Node).Data.GetAuthor())
		file.WriteString(p_node.Value.(*Node).Data.GetName())

		//fmt.Println(p_node.Value.(*Node).Data)

		if p_node.Value.(*Node).left != nil {
			queue.PushBack(p_node.Value.(*Node).left)
		}
		if p_node.Value.(*Node).right != nil {
			queue.PushBack(p_node.Value.(*Node).right)
		}
	}
}


func bfsParalel(pRoot *Node, str string) {
	defer Wg.Done()
	bfs(pRoot, str)
}


func (t *Tree) Bfs() {
	bfs(t.root, "CasualFile.dat")
}

func (t *Tree) BfsParalel() {
	file, err := os.Create("/home/thereptile/" +
		"GoglandProjects/CompArchitecture_labs/lab1_sem2/goRoutine0.dat")
	if err != nil {
		return
	}
	defer file.Close()

	pFirst := t.root
	pLeft1 := pFirst.left
	pRight1 := pFirst.right

	file.WriteString(pLeft1.Data.GetAuthor())
	file.WriteString(pLeft1.Data.GetName())

	file.WriteString(pRight1.Data.GetAuthor())
	file.WriteString(pRight1.Data.GetName())

	Wg.Add(GoRoutinesCount)

	go bfsParalel(pLeft1.left, "/home/thereptile/" +
	"GoglandProjects/CompArchitecture_labs/lab1_sem2/goRoutine1.dat")
	go bfsParalel(pLeft1.right, "/home/thereptile/" +
		"GoglandProjects/CompArchitecture_labs/lab1_sem2/goRoutine2.dat")
	go bfsParalel(pRight1.left, "/home/thereptile/" +
		"GoglandProjects/CompArchitecture_labs/lab1_sem2/goRoutine3.dat")
	go bfsParalel(pRight1.right, "/home/thereptile/" +
		"GoglandProjects/CompArchitecture_labs/lab1_sem2/goRoutine4.dat")
}



func dfs(pRoot *Node, str string) {
	stack := list.New()
	stack.PushBack(pRoot)

	file, err := os.Create(str)
	if err != nil {
		return
	}
	defer file.Close()

	for ; stack.Len() > 0 ; {
		p_node := stack.Back()

		stack.Remove(p_node)
		file.WriteString(p_node.Value.(*Node).Data.GetAuthor())
		file.WriteString(p_node.Value.(*Node).Data.GetName())

		if p_node.Value.(*Node).left != nil {
			stack.PushBack(p_node.Value.(*Node).left)
		}
		if p_node.Value.(*Node).right != nil {
			stack.PushBack(p_node.Value.(*Node).right)
		}
	}
}


func dfsParalel(pRoot *Node, str string) {
	defer Wg.Done()
	dfs(pRoot, str)
}


func (t *Tree) Dfs() {
	dfs(t.root, "CasualFile.dat")
}

func (t *Tree) DfsParalel() {
	file, err := os.Create("/home/thereptile/" +
		"GoglandProjects/CompArchitecture_labs/lab1_sem2/goRoutine0.dat")
	if err != nil {
		return
	}
	defer file.Close()

	pFirst := t.root
	pLeft1 := pFirst.left
	pRight1 := pFirst.right

	file.WriteString(pLeft1.Data.GetAuthor())
	file.WriteString(pLeft1.Data.GetName())

	file.WriteString(pRight1.Data.GetAuthor())
	file.WriteString(pRight1.Data.GetName())

	Wg.Add(GoRoutinesCount)

	go dfsParalel(pLeft1.left, "/home/thereptile/" +
		"GoglandProjects/CompArchitecture_labs/lab1_sem2/goRoutine1.dat")
	go dfsParalel(pLeft1.right, "/home/thereptile/" +
		"GoglandProjects/CompArchitecture_labs/lab1_sem2/goRoutine2.dat")
	go dfsParalel(pRight1.left, "/home/thereptile/" +
		"GoglandProjects/CompArchitecture_labs/lab1_sem2/goRoutine3.dat")
	go dfsParalel(pRight1.right, "/home/thereptile/" +
		"GoglandProjects/CompArchitecture_labs/lab1_sem2/goRoutine4.dat")
}


func walkImpl(pRoot *Node, ch chan Abstract) {
	if pRoot == nil {
		return
	}
	walkImpl(pRoot.left, ch)
	ch <- pRoot.Data
	walkImpl(pRoot.right, ch)
}

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func (t *Tree) Walk(ch chan Abstract) {
	walkImpl(t.root, ch)
	// Need to close the channel here
	close(ch)
}

func (t *Tree) Same(t2 *Tree) bool {
	w1, w2 := make(chan Abstract), make(chan Abstract)

	go t.Walk(w1)
	go t2.Walk(w2)

	for {
		v1, ok1 := <-w1
		v2, ok2 := <-w2
		if !ok1 || !ok2 {
			return ok1 == ok2
		}
		if v1 != v2 {
			return false
		}
	}
}