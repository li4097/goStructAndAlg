package main

import (
	"fmt"
	"struct/structure"
	"struct/utility"
)

type MyData struct {
	data int
}

func (i *MyData) Less(data utility.Compare) bool {
	if o, err := data.(*MyData); err == false {
		return false
	} else {
		return i.data < o.data
	}
}
func BinaryTreeFormat(n *structure.Node) {
	if data, err := n.GetData().(*MyData); err == false {
		fmt.Println("data is not MyData")
	} else {
		strFormat := ""
		var i uint8
		for i = 0; i < n.GetLevel(); i++ {
			strFormat += "   "
		}
		strFormat += "|-"
		fmt.Println(strFormat, data.data)
	}
}
func BinaryTreePrint(n *structure.Node) {
	if data, err := n.GetData().(*MyData); err == false {
		fmt.Println("data is not MyData")
	} else {
		fmt.Print(data.data, ",")
	}
}
func main() {

	var b structure.BinaryTree
	b.Insert(&MyData{5})
	b.Insert(&MyData{2})
	b.Insert(&MyData{3})
	b.Insert(&MyData{9})
	b.Insert(&MyData{6})
	b.Insert(&MyData{7})
	b.Insert(&MyData{1})

	b.TraversalFormat(BinaryTreeFormat)
	fmt.Println("\nL:")
	b.TraversalL(BinaryTreePrint)
	fmt.Println("\nM:")
	b.TraversalM(BinaryTreePrint)
	fmt.Println("\nR:")
	b.TraversalR(BinaryTreePrint)

	return
}
