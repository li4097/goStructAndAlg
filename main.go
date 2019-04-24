package main

import (
	"fmt"
	"struct/structure"
)

func main() {

	var b structure.BinaryTree
	b.Insert(&MyData{5})
	b.Insert(&MyData{2})
	b.Insert(&MyData{3})
	b.Insert(&MyData{9})
	b.Insert(&MyData{6})
	b.Insert(&MyData{7})
	b.Insert(&MyData{1})

	b.TraversalFormat(structure.BinaryTreeFormat)
	fmt.Println("\nL:")
	b.TraversalL(structure.BinaryTreePrint)
	fmt.Println("\nM:")
	b.TraversalM(structure.BinaryTreePrint)
	fmt.Println("\nR:")
	b.TraversalR(structure.BinaryTreePrint)

	return
}
