/**
 * @Author: liubaoshuai3
 * @Date: 2020/4/20 22:48
 * @File: heapr
 * @Description:
 */

package tools

import "sort"

// maybe implement the interface do a heap
type Interface interface {
	sort.Interface
	Push(x interface{})
	Pop()
}