package structure

import (
	"errors"
	"fmt"
	"struct/utility"
)

type Compare utility.Compare

type Stack struct {
	data  []*Node
	index int
}

func (self *Stack) Push(data *Node) {
	self.data[self.index] = data
	self.index++
}
func (self *Stack) Pop() *Node {
	if 0 == self.index {
		return nil
	}
	self.index--
	return self.data[self.index]
}
func (self *Stack) Empty() bool {
	return self.index == 0
}

type Node struct {
	data  Compare
	level uint8
	left  *Node
	right *Node
}

type BinaryTree struct {
	head *Node
	size uint32
}

func (n *Node) GetLeft() *Node {

	return n.left
}
func (n *Node) GetRight() *Node {
	return n.right
}

func (n *Node) GetData() Compare {
	return n.data
}
func (n *Node) GetLevel() uint8 {
	return n.level
}

func (self *BinaryTree) Insert(data Compare) error {
	if self.head == nil {
		self.head = new(Node)
		self.head.data = data
		self.head.level = 0
		self.head.left = nil
		self.head.right = nil
		self.size++
		return nil
	}
	n := self.head
	for true {
		if n.data.Less(data) {
			if n.right == nil {
				n.right = &Node{data, n.level + 1, nil, nil}
				break
			} else {
				n = n.right
			}
		} else if data.Less(n.data) {
			if n.left == nil {
				n.left = &Node{data, n.level + 1, nil, nil}
				break
			} else {
				n = n.left
			}
		} else {
			return errors.New("Data exist")
		}
	}
	self.size++
	return nil
}

func (self *BinaryTree) Search(data Compare) *Node {
	n := self.head
	for n != nil {
		if n.data.Less(data) {
			n = n.right
		} else if data.Less(n.data) {
			n = n.left
		} else {
			return n
		}
	}
	return nil
}

func (self *BinaryTree) TraversalL(foo func(*Node)) {
	n := self.head
	s := Stack{make([]*Node, 10, 10), 0}
	for n != nil {
		for n != nil {
			s.Push(n)
			n = n.left
		}
		for n == nil {
			n = s.Pop()
			if n == nil {
				break
			}
			foo(n)
			n = n.right
		}
	}
}
func (self *BinaryTree) TraversalR(foo func(*Node)) {
	n := self.head
	s := Stack{make([]*Node, 10, 10), 0}
	sL := Stack{make([]*Node, 10, 10), 0}

	for nil != n {
		for nil != n {
			s.Push(n)
			if nil != n.left {
				sL.Push(n.left)
			}
			n = n.right
		}
		n = sL.Pop()
	}

	n = s.Pop()
	for nil != n {
		foo(n)
		n = s.Pop()
	}
}
func (self *BinaryTree) TraversalM(foo func(*Node)) {
	n := self.head
	s := Stack{make([]*Node, 10, 10), 0}
	for n != nil {
		for n != nil {

			foo(n)
			s.Push(n)
			n = n.left
		}
		for n == nil {
			n = s.Pop()
			if n == nil {
				break
			}
			n = n.right
		}
	}
}

func (self *BinaryTree) TraversalFormat(foo func(*Node)) {
	n := self.head
	s := Stack{make([]*Node, 10, 10), 0}
	for n != nil || !s.Empty() {
		if n != nil {
			foo(n)
			s.Push(n)
			n = n.left
		} else {
			n = s.Pop()
			//foo(n)
			n = n.right
		}
	}
}
func (self *BinaryTree) Delete(data Compare) error {

	n := self.head
	var prev *Node = nil
	isLeft := true
	for n != nil {
		if n.data.Less(data) {
			prev = n
			n = n.right
			isLeft = false
		} else if data.Less(n.data) {
			prev = n
			n = n.left
			isLeft = true
		} else {
			tmp := n.left
			if nil == tmp {
				if isLeft {
					prev.left = n.right
				} else {
					prev.right = n.right
				}
			} else {
				for tmp.right != nil {
					tmp = tmp.right
				}
				tmp.right = n.right

				if isLeft {
					prev.left = tmp
				} else {
					prev.right = tmp
				}
			}
			break
		}
	}
	return nil
}
func BinaryTreeFormat(n *Node) {
	if data, err := n.GetData().(*utility.MyData); err == false {
		fmt.Println("data is not MyData")
	} else {
		strFormat := ""
		var i uint8
		for i = 0; i < n.GetLevel(); i++ {
			strFormat += "   "
		}
		strFormat += "|-"
		fmt.Println(strFormat, data.Data)
	}
}
func BinaryTreePrint(n *Node) {
	if data, err := n.GetData().(*utility.MyData); err == false {
		fmt.Println("data is not MyData")
	} else {
		fmt.Print(data.Data, ",")
	}
}
