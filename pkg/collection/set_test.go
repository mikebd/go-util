package collection

import "fmt"

func ExampleSet() {
	set := make(Set[int])

	set.Add(1)
	set.Add(2)
	set.Add(3)

	fmt.Println(set.Contains(1))
	fmt.Println(set.Contains(4))

	set.Remove(2)

	fmt.Println(set.Contains(2))

	fmt.Println(len(set))

	// Output:
	// true
	// false
	// false
	// 2
}
