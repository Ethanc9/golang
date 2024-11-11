package linkedlist

import "fmt"

type node struct{
	value int
	next *node
	prev *node
}

type linkedlist struct{
	head *node
	tail *node
}

func (ll *linkedlist) insert(x int){
	// Create a new node
	newNode := &node{value: x, next: nil, prev: nil}

	// If list is not initiaziled
	if(ll.head == nil){
		ll.head = newNode
		ll.tail = newNode
		return
	}

	curr := ll.head
	for curr.next != nil {
		curr = curr.next
	}	

	curr.next = newNode
	newNode.prev = curr
	ll.tail = newNode
	return
}

// Return node of first sighting
func (ll *linkedlist) contains(x int) *node {
	curr := ll.head
	for curr != nil { // Traverse until the end
		if curr.value == x { // If the value matches
			return curr
		}
		curr = curr.next
	}
	return nil // Return nil if not found
}

// Remove Last
func (ll *linkedlist) removeTail(){
	curr := ll.tail

	if curr.prev != nil{
		curr.prev.next = nil
		ll.tail = curr.prev
		return
	} else {
		ll.head = nil
		ll.tail = nil
		return
	}
}

// Remove First
func (ll *linkedlist) removeHead(){
	curr := ll.head
	if curr.next != nil{
		curr.next = ll.head
		curr.next.prev = nil
		return
	} else {
		ll.head = nil
		ll.tail = nil
		return
	}
}



func main() {
	// Create a new linked list
	ll := &linkedlist{}

	// Test insert method
	ll.insert(10)
	ll.removeHead()
	ll.insert(20)
	ll.insert(30)

	ll.removeTail()
	// Print the values in the linked list
	curr := ll.head


	ll.insert(10)

	curr = ll.head
	for curr != nil {
		fmt.Printf("%d ", curr.value)
		curr = curr.next
	}
}