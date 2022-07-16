package linkedList

import "fmt"

func Example() {
	list := Init()

	fmt.Println("========== Linked list examples ==========\n")
	fmt.Println("========== case 1: add and iter")
	list.Add(1)
	list.Add(2)
	list.Add(3)
	fmt.Println("Add: 1,2,3")
	fmt.Println("Len: ", list.Len())
	fmt.Println("Start iter:")
	for list.Has() {
		fmt.Println(list.Next().Value)
	}

	fmt.Println("\n========== case 2: HasValue")
	fmt.Printf("Has 2: %v\n", list.HasValue(2))
	fmt.Printf("Has 8: %v\n", list.HasValue(8))

	fmt.Println("Add: 4 value")
	list.Add(4)
	fmt.Println("Len: ", list.Len())

	fmt.Println("\nStart iter:")
	for list.Has() {
		fmt.Println(list.Next().Value)
	}

	fmt.Println("\n========== case 3: delete middle elem")
	fmt.Println("Delete: 2")
	list.Delete(2)
	fmt.Println("Len: ", list.Len())
	fmt.Println("\nStart iter:")
	for list.Has() {
		fmt.Println(list.Next().Value)
	}

	fmt.Println("\n========== case 4: clear")
	fmt.Println("Clear list")
	list.Clear()
	fmt.Println("Len: ", list.Len())
	for list.Has() {
		fmt.Println(list.Next().Value)
	}
	fmt.Println("\n")
}
