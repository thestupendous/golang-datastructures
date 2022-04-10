package main

import "fmt"

//max heap
type MaxHeap struct{
	HeapSlice []int
}

//insert method builds up the heap
func (h *MaxHeap) Insert(val int) {
	h.HeapSlice = append(h.HeapSlice,val)
	i := len(h.HeapSlice)-1
	for ; ! (i/2<1) ; i/=2 {
		if h.HeapSlice[i/2] < h.HeapSlice[i] {
			h.HeapSlice[i/2],h.HeapSlice[i] = h.HeapSlice[i],h.HeapSlice[i/2]
		} else {
			break
		}
	}
}


//extract method returns the max item
func (h *MaxHeap) Extract() int {
	max := h.HeapSlice[1]
	i := 1
	lastI := len(h.HeapSlice)-1
	for ( (2*i<=lastI) && (2*i+1<=lastI) ) {
		//compare
		if (2*i < 2*i + 1) {
			h.HeapSlice[i],h.HeapSlice[2*i] = h.HeapSlice[2*i],h.HeapSlice[i]
			i = 2*i
		} else {
			h.HeapSlice[i],h.HeapSlice[2*i+1] = h.HeapSlice[2*i+1],h.HeapSlice[i]
			i = 2*i+1
		}
	}
	h.HeapSlice = h.HeapSlice[:lastI]
	return max
}

func main() {
	heap := MaxHeap{}
	var val int
	for i:=0;i<6;i++ {
		fmt.Scanf("%d",&val)
		heap.Insert(val)
	}

	fmt.Println("heap : ",heap)

	fmt.Println("max :",heap.Extract())

	fmt.Println("heap : ",heap)

}




