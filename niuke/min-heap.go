package main

type MinHeap struct {
	data []int
}

func (h *MinHeap) Push(Val int) {
	h.data = append(h.data, Val)
	h.siftUp(len(h.data) - 1)
}

func (h *MinHeap) Pop() int {
	if len(h.data) == 0 {
		return -1
	}

	min := h.data[0]
	h.data[0] = h.data[len(h.data)-1]
	h.data = h.data[:len(h.data)-1]

	h.siftDown(0)
	return min
}

//上浮
func (h *MinHeap) siftUp(i int) {
	for i > 0 {
		parent := (i - 1) / 2
		if h.data[parent] <= h.data[i] {
			break
		}
		h.data[parent], h.data[i] = h.data[i], h.data[parent]
		i = parent
	}
}

//下沉
func (h *MinHeap) siftDown(i int) {
	for {
		left := i*2 + 1
		right := i*2 + 2

		minIndex := i
		if left < len(h.data) && h.data[left] < h.data[minIndex] {
			minIndex = left
		}
		if right < len(h.data) && h.data[right] < h.data[minIndex] {
			minIndex = right
		}
		if minIndex == i {
			break
		}

		h.data[i], h.data[minIndex] = h.data[minIndex], h.data[i]
		i = minIndex
	}
}
