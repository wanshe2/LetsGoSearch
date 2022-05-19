package trie

import (
	"container/heap"
	"fmt"
)

type heapNode struct {
	node *Node
	deep int32 // deep represents the relationship with the pattern string
}

func (node1 *heapNode) compare(node2 *heapNode) bool {
	if node1.deep != node2.deep {
		return node1.deep < node2.deep
	} else {
		return node1.node.count < node2.node.count
	}
}

// Heap is a min-heap of *heapNode.
type Heap []*heapNode

func (h Heap) Len() int { return len(h) }

// Less 我需要设计一个排序算法, 首先要侧重deep, 但是不能忽略 count
// 是否考虑 deep 设计为浮点数, deep^2 * count ?
func (h Heap) Less(i, j int) bool {
	return h[i].compare(h[j])

	//if h[i].node.count != h[j].node.count {
	//	return h[i].node.count < h[j].node.count
	//} else {
	//	考虑是否维护一个 deep, 通过 deep 来记录它和模式串的相关度
	//return h[i].deep < h[j].deep
	//}
}
func (h Heap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *Heap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(*heapNode))
}

func (h *Heap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// This example inserts several ints into an IntHeap, checks the minimum,
// and removes them in order of priority.
func main() {
	h := &Heap{}
	heap.Init(h)
	heap.Push(h, &heapNode{node: newNode()})
	fmt.Printf("minimum: %d\n", (*h)[0]) // minimum: 1
	for h.Len() > 0 {
		fmt.Printf("%d ", heap.Pop(h)) // 1 2 3 5
	}
}

//type Heap struct {
//	size int
//	heap []*Node
//}
//
//func newHeap(nodes []*Node) *Heap {
//	heap := &Heap{size: len(nodes), heap: nodes}
//	if heap.size <= 1 {
//		return heap
//	}
//	for i := (heap.size - 1) / 2; i >= 0; i-- {
//		heap.shiftDown(i)
//	}
//}
//
//func (h *Heap) push(node *Node) {
//	h.heap = append(h.heap, node)
//	h.size += 1
//	h.shiftUp(h.size - 1)
//}
//
//func (h *Heap) peak() *Node {
//	return h.heap[0]
//}
//
//func (h *Heap) pop() *Node {
//	// swap heap[0] and heap[size - 1]
//	h.heap[0], h.heap[h.size-1] = h.heap[h.size-1], h.heap[0]
//	node := h.heap[h.size-1]
//	h.heap = delete(h.heap, node)
//
//}
