package linkedlist

import "fmt"

type Node struct {
	data string
	prev *Node
	next *Node
}

type List struct {
	head *Node
}

func (l *List) add(value string) {
	newNode := &Node{data: value}

	if l.head == nil {
		l.head = newNode
		return
	}

	curr := l.head
	for curr.next != nil {
		curr = curr.next
	}

	curr.next = newNode
}

func (l *List) remove(value string) {
	if l.head == nil {
		return
	}

	if l.head.data == value {
		l.head = l.head.next
		return
	}

	curr := l.head
	for curr.next != nil && curr.next.data != value {
		curr = curr.next
	}

	if curr.next != nil {
		curr.next = curr.next.next
	}
}

func (l *List) print() {
	curr := l.head
	for curr != nil {
		fmt.Printf("%s ", curr.data)
		curr = curr.next
	}
	fmt.Println()
}

func makeList(input []string) {
	list := &List{}

	for _, v := range input {
		list.add(v)
	}

	list.print()
}

func UseLinkedList() {
	var text = []string{"hello", "world", "from", "the", "computer"}

	makeList(text)
}
