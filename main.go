package main

func main() {
	tree := IntTree{}
	/* tree.insert(13)
	tree.insert(10)
	tree.insert(2)
	tree.insert(12)
	tree.insert(15)
	tree.insert(25)
	tree.insert(20)
	tree.insert(31)
	tree.insert(29) */

	// Easier sample to look at
	/* tree.insert(4)
	tree.insert(5)
	tree.insert(2)
	tree.insert(1)
	tree.insert(3) */
	tree.fill([]int{4, 5, 2, 1, 3})

	// print the whole tree
	tree.printNode(tree.root, "")
}
