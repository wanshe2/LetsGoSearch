package trie

func (node1 *heapNode) maxHeapCompare(node2 *heapNode) bool {
	return node1.deep > node2.deep
}

// MaxHeap is a max-heap of *heapNode.
type MaxHeap []*heapNode

func (h MaxHeap) Len() int { return len(h) }

// Less 我需要设计一个排序算法, 首先要侧重deep, 但是不能忽略 count
// 是否考虑 deep 设计为浮点数, deep^2 * count ?
func (h MaxHeap) Less(i, j int) bool {
	return h[i].maxHeapCompare(h[j])
}

func (h MaxHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *MaxHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(*heapNode))
}

func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}