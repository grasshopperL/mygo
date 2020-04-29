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