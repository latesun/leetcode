package kthlargest

type minHeap struct {
	nums   []int
	size   int
	stored int
}

func newMinHeap(k int, nums []int) minHeap {
	h := minHeap{
		nums: make([]int, k+1),
		size: k,
	}

	for i := range nums {
		h.add(nums[i])
	}

	return h
}

func (h *minHeap) add(val int) {
	if h.stored == h.size {
		if h.nums[1] < val {
			h.nums[1] = val
			h.heapify(1)
		}
		return
	}

	h.stored++
	h.nums[h.stored] = val
	if h.size == h.stored {
		for i := h.size / 2; i >= 1; i-- {
			h.heapify(i)
		}
	}
}

func (h *minHeap) heapify(i int) {
	if h.stored <= 1 {
		return
	}

	for {
		minPos := i
		if i*2 <= len(h.nums)-1 && h.nums[i*2] < h.nums[i] {
			minPos = i * 2
		}

		if i*2+1 <= len(h.nums)-1 && h.nums[i*2+1] < h.nums[minPos] {
			minPos = i*2 + 1
		}

		if minPos == i {
			return
		}

		h.nums[i], h.nums[minPos] = h.nums[minPos], h.nums[i]
		i = minPos
	}
}

func (h *minHeap) top() int {
	if h.stored < 1 {
		return -1
	}
	return h.nums[1]
}

type KthLargest struct {
	minHeap
}

func Constructor(k int, nums []int) KthLargest {
	return KthLargest{minHeap: newMinHeap(k, nums)}
}

func (this *KthLargest) Add(val int) int {
	this.add(val)
	return this.top()
}
