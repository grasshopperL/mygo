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

// make the sort by data[m0] <= data[m1] <= data[m2]
func medianOfThree(data Interface, m0, m1, m2 int) {
	if !data.Less(m0, m1) {
		data.Swap(m0, m1)
	}
	if !data.Less(m1, m2) {
		data.Swap(m1, m2)
		if !data.Less(m0, m2) {
			data.Swap(m0, m2)
		}
	}
}

//  swap value by range
func swapRange(data Interface, a, b, n int) {
	for i := 0; i < n; i++ {
		data.Swap(a + i, b + i)
	}
}

// I don't know why
func quickSort(data Interface, a, b, maxDepth int) {
	for b - a > 12 {
		if maxDepth == 0 {
			heapSort(data, a, b)
			return
		}
		maxDepth--
		mlo, mhi := doPivot(data, a, b)
		if mlo - a < b - mhi {
			quickSort(data, a, mlo, maxDepth)
			a = mhi 
		} else {
			quickSort(data, mhi, b, maxDepth)
			b = mlo
		}
	}
	if b - a  > 1 {
		for i := a + 6; i < b; i++ {
			if data.Less(i, i-6) {
				data.Swap(i, i-6)
			}
		}
		insertionSort(data, a, b)
	}
}