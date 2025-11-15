package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
)

// Edge represents a weighted directed edge in the graph.
type Edge struct {
	to     int
	weight int
}

// Graph holds adjacency lists for all nodes.
type Graph struct {
	n     int
	edges [][]Edge
}

// NewGraph creates a graph with n nodes.
func NewGraph(n int) *Graph {
	g := &Graph{
		n:     n,
		edges: make([][]Edge, n),
	}
	return g
}

// AddEdge adds a directed edge u -> v with given weight.
func (g *Graph) AddEdge(u, v, w int) {
	g.edges[u] = append(g.edges[u], Edge{to: v, weight: w})
}

// Item represents a node in the priority queue for Dijkstra.
type Item struct {
	node int
	dist int
	idx  int
}

// A PriorityQueue implements heap.Interface for Items.
type PriorityQueue []*Item

func (pq PriorityQueue) Len() int {
	return len(pq)
}

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].dist < pq[j].dist
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].idx = i
	pq[j].idx = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	item := x.(*Item)
	item.idx = len(*pq)
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	item.idx = -1
	*pq = old[0 : n-1]
	return item
}

// dijkstra computes shortest paths from a single source to all nodes.
// It returns the distance array and a parent array to reconstruct paths.
func dijkstra(g *Graph, src int) ([]int, []int) {
	const INF = int(1e18)

	dist := make([]int, g.n)
	parent := make([]int, g.n)
	for i := 0; i < g.n; i++ {
		dist[i] = INF
		parent[i] = -1
	}

	dist[src] = 0

	pq := &PriorityQueue{}
	heap.Init(pq)
	heap.Push(pq, &Item{node: src, dist: 0})

	for pq.Len() > 0 {
		item := heap.Pop(pq).(*Item)
		u := item.node
		if item.dist != dist[u] {
			// Outdated entry.
			continue
		}

		for _, e := range g.edges[u] {
			v := e.to
			nd := dist[u] + e.weight
			if nd < dist[v] {
				dist[v] = nd
				parent[v] = u
				heap.Push(pq, &Item{node: v, dist: nd})
			}
		}
	}

	return dist, parent
}

// reconstructPath returns the shortest path from src to dst using parent array.
func reconstructPath(parent []int, src, dst int) []int {
	if parent[dst] == -1 && src != dst {
		// No path.
		return nil
	}
	path := []int{}
	cur := dst
	for cur != -1 {
		path = append(path, cur)
		if cur == src {
			break
		}
		cur = parent[cur]
	}
	// Reverse path.
	for i, j := 0, len(path)-1; i < j; i, j = i+1, j-1 {
		path[i], path[j] = path[j], path[i]
	}
	if len(path) == 0 || path[0] != src {
		return nil
	}
	return path
}

// Fast input reader.
type FastScanner struct {
	r *bufio.Reader
}

func NewFastScanner() *FastScanner {
	return &FastScanner{r: bufio.NewReader(os.Stdin)}
}

func (fs *FastScanner) NextInt() int {
	sign := 1
	val := 0
	c, _ := fs.r.ReadByte()
	for (c < '0' || c > '9') && c != '-' {
		c, _ = fs.r.ReadByte()
	}
	if c == '-' {
		sign = -1
		c, _ = fs.r.ReadByte()
	}
	for c >= '0' && c <= '9' {
		val = val*10 + int(c-'0')
		c, _ = fs.r.ReadByte()
	}
	return sign * val
}

func main() {
	// Problem:
	// You are given a directed weighted graph with n nodes (1..n) and m edges.
	// Then you are given q queries, each asking for the shortest distance
	// and one actual shortest path between node s and t.
	//
	// Input format:
	// n m q
	// m lines: u v w
	// q lines: s t
	//
	// Output for each query:
	// If there is no path, print:
	//   -1
	// Otherwise print two lines:
	//   d
	//   k a1 a2 ... ak
	// where d is the shortest distance, and a1..ak is one shortest path.
	//
	// This program runs Dijkstra from every distinct source that appears
	// in the queries and caches the result.
	//
	// This is not the most optimal approach for every constraint setup,
	// but it is simple and educational.

	fs := NewFastScanner()
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	n := fs.NextInt()
	m := fs.NextInt()
	q := fs.NextInt()

	g := NewGraph(n)

	for i := 0; i < m; i++ {
		u := fs.NextInt() - 1
		v := fs.NextInt() - 1
		w := fs.NextInt()
		g.AddEdge(u, v, w)
	}

	type Query struct {
		s int
		t int
	}
	queries := make([]Query, q)
	for i := 0; i < q; i++ {
		s := fs.NextInt() - 1
		t := fs.NextInt() - 1
		queries[i] = Query{s: s, t: t}
	}

	distCache := make(map[int][]int)
	parentCache := make(map[int][]int)

	for _, query := range queries {
		src := query.s
		dst := query.t

		if _, ok := distCache[src]; !ok {
			d, p := dijkstra(g, src)
			distCache[src] = d
			parentCache[src] = p
		}

		dist := distCache[src]
		parent := parentCache[src]

		const INF = int(1e18)
		if dist[dst] >= INF {
			fmt.Fprintln(out, -1)
			continue
		}

		path := reconstructPath(parent, src, dst)
		if path == nil {
			fmt.Fprintln(out, -1)
			continue
		}

		fmt.Fprintln(out, dist[dst])
		fmt.Fprint(out, len(path))
		for _, v := range path {
			fmt.Fprint(out, " ", v+1)
		}
		fmt.Fprintln(out)
	}
}
