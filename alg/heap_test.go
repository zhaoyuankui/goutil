package alg

import (
	"container/heap"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_IntHeap(t *testing.T) {
	h := &IntHeap{}
	h.Push(1)
	h.Push(2)
	h.Push(3)
	h.Push(5)
	h.Push(4)
	h.Push(6)
	assert.Equal(t, &IntHeap{1, 2, 3, 5, 4, 6}, h)
	assert.Equal(t, 6, h.Pop().(int))
	assert.Equal(t, 4, h.Pop().(int))
	assert.Equal(t, &IntHeap{1, 2, 3, 5}, h)
	heap.Init(h)
	assert.Equal(t, &IntHeap{1, 2, 3, 5}, h)
	heap.Push(h, 4)
	assert.Equal(t, &IntHeap{1, 2, 3, 5, 4}, h)
	heap.Push(h, 0)
	assert.Equal(t, &IntHeap{0, 2, 1, 5, 4, 3}, h)
	assert.Equal(t, 0, heap.Pop(h).(int))
	assert.Equal(t, &IntHeap{1, 2, 3, 5, 4}, h)
	assert.Equal(t, 1, heap.Pop(h).(int))
	assert.Equal(t, &IntHeap{2, 4, 3, 5}, h)
	assert.Equal(t, 2, heap.Pop(h).(int))
	assert.Equal(t, &IntHeap{3, 4, 5}, h)
	assert.Equal(t, 3, heap.Pop(h).(int))
	assert.Equal(t, &IntHeap{4, 5}, h)
	assert.Equal(t, 4, heap.Pop(h).(int))
	assert.Equal(t, &IntHeap{5}, h)
	assert.Equal(t, 5, heap.Pop(h).(int))
	assert.Equal(t, &IntHeap{}, h)
}

func Test_StringHeap(t *testing.T) {
	h := &StringHeap{}
	h.Push("1")
	h.Push("2")
	h.Push("3")
	h.Push("5")
	h.Push("4")
	h.Push("6")
	assert.Equal(t, &StringHeap{"1", "2", "3", "5", "4", "6"}, h)
	assert.Equal(t, "6", h.Pop().(string))
	assert.Equal(t, "4", h.Pop().(string))
	assert.Equal(t, &StringHeap{"1", "2", "3", "5"}, h)
	heap.Init(h)
	assert.Equal(t, &StringHeap{"1", "2", "3", "5"}, h)
	heap.Push(h, "4")
	assert.Equal(t, &StringHeap{"1", "2", "3", "5", "4"}, h)
	heap.Push(h, "0")
	assert.Equal(t, &StringHeap{"0", "2", "1", "5", "4", "3"}, h)
	assert.Equal(t, "0", heap.Pop(h).(string))
	assert.Equal(t, &StringHeap{"1", "2", "3", "5", "4"}, h)
	assert.Equal(t, "1", heap.Pop(h).(string))
	assert.Equal(t, &StringHeap{"2", "4", "3", "5"}, h)
	assert.Equal(t, "2", heap.Pop(h).(string))
	assert.Equal(t, &StringHeap{"3", "4", "5"}, h)
	assert.Equal(t, "3", heap.Pop(h).(string))
	assert.Equal(t, &StringHeap{"4", "5"}, h)
	assert.Equal(t, "4", heap.Pop(h).(string))
	assert.Equal(t, &StringHeap{"5"}, h)
	assert.Equal(t, "5", heap.Pop(h).(string))
	assert.Equal(t, &StringHeap{}, h)
}
