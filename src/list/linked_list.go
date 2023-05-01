// Package list ...
package list

// Node ...
type Node[A any] struct {
	Value A
	Next  *Node[A]
}

// List ...
type List[A comparable] struct {
	head *Node[A]
	tail *Node[A]
	size int
}

// FromList create List from Slice
func FromList[A comparable](lst []A) *List[A] {
	list := &List[A]{}
	var size int
	for i := 0; i < len(lst); i++ {
		list.Append(lst[i])
		size++
	}
	list.size = size
	return list
}

// Append appends element to List
func (lst *List[A]) Append(elm A) List[A] {
	newNode := &Node[A]{Value: elm}
	if lst.head == nil {
		lst.head = newNode
		lst.tail = newNode
		lst.size = 1
		return *lst
	}

	lst.tail.Next = newNode
	lst.tail = newNode
	lst.size += 1
	return *lst
}

// Prepend prepends element to List
func (lst *List[A]) Prepend(elm A) List[A] {
	newNode := &Node[A]{Value: elm}
	if lst.head == nil {
		lst.head = newNode
		lst.tail = newNode
		lst.size = 1
		return *lst
	}
	newNode.Next = lst.head
	lst.head = newNode
	lst.size += 1
	return *lst
}

// Delete deletes element from List
func (lst *List[A]) Delete(elm A) List[A] {
	if lst.head == nil {
		return *lst
	}
	if lst.head.Value == elm {
		lst.head = lst.head.Next
		lst.size -= 1
		return *lst
	}

	current := lst.head
	for current.Next != nil {
		if current.Next.Value == elm {
			current.Next = current.Next.Next
			lst.size -= 1
			break
		}
		current = current.Next
	}
	return *lst
}

// Find finds an element in a List
func (lst *List[A]) Find(elm A) Node[A] {
	current := lst.head
	for current != nil {
		if current.Value == elm {
			return *current
		}
		current = current.Next
	}
	return Node[A]{}
}

// Size return the size of the List
func (lst *List[A]) Size() int {
	return lst.size
}

// Head returns the Head of List
func (lst *List[A]) Head() *Node[A] {
	return lst.head
}

// Tail return the Tail of the List
func (lst *List[A]) Tail() *Node[A] {
	return lst.tail
}
