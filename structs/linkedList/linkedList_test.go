package linkedList

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLinkedList_Delete(t *testing.T) {
	// case 1: delete middle elem
	list := Init()
	list.Add(1)
	list.Add(2)
	list.Add(3)
	list.Delete(2)

	var result []int
	for list.Has() {
		result = append(result, list.Next().Value)
	}
	assert.Equal(t, []int{3, 1}, result)

	// case 2: delete first elem
	list = Init()
	list.Add(1)
	list.Add(2)
	list.Add(3)
	list.Delete(3)

	result = []int{}
	for list.Has() {
		result = append(result, list.Next().Value)
	}
	assert.Equal(t, []int{2, 1}, result)

	// case 3: delete last elem
	list = Init()
	list.Add(1)
	list.Add(2)
	list.Add(3)
	list.Delete(1)

	result = []int{}
	for list.Has() {
		result = append(result, list.Next().Value)
	}
	assert.Equal(t, []int{3, 2}, result)
}
