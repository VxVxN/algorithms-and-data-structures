package list

type Cell struct {
	Next     *Cell
	Previous *Cell
	Value    int
}

type List struct {
	currentElem *Cell
	head        *Cell
	tail        *Cell
	len         int
}

func Init() *List {
	head := &Cell{}
	tail := &Cell{}
	tail.Previous = head
	head.Next = tail
	return &List{
		head: head,
		tail: tail,
		len:  0,
	}
}

func (list *List) Add(value int) {
	newCell := &Cell{
		Value: value,
	}
	if list.head.Next == nil {
		list.head.Next = newCell
	} else {
		list.head.Next.Previous = newCell
		newCell.Next = list.head.Next
	}
	newCell.Previous = list.head

	list.head.Next = newCell
	list.currentElem = list.head
	list.len++
}

func (list *List) Len() int {
	return list.len
}

func (list *List) Has() bool {
	if list.currentElem == list.tail {
		list.currentElem = list.head
	}
	has := list.currentElem.Next != list.tail
	if !has {
		list.currentElem = list.head
	}
	return has
}

func (list *List) HasPrevious() bool {
	if list.currentElem == list.head {
		list.currentElem = list.tail
	}
	has := list.currentElem.Previous != list.head
	if !has {
		list.currentElem = list.tail
	}
	return has
}

func (list *List) Next() *Cell {
	list.currentElem = list.currentElem.Next
	return list.currentElem
}

func (list *List) Previous() *Cell {
	list.currentElem = list.currentElem.Previous
	return list.currentElem
}

func (list *List) HasValue(value int) bool {
	for list.Has() {
		if list.Next().Value == value {
			list.currentElem = list.head
			return true
		}
	}
	list.currentElem = list.head
	return false
}

func (list *List) Delete(value int) bool {
	for list.Has() {
		nextElem := list.currentElem.Next
		if nextElem != nil && nextElem.Value == value {
			list.currentElem.Next = nextElem.Next
			nextElem.Next.Previous = list.currentElem
			list.currentElem = list.head
			list.len--
			return true
		}
		list.currentElem = list.currentElem.Next
	}
	list.currentElem = list.head
	return false
}

func (list *List) Clear() bool {
	list.head.Next = list.tail
	list.tail.Previous = list.head
	list.len = 0
	list.currentElem = list.head
	return false
}
