package collection

import "fmt"

func ExampleNewStack() {
	stack := NewStack[int]()
	fmt.Println("Is empty:", stack.IsEmpty())
	fmt.Println("Len:", stack.Len())
	fmt.Println("Cap:", stack.Cap())
	_, peekOK := stack.Peek()
	fmt.Println("Peek OK:", peekOK)
	_, popOK := stack.Pop()
	fmt.Println("Pop OK:", popOK)

	// Output:
	// Is empty: true
	// Len: 0
	// Cap: 0
	// Peek OK: false
	// Pop OK: false
}

func ExampleNewStackWithCapacity() {
	stack := NewStackWithCapacity[int](10)
	fmt.Println("Is empty:", stack.IsEmpty())
	fmt.Println("Len:", stack.Len())
	fmt.Println("Cap:", stack.Cap())
	_, peekOK := stack.Peek()
	fmt.Println("Peek OK:", peekOK)
	_, popOK := stack.Pop()
	fmt.Println("Pop OK:", popOK)

	// Output:
	// Is empty: true
	// Len: 0
	// Cap: 10
	// Peek OK: false
	// Pop OK: false
}

func ExampleStack_Clear() {
	{
		stack := NewStack[int]()
		stack.Push(1)
		stack.Push(2)
		stack.Push(3)
		fmt.Println("Len before clear:", stack.Len())
		capBeforeClear := stack.Cap()
		stack.Clear()
		fmt.Println("Len after clear:", stack.Len())
		fmt.Println("Cap unchanged:", stack.Cap() == capBeforeClear)
	}

	stack := NewStackWithCapacity[int](10)
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	stack.Clear()
	fmt.Println("Cap after clear:", stack.Cap())

	// Output:
	// Len before clear: 3
	// Len after clear: 0
	// Cap unchanged: true
	// Cap after clear: 10
}

func ExampleStack_Peek() {
	stack := NewStack[int]()
	stack.Push(1)
	top, ok := stack.Peek()
	fmt.Println("Top:", top)
	fmt.Println("OK:", ok)
	fmt.Println("Len:", stack.Len())

	// Output:
	// Top: 1
	// OK: true
	// Len: 1
}

func ExampleStack_Pop() {
	stack := NewStack[int]()
	stack.Push(1)
	top, popOk := stack.Pop()
	fmt.Println("Top:", top)
	fmt.Println("Pop OK:", popOk)
	fmt.Println("Len:", stack.Len())
	_, peekOk := stack.Peek()
	fmt.Println("Peek OK:", peekOk)

	// Output:
	// Top: 1
	// Pop OK: true
	// Len: 0
	// Peek OK: false
}
