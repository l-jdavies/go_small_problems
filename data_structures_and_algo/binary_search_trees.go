/*
A binary tree is a tree in which each node has zero, one or two children.

A binary search tree is a binary tree that abides by additional rules:
* Each node can have a most one left child and one right child.
* A node's left descendants can only contain values that are less than the node itself.
* A node's right descendants can only contain values that are greater than the node itself.

Most complex aspect is deleting a node. Algorithm for deleting a node is:
* If the node being deleted has no children, simply delete it.
* If the node being deleted has one child, delete the node and plug the child into the spot where the deleted node was.
* When deleting a node with two children, replace the deleted node with the successor node. The successor node is the child node whose value is the least of all values that are greater than the deleted node.
* To find the successor node: visit the right child of the deleted value, and then keep on visiting the left child of each subsequent child until there are no more left children. The bottom value is the successor node.
* If the successor node has a right child, after plugging the successor node into the spot of the deleted node, take the former right child of the suc- cessor node and turn it into the left child of the former parent of the suc- cessor node.


Binary search trees have a time complexity of O(log N)
*/

package main

import "fmt"

type node struct {
	key   int
	left  *node
	right *node
}

// add a node to the binary search tree
func (n *node) insert(k int) {
	if n.key < k {
		if n.right == nil {
			n.right = &node{key: k}
		} else {
			n.right.insert(k)
		}
	} else if n.key > k {
		if n.left == nil {
			n.left = &node{key: k}
		} else {
			n.left.insert(k)
		}
	}
}

// returns true if binary search tree contains argument, else returns false
func (n *node) search(k int) bool {
	if n == nil {
		return false
	}

	if n.key < k {
		return n.right.search(k)
	} else if n.key > k {
		return n.left.search(k)
	}

	return true
}

func (n *node) delete(k int) {
	if n == nil {
		return
	}

	if k < n.key {
		n.left.delete(k)
	} else if k > n.key {
		n.right.delete(k)
	} else if k == n.key {
		if n.left == nil {
			n = n.right
			return
		} else if n.right == nil {
			n = n.left
			return
		}
	}
}

// print all nodes, starting from root
func (n *node) traverseTree() {
	cont := true
	if n == nil {
		cont = false
	}

	if cont {
		n.left.traverseTree()
		fmt.Println(n.key)
		n.right.traverseTree()
	}
}

func main() {
	binarySearch := &node{key: 100}
	binarySearch.insert(10)
	binarySearch.insert(50)
	binarySearch.insert(3000)
	binarySearch.traverseTree()
	//fmt.Println(binarySearch.search(50))
	//fmt.Println(binarySearch.search(5))

	binarySearch.delete(3000)
	binarySearch.traverseTree()
}
