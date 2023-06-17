package main

import "fmt"

type MaxHeap struct{
	heap []int
}

func (h *MaxHeap) Insert(key int){
	h.heap = append(h.heap, key)
	h.maxHeapifyUp(len(h.heap)-1)
}
func (h *MaxHeap) maxHeapifyUp(index int){
	for h.heap[parent(index)] < h.heap[index]{
		h.swap(parent(index), index)
		index = parent(index)
	}
}
func parent(i int) int{
	return (i-1)/2
}
func left(i int) int{
	return 2*i+1
}
func right(i int) int{
	return 2*i+2
}
func (h *MaxHeap) swap(i1, i2 int){
	h.heap[i1], h.heap[i2] = h.heap[i2], h.heap[i1]
}
func main(){
	m := &MaxHeap{}
	buildHeap := []int{10, 20, 30, 5, 7, 9, 11, 13, 15, 17}
	for _, v := range buildHeap{
		m.Insert(v)
		fmt.Println(m)
	}

}