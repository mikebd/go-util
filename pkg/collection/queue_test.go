package collection

import "fmt"

func ExampleNewQueue() {
	queue := NewQueue[int]()
	fmt.Println("Is empty:", queue.IsEmpty())
	fmt.Println("Len:", queue.Len())
	fmt.Println("Cap:", queue.Cap())
	_, peekOK := queue.Peek()
	fmt.Println("Peek OK:", peekOK)
	_, removeOK := queue.Remove()
	fmt.Println("Remove OK:", removeOK)

	// Output:
	// Is empty: true
	// Len: 0
	// Cap: 0
	// Peek OK: false
	// Remove OK: false
}

func ExampleNewQueueWithCapacity() {
	queue := NewQueueWithCapacity[int](10)
	fmt.Println("Is empty:", queue.IsEmpty())
	fmt.Println("Len:", queue.Len())
	fmt.Println("Cap:", queue.Cap())
	_, peekOk := queue.Peek()
	fmt.Println("Peek OK:", peekOk)
	_, removeOk := queue.Remove()
	fmt.Println("Remove OK:", removeOk)

	// Output:
	// Is empty: true
	// Len: 0
	// Cap: 10
	// Peek OK: false
	// Remove OK: false
}

func ExampleQueue_Clear() {
	{
		queue := NewQueue[int]()
		queue.Add(1)
		queue.Add(2)
		queue.Add(3)
		fmt.Println("Len before clear:", queue.Len())
		capBeforeClear := queue.Cap()
		queue.Clear()
		fmt.Println("Len after clear:", queue.Len())
		fmt.Println("Cap unchanged:", queue.Cap() == capBeforeClear)
	}

	queue := NewQueueWithCapacity[int](10)
	queue.Add(1)
	queue.Add(2)
	queue.Add(3)
	queue.Clear()
	fmt.Println("Cap after clear:", queue.Cap())

	// Output:
	// Len before clear: 3
	// Len after clear: 0
	// Cap unchanged: true
	// Cap after clear: 10
}

func ExampleQueue_Peek() {
	queue := NewQueue[int]()
	queue.Add(1)
	value, ok := queue.Peek()
	fmt.Println("Value:", value)
	fmt.Println("OK:", ok)
	fmt.Println("Len:", queue.Len())

	// Output:
	// Value: 1
	// OK: true
	// Len: 1
}

func ExampleQueue_Remove() {
	queue := NewQueue[int]()
	queue.Add(1)
	value, removeOk := queue.Remove()
	fmt.Println("Value:", value)
	fmt.Println("Remove OK:", removeOk)
	fmt.Println("Len:", queue.Len())
	_, peekOk := queue.Peek()
	fmt.Println("Peek OK:", peekOk)

	// Output:
	// Value: 1
	// Remove OK: true
	// Len: 0
	// Peek OK: false
}
