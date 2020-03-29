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

// add an edge
func (g *Graph) AddEdge(s int, t int) {
	g.adj[s].PushBack(t)
	g.adj[t].PushBack(s)
}

// search path by bfs
func (g *Graph) BFS(s int, t int) {
	// todo
	if s == t {
		return
	}
	pre := make([]int, g.v)
	for i := range pre {
		pre[i] = -1
	}
	var queue []int
	visited := make([]bool, g.v)
	queue = append(queue, s)
	visited[s] = true
	isFound := false
	for len(queue) > 0 && !isFound{
		
	}
}

