/**
 * @Author: liubaoshuai3
 * @Date: 2020/4/28 15:08
 * @File: pq
 * @Description:
 */

package tools

// struct of item in queue
type Item struct {
	value string
	priority int
	index int
}

// container of queue
type PQ []*Item

// length of pq
func (pq *PQ) Len() int {
	return len(*pq)
}

// compare to decide make min-heap or max-heap
func (pq *PQ) Less(i, j int) bool {
	return (*(*pq)[i]).priority > (*pq)[j].priority
}

// swap the value by index
func (pq *PQ) Swap(i, j int) {
	(*pq)[i].index = j
	(*pq)[j].index = i
	(*pq)[i], (*pq)[j] = (*pq)[j], (*pq)[i]
}

// push x to pq
func (pq *PQ) Push(x interface{}) {
	l := len(*pq)
	item := x.(Item)
	item.index = l
	*pq = append(*pq, &item)
}
// remove tail item
func (pq *PQ) Pop() interface{} {
	l := len(*pq)
	item := (*pq)[l - 1]
	(*pq)[l - 1] = nil
	*pq = (*pq)[:l - 1]
	item.index = -1
	return item.value
}

// update the item in heap
func (pq *PQ) update(item *Item, value string, priority int) {
	item.priority = priority
	item.value = value
	Fix(pq, item.index)
}



