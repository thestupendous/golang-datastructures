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
	for ; ! (parent(i)<1) ; i=parent(i) {
		if h.HeapSlice[parent(i)] < h.HeapSlice[i] {
			h.Swap(i,parent(i))
		} else {
			break
		}
	}
}

//returns index of parent
func parent(i int) int {
	return i/2
}

//returns index of left child
func left(i int) int {
	return 2*i
}

//returns index of right child
func right(i int) int {
	return 2*i + 1
}

//swapping 2 values in heap by using their indices
func (h *MaxHeap) Swap(i, j int) {
	h.HeapSlice[i],h.HeapSlice[j] = h.HeapSlice[j],h.HeapSlice[i]
}

//extract method returns the max item
func (h *MaxHeap) Extract() int {
	if len(h.HeapSlice) == 0 {
		fmt.Println("Heap already Empty")
		return -999
	}

	max := h.HeapSlice[1]
	i := 1
	lastI := len(h.HeapSlice)-1
	h.HeapSlice[1] = h.HeapSlice[lastI]
	h.HeapSlice = h.HeapSlice[:lastI]
	lastI--
	goLeft := true


	//shifting down the heap to re organise the heap
	for left(i) <= lastI {
		if right(i) <= lastI {
			if h.HeapSlice[left(i)] > h.HeapSlice[right(i)] {		//when left child is greater
				goLeft = true
			} else {					//when right child is greater
				goLeft = false
			}
		} else {						//when only left child is there
			goLeft = true
		}

		if goLeft {
			if h.HeapSlice[i] < h.HeapSlice[left(i)] {
				h.Swap(i,left(i))
				i = left(i)
			} else {
				break
			}
		} else {
			if h.HeapSlice[i] < h.HeapSlice[right(i)] {
				h.Swap(i,right(i))
				i = right(i)
			} else {
				break
			}
		}
	}



	return max
}

func main() {
	heap := MaxHeap{}
	//indexing is starting from second location in the heap array, just for simplicity.
	// thus adding dummy value in the first location
	heap.HeapSlice = append(heap.HeapSlice,-999)
	var val,size int
	size = 6
	for i:=0;i<size;i++ {
		fmt.Scanf("%d",&val)
		heap.Insert(val)
	}

	fmt.Println("heap : ",heap)

	for i:=0; i<size; i++ {
		fmt.Println("max :",heap.Extract()," and heap : ",heap)
	}

}




