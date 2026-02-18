// Package linkedlist provides linked list implementations.
package linkedlist

import (
	"github.com/josestg/dsa/internal/generics"
	"github.com/josestg/dsa/sequence"
)

// BinaryNode is a node in a doubly linked list.
// Each node holds data and pointers to both the next and previous nodes.
//
//	         ┌───────────────────────────┐
//	◄────────│ prev │  data  │  next     │────────►
//	         └───────────────────────────┘
type BinaryNode[E any] struct {
	data E
	next *BinaryNode[E]
	prev *BinaryNode[E]
}

// NewBinaryNode creates a new node with the given data and neighbor pointers.
//
//	NewBinaryNode(42, nil, nil) creates:
//
//	       ┌───────────────────────────┐
//	nil ◄──│ prev │  42   │  next      │──► nil
//	       └───────────────────────────┘
func NewBinaryNode[E any](data E, next *BinaryNode[E], prev *BinaryNode[E]) *BinaryNode[E] {
	return &BinaryNode[E]{
		data: data,
		next: next,
		prev: prev,
	}
}

// DoublyLinkedList is a linked list where each node has both next and prev pointers.
// This enables efficient operations at both ends and bidirectional traversal.
//
//	head                           tail
//	  ↓                             ↓
//	┌───┐     ┌───┐     ┌───┐     ┌───┐
//	│ A │◄───►│ B │◄───►│ C │◄───►│ D │
//	└───┘     └───┘     └───┘     └───┘
//	  ↓                             ↓
//	 nil                           nil
//
// With prev pointers:
//   - Pop (remove from back): O(1) - can directly access prev node
//   - Backward iteration: O(n) - no need to build reversed copy
//
// Trade-off: Uses more memory than SinglyLinkedList (extra prev pointer per node).
type DoublyLinkedList[E any] struct {
	head *BinaryNode[E]
	tail *BinaryNode[E]
	size int
}

// NewDoublyLinkedList creates an empty doubly linked list.
//
//	head     tail
//	  ↓        ↓
//	 nil      nil      size = 0
func NewDoublyLinkedList[E any]() *DoublyLinkedList[E] {
	return &DoublyLinkedList[E]{}
}

// Size returns the number of elements in the list.
//
//	head                           tail
//	  ↓                              ↓
//	┌───┐     ┌───┐     ┌───┐     ┌───┐
//	│ A │◄───►│ B │◄───►│ C │◄───►│ D │
//	└───┘     └───┘     └───┘     └───┘
//
//	Size() → 4
//
// complexity:
//   - time : O(1)
//   - space: O(1)
func (l *DoublyLinkedList[E]) Size() int {
	return l.size
}

// Empty returns true if the list has no elements.
//
//	Empty list:              Non-empty list:
//	head = nil               head ◄──► [A] ◄──► [B]
//	tail = nil               tail ──────────────┘
//	size = 0                 size = 2
//
//	Empty() → true           Empty() → false
//
// complexity:
//   - time : O(1)
//   - space: O(1)
func (l *DoublyLinkedList[E]) Empty() bool {
	return l.size == 0 && l.head == nil && l.tail == nil
}

// Tail returns the last element without removing it.
//
//	head                           tail
//	  ↓                              ↓
//	┌───┐     ┌───┐     ┌───┐     ┌───┐
//	│ A │◄───►│ B │◄───►│ C │◄───►│ D │
//	└───┘     └───┘     └───┘     └───┘
//	                                ↑
//	                           Tail() → D
//
// complexity:
//   - time : O(1)
//   - space: O(1)
//
// Panics if the list is empty.
func (l *DoublyLinkedList[E]) Tail() E {
	if v, ok := l.TryTail(); !ok {
		panic("DoublyLinkedList.Tail: list is empty")
	} else {
		return v
	}
}

// TryTail attempts to return the last element.
// Returns (value, true) on success, or (zero, false) if empty.
//
// complexity:
//   - time : O(1)
//   - space: O(1)
func (l *DoublyLinkedList[E]) TryTail() (E, bool) {
	if l.Empty() {
		return generics.ZeroValue[E](), false
	}
	return l.tail.data, true
}

// Head returns the first element without removing it.
//
//	head                           tail
//	  ↓                              ↓
//	┌───┐     ┌───┐     ┌───┐     ┌───┐
//	│ A │◄───►│ B │◄───►│ C │◄───►│ D │
//	└───┘     └───┘     └───┘     └───┘
//	  ↑
//	Head() → A
//
// complexity:
//   - time : O(1)
//   - space: O(1)
//
// Panics if the list is empty.
func (l *DoublyLinkedList[E]) Head() E {
	if v, ok := l.TryHead(); !ok {
		panic("DoublyLinkedList.Head: list is empty")
	} else {
		return v
	}
}

// TryHead attempts to return the first element.
// Returns (value, true) on success, or (zero, false) if empty.
//
// complexity:
//   - time : O(1)
//   - space: O(1)
func (l *DoublyLinkedList[E]) TryHead() (E, bool) {
	if l.Empty() {
		return generics.ZeroValue[E](), false
	}
	return l.head.data, true
}

// Prepend adds an element to the front of the list.
//
//	Before Prepend(Z):
//
//	head                 tail
//	  ↓                    ↓
//	┌───┐     ┌───┐     ┌───┐
//	│ A │◄───►│ B │◄───►│ C │
//	└───┘     └───┘     └───┘
//
//	After Prepend(Z):
//
//	head                           tail
//	  ↓                              ↓
//	┌───┐     ┌───┐     ┌───┐     ┌───┐
//	│ Z │◄───►│ A │◄───►│ B │◄───►│ C │
//	└───┘     └───┘     └───┘     └───┘
//
// complexity:
//   - time : O(1)
//   - space: O(1)
//
// SCORE: 10
func (l *DoublyLinkedList[E]) Prepend(data E) {
	nn := NewBinaryNode(data, l.head, nil)

	if l.Empty() {
		l.head = nn
		l.tail = nn
	} else {
		l.head.prev = nn
		l.head = nn
	}
	l.size++
}

// Append adds an element to the back of the list.
//
//	Before Append(D):
//
//	head                 tail
//	  ↓                    ↓
//	┌───┐     ┌───┐     ┌───┐
//	│ A │◄───►│ B │◄───►│ C │
//	└───┘     └───┘     └───┘
//
//	After Append(D):
//
//	head                           tail
//	  ↓                              ↓
//	┌───┐     ┌───┐     ┌───┐     ┌───┐
//	│ A │◄───►│ B │◄───►│ C │◄───►│ D │
//	└───┘     └───┘     └───┘     └───┘
//
// complexity:
//   - time : O(1)
//   - space: O(1)
//
// SCORE: 15
func (l *DoublyLinkedList[E]) Append(data E) {
	nn := NewBinaryNode(data, nil, l.tail)
	if l.Empty() {
		l.head = nn
		l.tail = nn
	} else {
		l.tail.next = nn
		l.tail = nn
	}

	l.size++
}

// Pop removes and returns the last element.
//
//	Before Pop():
//
//	head                           tail
//	  ↓                              ↓
//	┌───┐     ┌───┐     ┌───┐     ┌───┐
//	│ A │◄───►│ B │◄───►│ C │◄───►│ D │
//	└───┘     └───┘     └───┘     └───┘
//
//	After Pop():
//
//	head                 tail
//	  ↓                    ↓
//	┌───┐     ┌───┐     ┌───┐
//	│ A │◄───►│ B │◄───►│ C │
//	└───┘     └───┘     └───┘
//
//	Pop() → D (removed element)
//
// complexity:
//   - time : O(1) - thanks to prev pointer!
//   - space: O(1)
//
// Note: This is O(1) unlike SinglyLinkedList.Pop() which is O(n).
// The prev pointer lets us directly access the second-to-last node.
//
// Panics if the list is empty.
func (l *DoublyLinkedList[E]) Pop() E {
	if data, ok := l.TryPop(); !ok {
		panic("DoublyLinkedList.Pop: list is empty")
	} else {
		return data
	}
}

// TryPop attempts to remove and return the last element.
// Returns (value, true) on success, or (zero, false) if empty.
//
// This is a non-panicking version of Pop.
//
// complexity:
//   - time : O(1)
//   - space: O(1)
//
// SCORE: 15
func (l *DoublyLinkedList[E]) TryPop() (E, bool) {
	if l.Empty() {
		return generics.ZeroValue[E](), false
	}

	currLastElement := l.tail.data
	if l.Size() == 1 {
		l.reset()
	} else {
		prev := l.tail.prev
		prev.next = nil
		l.tail = prev
		l.size--
	}
	return currLastElement, true
}

// Shift removes and returns the first element.
//
//	Before Shift():
//
//	head                           tail
//	  ↓                              ↓
//	┌───┐     ┌───┐     ┌───┐     ┌───┐
//	│ A │◄───►│ B │◄───►│ C │◄───►│ D │
//	└───┘     └───┘     └───┘     └───┘
//
//	After Shift():
//
//	head                 tail
//	  ↓                    ↓
//	┌───┐     ┌───┐     ┌───┐
//	│ B │◄───►│ C │◄───►│ D │
//	└───┘     └───┘     └───┘
//
//	Shift() → A (removed element)
//
// complexity:
//   - time : O(1)
//   - space: O(1)
//
// Panics if the list is empty.
func (l *DoublyLinkedList[E]) Shift() E {
	if data, ok := l.TryShift(); !ok {
		panic("DoublyLinkedList.Shift: list is empty")
	} else {
		return data
	}
}

// TryShift attempts to remove and return the first element.
// Returns (value, true) on success, or (zero, false) if empty.
//
// This is a non-panicking version of Shift.
//
// complexity:
//   - time : O(1)
//   - space: O(1)
//
// SCORE: 10
func (l *DoublyLinkedList[E]) TryShift() (E, bool) {
	if l.Empty() {
		return generics.ZeroValue[E](), false
	}
	tempHeadElement := l.head.data
	if l.Size() == 1 {
		l.reset()
	} else {
		newHeadElement := l.head.next
		newHeadElement.prev = nil
		l.head = newHeadElement
		l.size--
	}

	return tempHeadElement, true
}

// Iter iterates over all elements from front to back.
//
//	head                           tail
//	  ↓                              ↓
//	┌───┐     ┌───┐     ┌───┐     ┌───┐
//	│ A │◄───►│ B │◄───►│ C │◄───►│ D │
//	└───┘     └───┘     └───┘     └───┘
//	  ↓         ↓         ↓         ↓
//	 1st       2nd       3rd       4th     ← iteration order
//
// Example:
//
//	for value := range list.Iter {
//	    fmt.Println(value)  // prints A, B, C, D
//	}
//
// complexity:
//   - time : O(n)
//   - space: O(1)
func (l *DoublyLinkedList[E]) Iter(yield func(E) bool) {
	l.iterForward(func(u *BinaryNode[E]) bool { return yield(u.data) })
}

// IterBackward iterates over all elements from back to front.
//
//	head                           tail
//	  ↓                              ↓
//	┌───┐     ┌───┐     ┌───┐     ┌───┐
//	│ A │◄───►│ B │◄───►│ C │◄───►│ D │
//	└───┘     └───┘     └───┘     └───┘
//	  ↓         ↓         ↓         ↓
//	 4th       3rd       2nd       1st     ← iteration order
//
// Example:
//
//	for value := range list.IterBackward {
//	    fmt.Println(value)  // prints D, C, B, A
//	}
//
// complexity:
//   - time : O(n)
//   - space: O(1) - thanks to prev pointers!
//
// Note: Unlike SinglyLinkedList.IterBackward(), this is O(1) space
// because we can follow prev pointers directly.
func (l *DoublyLinkedList[E]) IterBackward(yield func(E) bool) {
	l.iterBackward(func(u *BinaryNode[E]) bool { return yield(u.data) })
}

// Enum iterates over all elements with their indices from front to back.
//
//	head                           tail
//	  ↓                              ↓
//	┌───┐     ┌───┐     ┌───┐     ┌───┐
//	│ A │◄───►│ B │◄───►│ C │◄───►│ D │
//	└───┘     └───┘     └───┘     └───┘
//	  0         1         2         3      ← indices
//
// Example:
//
//	for index, value := range list.Enum {
//	    fmt.Printf("%d: %v\n", index, value)
//	}
//
// complexity:
//   - time : O(n)
//   - space: O(1)
func (l *DoublyLinkedList[E]) Enum(yield func(int, E) bool) {
	i := 0
	for v := range l.Iter {
		if !yield(i, v) {
			break
		}
		i++
	}
}

// EnumBackward iterates over all elements with their indices from back to front.
//
//	head                           tail
//	  ↓                              ↓
//	┌───┐     ┌───┐     ┌───┐     ┌───┐
//	│ A │◄───►│ B │◄───►│ C │◄───►│ D │
//	└───┘     └───┘     └───┘     └───┘
//	  0         1         2         3      ← indices
//	  ↓         ↓         ↓         ↓
//	 4th       3rd       2nd       1st     ← enumeration order
//
// complexity:
//   - time : O(n)
//   - space: O(1)
func (l *DoublyLinkedList[E]) EnumBackward(yield func(int, E) bool) {
	i := l.Size() - 1
	for v := range l.IterBackward {
		if !yield(i, v) {
			break
		}
		i--
	}
}

// SCORE: 10
func (l *DoublyLinkedList[E]) iterForward(yield func(node *BinaryNode[E]) bool) {
	p := l.head
	for p != nil {
		if !yield(p) {
			break
		}
		p = p.next
	}
}

// SCORE: 10
func (l *DoublyLinkedList[E]) iterBackward(yield func(*BinaryNode[E]) bool) {
	p := l.tail
	for p != nil {
		if !yield(p) {
			break
		}
		p = p.prev
	}
}

// Get retrieves the element at the given index.
//
//	head                           tail       ← cursors
//	  ↓                              ↓
//	┌───┐     ┌───┐     ┌───┐     ┌───┐
//	│ A │◄───►│ B │◄───►│ C │◄───►│ D │                     ← values
//	└───┘     └───┘     └───┘     └───┘
//	  0         1         2         3                        ← indices
//
//	Get(0) → A
//	Get(2) → C
//	Get(3) → D
//
// complexity:
//   - time : O(n) worst case, O(n/2) average
//   - space: O(1)
//
// Note: The implementation can start from either head or tail,
// choosing whichever is closer to the target index.
//
// Panics if:
//   - The list is empty
//   - index < 0 or index >= Size()
func (l *DoublyLinkedList[E]) Get(index int) E {
	if v, ok := l.TryGet(index); !ok {
		if l.Empty() {
			panic("DoublyLinkedList.Get: list is empty")
		}
		panic("DoublyLinkedList.Get: index out of range")
	} else {
		return v
	}
}

// TryGet attempts to retrieve the element at the given index.
// Returns (value, true) on success, or (zero, false) if index is out of bounds.
//
// complexity:
//   - time : O(n)
//   - space: O(1)
//
// SCORE: 10
func (l *DoublyLinkedList[E]) TryGet(index int) (E, bool) {
	if l.Empty() || index < 0 || index >= l.size {
		return generics.ZeroValue[E](), false
	}

	p := l.head
	curr := 0
	for p != nil {
		if curr == index {
			return p.data, true
		}
		p = p.next
		curr++
	}
	return generics.ZeroValue[E](), false
}

// Set updates the element at the given index.
//
//	Before Set(2, X):
//
//	head                           tail       ← cursors
//	  ↓                              ↓
//	┌───┐     ┌───┐     ┌───┐     ┌───┐
//	│ A │◄───►│ B │◄───►│ C │◄───►│ D │                     ← values
//	└───┘     └───┘     └───┘     └───┘
//	  0         1         2         3                        ← indices
//
//	After Set(2, X):
//
//	head                           tail       ← cursors
//	  ↓                              ↓
//	┌───┐     ┌───┐     ┌───┐     ┌───┐
//	│ A │◄───►│ B │◄───►│ X │◄───►│ D │                     ← values
//	└───┘     └───┘     └───┘     └───┘
//	  0         1         2         3                        ← indices
//	                      ↑
//	                  updated
//
// complexity:
//   - time : O(n) worst case, O(n/2) average
//   - space: O(1)
//
// Panics if:
//   - The list is empty
//   - index < 0 or index >= Size()
func (l *DoublyLinkedList[E]) Set(index int, data E) {
	if !l.TrySet(index, data) {
		if l.Empty() {
			panic("DoublyLinkedList.Set: list is empty")
		}
		panic("DoublyLinkedList.Set: index out of range")
	}
}

// TrySet attempts to update the element at the given index.
// Returns true on success, false if index is out of bounds.
//
// complexity:
//   - time : O(n)
//   - space: O(1)
//
// SCORE: 5
func (l *DoublyLinkedList[E]) TrySet(index int, data E) bool {
	if l.Empty() || index < 0 || index >= l.size {
		return false
	}

	p := l.head
	curr := 0
	for p != nil {
		if curr == index {
			p.data = data
			return true
		}
		p = p.next
		curr++
	}
	return false
}

func (l *DoublyLinkedList[E]) checkBounds(index int) {
	if index < 0 || index >= l.Size() {
		panic("DoublyLinkedList.checkBounds: index out of range")
	}
}

// String returns the string representation of the list.
//
//	head                           tail
//	  ↓                              ↓
//	┌───┐     ┌───┐     ┌───┐     ┌───┐
//	│ 1 │◄───►│ 2 │◄───►│ 3 │◄───►│ 4 │
//	└───┘     └───┘     └───┘     └───┘
//
//	String() → "[1 2 3 4]"
//
// complexity:
//   - time : O(n)
//   - space: O(n)
func (l *DoublyLinkedList[E]) String() string {
	return sequence.String(l.Iter)
}

func (l *DoublyLinkedList[E]) reset() {
	l.head = nil
	l.tail = nil
	l.size = 0
}

// Insert adds an element at the given index.
// Elements at and after the index shift to higher indices.
//
//	Before Insert(2, X):
//
//	head                 tail                   ← cursors
//	  ↓                    ↓
//	┌───┐     ┌───┐     ┌───┐
//	│ A │◄───►│ B │◄───►│ C │                               ← values
//	└───┘     └───┘     └───┘
//	  0         1         2                                  ← indices
//	                ↑
//	          insert here
//
//	After Insert(2, X):
//
//	head                           tail       ← cursors
//	  ↓                              ↓
//	┌───┐     ┌───┐     ┌───┐     ┌───┐
//	│ A │◄───►│ B │◄───►│ X │◄───►│ C │                     ← values
//	└───┘     └───┘     └───┘     └───┘
//	  0         1         2         3                        ← indices
//	                      ↑
//	                  inserted
//
// Special cases:
//   - Insert(0, X) is equivalent to Prepend(X)
//   - Insert(Size(), X) is equivalent to Append(X)
//
// complexity:
//   - time : O(k) where k is the given index
//   - space: O(1)
//
// Panics if:
//   - index < 0 or index > Size()
//
// SCORE: 5
func (l *DoublyLinkedList[E]) Insert(index int, data E) {
	if index == 0 {
		np := NewBinaryNode(data, l.head, nil)
		if l.Empty() {
			l.head = np
			l.tail = np
		} else {
			l.head.prev = np
			l.head = np
		}
		l.size++
		return
	}

	if index == l.Size() {
		nn := NewBinaryNode(data, nil, l.tail)
		if l.Empty() {
			l.head = nn
			l.tail = nn
		} else {
			l.tail.next = nn
			l.tail = nn
		}
		l.size++
		return
	}
	l.checkBounds(index)

	p := l.head
	curr := 0
	for curr < index-1 {
		p = p.next
		curr++
	}

	nd := NewBinaryNode(data, p.next, p)

	p.next.prev = nd
	p.next = nd
	l.size++
}

// Remove deletes and returns the element at the given index.
//
//	Before Remove(2):
//
//	head                           tail       ← cursors
//	  ↓                              ↓
//	┌───┐     ┌───┐     ┌───┐     ┌───┐
//	│ A │◄───►│ B │◄───►│ C │◄───►│ D │                     ← values
//	└───┘     └───┘     └───┘     └───┘
//	  0         1         2         3                        ← indices
//	                      ↑
//	                   remove
//
//	After Remove(2):
//
//	head                 tail                 ← cursors
//	  ↓                    ↓
//	┌───┐     ┌───┐     ┌───┐
//	│ A │◄───►│ B │◄───►│ D │                               ← values
//	└───┘     └───┘     └───┘
//	  0         1         2                                  ← indices
//
//	Remove(2) → C (removed element)
//
// complexity:
//   - time : O(k) where k is min(index, Size()-index)
//   - space: O(1)
//
// Note: Traverses from whichever end (head or tail) is closer.
//
// Panics if:
//   - index < 0 or index >= Size()
func (l *DoublyLinkedList[E]) Remove(index int) E {
	if v, ok := l.TryRemove(index); !ok {
		panic("DoublyLinkedList.Remove: index out of range")
	} else {
		return v
	}
}

// TryRemove attempts to remove the element at the given index.
// Returns (value, true) on success, or (zero, false) if index is out of bounds.
//
// complexity:
//   - time : O(n)
//   - space: O(1)
//
// SCORE: 10
func (l *DoublyLinkedList[E]) TryRemove(index int) (E, bool) {
	if l.Empty() || index < 0 || index >= l.size {
		return generics.ZeroValue[E](), false
	}
	if index == 0 {
		return l.TryShift()
	}

	if index == l.Size()-1 {
		return l.TryPop()
	}

	var curr *BinaryNode[E]
	curr = l.head
	for i := 0; i < index; i++ {
		curr = curr.next
	}

	val := curr.data

	curr.prev.next = curr.next
	curr.next.prev = curr.prev

	curr.prev = nil
	curr.next = nil

	l.size--

	return val, true
}
