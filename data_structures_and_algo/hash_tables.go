// Jumin Lee's github repo: https://github.com/lee-junmin/tutorial-source/blob/master/data-structure-golang/hash-table/main.go

/*
`hash` is a hash function, which creates an integer by summing the binary value of each rune in the string and determining the reminder by dividing the sum with the array size. This hash code is used to determine where in the `HashTable` a value should be stored.

To prevent conflicts i.e. multiple values with the same hash code (collision handling), this structure uses separate chaining whereby each element in the array references a linked list, where collisions arise, thereby enabling multiple values to be stored at the same array index.
*/

package main

import (
	"fmt"
)

// Set the size of the hash table array
const ArraySize = 7

// HashTable will hold an array. Each element in the array will be a pointer to a 'bucket'
type HashTable struct {
	array [ArraySize]*bucket
}

// bucket is a linked list
// just references the head of the linked list
type bucket struct {
	head *bucketNode
}

// bucketNode is a node of the linked list, which holds the key and reference to the next node in the linked list
type bucketNode struct {
	key  string
	next *bucketNode
}

/////////HashTable methods///////

// Insert invokes the hash function to generate a hash code then returns the bucket stored in that index of the HashTable and invokes the bucket method, insert, to create a node with the given key and insert the node in the bucket
func (h *HashTable) Insert(key string) {
	index := hash(key)
	h.array[index].insert(key)
}

// return true if HashTable contains key, else return false
func (h *HashTable) Search(key string) bool {
	index := hash(key)
	return h.array[index].search(key)
}

// Delete value based on key
func (h *HashTable) Delete(key string) {
	index := hash(key)
	h.array[index].delete(key)
}

/////bucket (i.e. linked list) methods/////

// insert creates a new node with the key and inserts node in the bucket by making the node the head of the linked list
func (b *bucket) insert(k string) {
	if !b.search(k) {
		newNode := &bucketNode{key: k}
		newNode.next = b.head
		b.head = newNode
	} else {
		fmt.Println(k, "already exists")
	}
}

func (b *bucket) search(k string) bool {
	current := b.head

	for current != nil {
		if current.key == k {
			return true
		}
		current = current.next
	}

	return false
}

// delete value from linked list by essentially cutting it out of the reference chain
func (b *bucket) delete(k string) {
	// guard clause for if the head of the list is the value to be deleted
	if b.head.key == k {
		b.head = b.head.next
		return
	}

	previousNode := b.head
	for previousNode.next != nil {
		if previousNode.next.key == k {
			previousNode.next = previousNode.next.next
			return
		}

		previousNode = previousNode.next
	}
}

//////////////////////////////////////////////////////////////////////

// hash function to create a hash code from the input string (key)
func hash(key string) int {
	sum := 0
	for _, v := range key {
		sum += int(v)
	}

	return sum % ArraySize
}

// create a bucket in each slot of the HashTable arr
func Init() *HashTable {
	result := &HashTable{}

	for i := range result.array {
		result.array[i] = &bucket{}
	}

	return result
}

func main() {
	hashTable := Init()
	list := []string{
		"ERIC",
		"STAN",
		"RANDY",
		"KYLE",
	}

	for _, v := range list {
		hashTable.Insert(v)
	}

	hashTable.Delete("ERIC")
}
