/**
 * @Author: liubaoshuai3
 * @Date: 2020/4/29 21:58
 * @File: sort
 * @Description:
 */

package sort

// the interface of sort
type Interface interface {
	Len() int
	Less(i, j int) bool
	Swap(i, j int)
}

// insert sort
func insertionSort(data Interface, a, b int)  {
	for i := a + 1; i < b; i++ {
		// my version
		//for j := i; j > a; j-- {
		//	if data.Less(j, j - 1) {
		//		data.Swap(j, j - 1)
		//	}
		//}
		for j := i; j > a && data.Less(j, j -1 ); j-- {
			data.Swap(j, j - 1)
		}
	}
}

func siftDown(data Interface, lo, hi, first int) {
	root := lo
	for {
		child := 2 * root + 1
		if child > hi {
			break
		}
		if child + 1 < hi && data.Less(first + child, first + child + 1) {
			data.Swap(first + child, first + child + 1)
		}
		if !data.Less(first + root, first + child) {
			return
		}
		data.Swap(first + root, first + child)
		root = child
	}
}

// heap sort
func heapSort(data Interface, a, b int)  {
	first := a
	lo := 0
	hi := b - a
	for i := (hi - 1) / 2; i >= 0; i-- {
		siftDown(data, i, hi, first)
	}
	for i := hi - 1; i >= 0; i-- {
		data.Swap(first, first + i)
		siftDown(data, lo, i, first)
	}
}