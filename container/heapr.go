/**
 * @Author: liubaoshuai3
 * @Date: 2020/7/21 18:13
 * @File: heapr
 * @Description:
 */

package container

import "sort"

type Interface interface {
	sort.Interface
	Push(x interface{})
	Pop() interface{}
}

func Init(h Interface) {
	n := h.Len()
	for i := n / 2 - 1; i >= 0; i-- {
		down(h, i, n)
	}
}

