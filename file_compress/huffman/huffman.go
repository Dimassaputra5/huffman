package huffman

import (
	"container/heap"
	"sort"
)

type Node struct {
	freq  int
	Char  byte
	Left  *Node
	Right *Node
}
type PriorityQueue []*Node

func (pq PriorityQueue) Len() int           { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool { return pq[i].freq < pq[j].freq }
func (pq PriorityQueue) Swap(i, j int)      { pq[i], pq[j] = pq[j], pq[i] }

func (pq *PriorityQueue) Push(x interface{}) {
	*pq = append(*pq, x.(*Node))
}
func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

func BuildTree(freq map[byte]int) *Node {
	pq := &PriorityQueue{}
	heap.Init(pq)
	keys := make([]byte, 0, len(freq))
	for ch := range freq {
		keys = append(keys, ch)
	}
	sort.Slice(keys, func(i, j int) bool { return keys[i] < keys[j] })
	for _, ch := range keys {
		heap.Push(pq, &Node{freq: freq[ch], Char: ch})
	}
	if pq.Len() == 1 {
		dummy := &Node{freq: 0, Char: 0}
		heap.Push(pq, dummy)
	}
	for pq.Len() > 1 {
		left := heap.Pop(pq).(*Node)
		right := heap.Pop(pq).(*Node)
		parent := &Node{
			freq:  left.freq + right.freq,
			Left:  left,
			Right: right}
		heap.Push(pq, parent)
	}
	return heap.Pop(pq).(*Node)
}
func GenerateCodes(root *Node) map[byte]Code {
	codes := make(map[byte]Code)
	var traverse func(*Node, uint64, uint8)
	traverse = func(n *Node, val uint64, bits uint8) {
		if n.Left == nil && n.Right == nil {
			codes[n.Char] = Code{Value: val, Bits: bits}
			return
		}
		if n.Left != nil {
			traverse(n.Left, (val<<1)|0, bits+1)
		}
		if n.Right != nil {
			traverse(n.Right, (val<<1)|1, bits+1)
		}
	}
	traverse(root, 0, 0)
	return codes
}

type Code struct {
	Value uint64
	Bits  uint8
}
