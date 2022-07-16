package list

import "fmt"

func Example() {
	list := Init()

	// case 1: add and iter
	list.Add(1)
	list.Add(2)
	list.Add(3)
	fmt.Println("Len: ", list.Len())
	fmt.Println("Start iter direct order:")
	for list.Has() {
		fmt.Println(list.Next().Value)
	}
	fmt.Println("Start iter reverse order:")
	for list.HasPrevious() {
		fmt.Println(list.Previous().Value)
	}
	fmt.Println("Start iter direct order:")
	for list.Has() {
		fmt.Println(list.Next().Value)
	}

	// case 2: HasValue
	fmt.Printf("\nHas 2: %v\n", list.HasValue(2))
	fmt.Printf("Has 8: %v\n", list.HasValue(8))

	fmt.Println("Add 4 value")
	list.Add(4)
	fmt.Println("Len: ", list.Len())

	fmt.Println("\nStart iter:")
	for list.Has() {
		fmt.Println(list.Next().Value)
	}

	// case 3: delete
	// fmt.Println("\nDelete 2")
	// list.Delete(2)
	// fmt.Println("Len: ", list.Len())
	// for list.Has() {
	// 	fmt.Println(list.Next().Value)
	// }
}
