/**
 * @Author: liubaoshuai3
 * @Date: 2020/3/26 20:36
 * @File: graph
 * @Description:
 */

package tools

import (
	"container/list"
	"fmt"
)

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
	for len(queue) > 0 && !isFound {
		top := queue[0]
		queue = queue[1:]
		linkedlist := g.adj[top]
		for e := linkedlist.Front(); e != nil; e = e.Next() {
			k := e.Value.(int)
			if !visited[k] {
				pre[k] = top
				if k == t {
					isFound = true
					break
				}
				queue = append(queue, k)
				visited[k] = true
			}
		}
	}
	if isFound {
		printPrev(pre, s, t)
	} else {
		fmt.Printf("no path found from %d to %d\n", s, t)
	}
}

// search path by dfs
func (g *Graph) DFS(s int, t int) {
	prev := make([]int, g.v)
	for i := range prev {
		prev[i] = -1
	}
	visited := make([]bool, g.v)
	visited[s] = true
	isFound := false
	g.recurse(s, t, prev, visited, isFound)
	printPrev(prev, s, t)
}

// find path recursively
func (g *Graph) recurse(s int, t int, p []int, v []bool, iF bool) {}

// print path recursively
func printPrev(prev []int, s int, t int) {
	if t == s || prev[t] == -1 {
		fmt.Printf("%d ", t)
	} else {
		printPrev(prev, s, prev[t])
		fmt.Printf("%d ", t)
	}

}

