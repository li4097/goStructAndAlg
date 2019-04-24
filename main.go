package main

import (
	"fmt"
	"struct/structure"
	"struct/utility"
)

func main() {

	var b structure.BinaryTree
	b.Insert(&utility.MyData{5})
	b.Insert(&utility.MyData{2})
	b.Insert(&utility.MyData{3})
	b.Insert(&utility.MyData{9})
	b.Insert(&utility.MyData{6})
	b.Insert(&utility.MyData{7})
	b.Insert(&utility.MyData{1})

	b.TraversalFormat(structure.BinaryTreeFormat)
	fmt.Println("\nL:")
	b.TraversalL(structure.BinaryTreePrint)
	fmt.Println("\nM:")
	b.TraversalM(structure.BinaryTreePrint)
	fmt.Println("\nR:")
	b.TraversalR(structure.BinaryTreePrint)

	return
}
