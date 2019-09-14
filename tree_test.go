package main

import (
	"fmt"
	"testing"
)

// TestBreadthFirst - Test Breadth-First tree traversal
func TestBreadthFirst(t *testing.T) {
	tree := IntTree{}

	tree.fill([]int{4, 5, 2, 1, 3})

	tree.breadthFirst()
}

// TestInOrder - Test inorder tree traversal
func TestInOrder(t *testing.T) {
	tree := IntTree{}

	tree.fill([]int{4, 5, 2, 1, 3})

	tree.inorder(tree.root)
}

// TestPreOrder - Test inorder tree traversal
func TestPreOrder(t *testing.T) {
	tree := IntTree{}

	tree.fill([]int{4, 5, 2, 1, 3})

	tree.preorder(tree.root)
}

// TestPostOrder - Test inorder tree traversal
func TestPostOrder(t *testing.T) {
	tree := IntTree{}

	tree.fill([]int{4, 5, 2, 1, 3})

	tree.postorder(tree.root)
}

// TestPostOrder - Test inorder tree traversal
func TestSearch(t *testing.T) {
	tree := IntTree{}

	tree.fill([]int{4, 5, 2, 1, 3})

	fmt.Println(tree.search(tree.root, 2).key)

}

// TestPostOrder - Test inorder tree traversal
func TestDepthFirstSearch(t *testing.T) {
	tree := IntTree{}

	tree.fill([]int{4, 5, 2, 1, 3})

	fmt.Println(tree.dfs(tree.root, 2).key)

}
