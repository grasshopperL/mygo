/**
 * @Author: liubaoshuai3
 * @Date: 2020/3/26 20:36
 * @File: graph
 * @Description:
 */

package tools

import "container/list"

// the struct of graph
type Graph struct {
	adj []*list.List
	v int
}

// init a graph
func NewGraph(c int) *Graph {
	g := &Graph{
		adj: make([]*list.List, c),
		v:   c,
	}
	for i := range g.adj {
		g.adj[i] = list.New()
	}
	return g
}