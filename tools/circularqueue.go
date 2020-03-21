/**
 * @Author: liubaoshuai3
 * @Date: 2020/3/21 21:01
 * @File: circularqueue
 * @Description:
 */

package tools

// struct circular queue
type CircularQueue struct {
	q interface{}
	h int
	t int
	c int
}
