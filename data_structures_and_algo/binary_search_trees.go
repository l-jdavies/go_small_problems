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
	toDelete, parent := n.findNodeAndParent(k, &node{}) // pass in empty node as parent

	fmt.Println(toDelete, parent)
	if toDelete.left == nil && toDelete.right == nil { // toDelete has no children
		removeNode(toDelete, parent)
	}

	if toDelete.hasBothChildren() {
		// implement
	} else if toDelete.hasLeftLeaf() { // left child is end of tree
		toDelete.replaceNode(toDelete.left, parent)
	} else if toDelete.hasRightLeaf() { // right child is end of tree
		toDelete.replaceNode(toDelete.right, parent)
	} else {

	}
}

func (n *node) replaceNode(replacement, parent *node) {
	if parent.left == n {
		parent.left = replacement
	} else if parent.right == n {
		parent.right = replacement
	}
}

func (n *node) hasBothChildren() bool {
	return n.left != nil && n.right != nil
}

func (n *node) hasLeftLeaf() bool {
	return n.left != nil && n.left.left == nil && n.left.right == nil
}

func (n *node) hasRightLeaf() bool {
	return n.right != nil && n.right.left == nil && n.right.right == nil
}

func removeNode(target, parent *node) {
	if parent.left == target {
		parent.left = nil
	} else if parent.right == target {
		parent.right = nil
	}
}

func (n *node) findNodeAndParent(k int, p *node) (target, parent *node) {
	if n.key < k {
		if n.right.key == k {
			target = n.right
			parent = n
			fmt.Println("right is match", n.right, n)
			//return
		} else {
			n.right.findNodeAndParent(k, n)
		}
	} else if n.key > k {
		if n.left.key == k {
			target = n.left
			parent = n
			fmt.Println("left is match", target, parent)
			//return
		} else {
			n.left.findNodeAndParent(k, n)
		}
	} else {
		target = n
		parent = p
	}

	return
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
	binarySearch.insert(5)
	binarySearch.insert(200)
	binarySearch.insert(5000)
	binarySearch.insert(1)
	//fmt.Println(binarySearch.search(50))
	//fmt.Println(binarySearch.search(5))

	binarySearch.traverseTree()
	binarySearch.delete(5) // 1 should replace 5
	binarySearch.traverseTree()
	//	binarySearch.delete(100)
}