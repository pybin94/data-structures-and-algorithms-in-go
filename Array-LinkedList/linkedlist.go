package main

import "fmt"

type Node[T any] struct {
	next *Node[T]
	val  T
}

func Append[T any](node *Node[T], next *Node[T]) *Node[T] {
	node.next = next
	return next
}

// Linked List
func main() {
	root := &Node[int]{nil, 10}
	node := root
	for i := 0; i < 5; i++ {
		node = Append(node, &Node[int]{nil, (i + 2) * 10})
	}

	for n := root; n != nil; n = n.next {
		fmt.Println("Val: ", n.val)
	}

	fmt.Println()

	node = root.next         // 20 node
	orignalNext := node.next // 30 node
	node = Append(node, &Node[int]{nil, 100})
	node.next = orignalNext

	for n := root; n != nil; n = n.next {
		fmt.Println("Val:", n.val)
	}

	fmt.Println()

	// search node
	node = root
	for i := 0; i < 3; i++ {
		node = node.next
	}
	fmt.Println("4th node:", node.val)
}
