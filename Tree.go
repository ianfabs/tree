package main

import (
	"container/list"
	"fmt"
)

type IntNode struct {
	key   int
	left  *IntNode
	right *IntNode
}

type IntTree struct {
	root *IntNode
}

func (tree *IntTree) printNode(node *IntNode, indent string) {
	var n = node
	var prefix = indent + "\\-"

	// var leftChild = node.left
	var rightChild = node.right
	var display = node.key
	fmt.Printf("%s %d\n", prefix, display)

	if n.left != nil {
		if rightChild != nil {
			tree.printNode(node.left, fmt.Sprintf("%s |", indent))
		} else {
			tree.printNode(node.left, fmt.Sprintf("%s  ", indent))
		}
	}
	if n.right != nil {
		tree.printNode(node.right, fmt.Sprintf("%s  ", indent))
	}

}

func (tree *IntTree) visit(node *IntNode) {
	fmt.Println(node.key)
}

func (tree *IntTree) fill(arr []int) {
	if len(arr) != 0 {
		tree.insert(arr[0])
		tree.fill(arr[1:])
	}
}

func (tree *IntTree) insert(el int) {
	tree.root = tree.recursiveInsert(tree.root, &IntNode{key: el})
}

func (tree *IntTree) recursiveInsert(node *IntNode, new *IntNode) *IntNode {
	if node == nil {
		return new
	} else if node.key < new.key {
		node.right = tree.recursiveInsert(node.right, new)
	} else {
		node.left = tree.recursiveInsert(node.left, new)
	}
	return node
}

func (tree *IntTree) search(p *IntNode, el int) *IntNode {
	for p != nil {
		if el == p.key {
			return p
		} else if el < p.key {
			p = p.left
		} else {
			p = p.right
		}
	}
	return nil
}

func (tree *IntTree) breadthFirst() {
	p := tree.root
	queue := Queue{queue: list.New()}

	if p != nil {
		queue.enqueue(&list.Element{Value: p})
		for !queue.isEmpty() {
			p = queue.dequeue().Value.(*IntNode)
			tree.visit(p)
			if p.left != nil {
				queue.enqueue(&list.Element{Value: p.left})
			}
			if p.right != nil {
				queue.enqueue(&list.Element{Value: p.right})
			}
		}
	}
}

func (tree *IntTree) inorder(p *IntNode) {
	if p != nil {
		tree.inorder(p.left)
		tree.visit(p)
		tree.inorder(p.right)
	}
}

func (tree *IntTree) preorder(p *IntNode) {
	if p != nil {
		tree.visit(p)
		tree.preorder(p.left)
		tree.preorder(p.right)
	}
}

func (tree *IntTree) postorder(p *IntNode) {
	if p != nil {
		tree.postorder(p.left)
		tree.postorder(p.right)
		tree.visit(p)
	}
}

func (tree *IntTree) dfs(p *IntNode, el int) *IntNode {
	if p != nil {
		if el == p.key {
			return p
		}
		tree.dfs(p.left, el)
		tree.dfs(p.right, el)
	}
	return nil
}
