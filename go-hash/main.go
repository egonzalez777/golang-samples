package main

import "fmt"

// Hash tables are the ascii values of each letter added and then divided by the array size
// The value returned is the identifier stored in the hash table

// Collision can happen, how to handle it.
// 1. Open addressing -> Store the second value to the next empty slot
// Searching, we check the original value, otherwise, we check +1 position.
// Open addressing can have a negative effect and cause hash tables to not be adequate
// 2. Seprate Chaining
// Linked lists used to chain the name within a position to handle these collisions.

// ArraySize const
const ArraySize = 7

// HashTable struct
type HashTable struct {
	array [ArraySize]*Bucket
}

// Bucket struct
type Bucket struct {
	head *BucketNode
}

// BucketNode struct
type BucketNode struct {
	key  string
	next *BucketNode
}

// Insert hash
func (h *HashTable) Insert(key string) {
	index := Hash(key)
	h.array[index].Insert(key)
}

// Search hash
func (h *HashTable) Search(key string) bool {
	index := Hash(key)
	return h.array[index].Search(key)
}

// Delete hash
func (h *HashTable) Delete(key string) {
	index := Hash(key)
	h.array[index].Delete(key)
}

// Insert bucket
func (b *Bucket) Insert(k string) {
	if !b.Search(k) {
		newNode := &BucketNode{key: k}
		newNode.next = b.head
		b.head = newNode
	} else {
		fmt.Println("Already exists")
	}
}

// Search bucket
func (b *Bucket) Search(k string) bool {
	currentNode := b.head
	for currentNode != nil {
		if currentNode.key == k {
			return true
		}
		currentNode = currentNode.next
	}
	return false
}

// Delete bucket
func (b *Bucket) Delete(k string) {
	if b.head.key == k {
		b.head = b.head.next
		return
	}
	previousNode := b.head
	for previousNode.next != nil {
		if previousNode.next.key == k {
			previousNode.next = previousNode.next.next
		}
		previousNode = previousNode.next
	}
}

// Hash func
func Hash(key string) int {
	sum := 0
	for _, v := range key {
		sum += int(v)
	}
	return sum % ArraySize
}

// Init func
func Init() *HashTable {
	result := &HashTable{}
	for i := range result.array {
		result.array[i] = &Bucket{}
	}
	return result
}

func main() {
	hash := Init()
	fmt.Println(hash)
	// This insert and search complexity is O(1), unless
	// the linked list becomes to great, which then alters this
	// to O(n)
	hash.Insert("randy")
	hash.Insert("eric")
	hash.Search("randy")
	hash.Delete("randy")
	fmt.Println("Randy is false: ", hash.Search("randy"))
	fmt.Println("Eric is true: ", hash.Search("eric"))
	// To be able to view this in action, run this in debug.
}
