package linkedList

type Cell struct {
	Next  *Cell
	Value int
}

type LinkedList struct {
	currentElem *Cell
	head        *Cell
	len         int
}

func Init() *LinkedList {
	return &LinkedList{
		head: &Cell{},
		len:  0,
	}
}

func (list *LinkedList) Add(value int) {
	list.head.Next = &Cell{
		Next:  list.head.Next,
		Value: value,
	}
	list.currentElem = list.head
	list.len++
}

func (list *LinkedList) AddAtEnd(value int) {
	for list.Has() {
		if list.Next().Next == nil {
			list.currentElem = &Cell{
				Value: value,
			}
			list.currentElem = list.head
			return
		}
	}
}

func (list *LinkedList) Len() int {
	return list.len
}

func (list *LinkedList) Has() bool {
	has := list.currentElem.Next != nil
	if !has {
		list.currentElem = list.head
	}
	return has
}

func (list *LinkedList) Next() *Cell {
	list.currentElem = list.currentElem.Next
	return list.currentElem
}

func (list *LinkedList) HasValue(value int) bool {
	for list.Has() {
		if list.Next().Value == value {
			list.currentElem = list.head
			return true
		}
	}
	list.currentElem = list.head
	return false
}

func (list *LinkedList) Clear() bool {
	list.head.Next = nil
	list.len = 0
	list.currentElem = list.head
	return false
}

func (list *LinkedList) Delete(value int) bool {
	for list.Has() {
		if list.Next().Next != nil && list.Next().Next.Value == value {
			list.currentElem.Next = list.currentElem.Next.Next
			list.currentElem = list.head
			list.len--
			return true
		}
	}
	list.currentElem = list.head
	return false
}
