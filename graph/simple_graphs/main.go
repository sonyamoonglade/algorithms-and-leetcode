package main

import "fmt"

func main() {
	g := make(Graph)

	g.add("a", "b")
	g.add("b", "c")
	g.add("c", "k")
	g.add("k", "l")
	g.add("l", "g")
	g.add("g", "l")
	g.add("l", "a")
	g.add("z", "v")

	// recursivDfs(g, "a")
	// bfs(g, "a")
	// undirectedGraphDfs(g, "a")
	// ok := hasPathDirected(g, "a", "v")
	ok := hasPathBfsUndirected(g, "z", "v")
	fmt.Println(ok)

	f := fib(6)
	fmt.Println(f)
}

type Graph map[string][]string

func (g Graph) add(k string, v ...string) {
	if _, ok := g[k]; ok == false {
		g[k] = append([]string{}, v...)
		return
	}
	g[k] = append(g[k], v...)
}

func recursivDfs(g Graph, src string) {
	nbs := g[src]
	for _, n := range nbs {
		recursivDfs(g, n)
	}
}

func bfs(g Graph, src string) {

	//Add initially src
	q := make(Queue, 0)
	q.Add(src)
	for len(q) != 0 {
		node := q.Get()
		nodes := g[node]
		for _, n := range nodes {
			q.Add(n)
		}
	}
}

func undirectedGraphDfs(g Graph, src string) {

	memo := make(map[string]struct{})
	memo[src] = struct{}{}

	q := make(Queue, 0)
	q.Add(src)

	for len(q) != 0 {
		node := q.Get()
		fmt.Println(node)
		nodes := g[node]
		for _, n := range nodes {
			//already visited
			if _, ok := memo[n]; ok == true {
				continue
			}
			q.Add(n)
			memo[n] = struct{}{}
		}
	}

}

func hasPathDirected(g Graph, src, dst string) bool {
	if src == dst {
		return true
	}

	for _, node := range g[src] {
		ok := hasPathDirected(g, node, dst)
		if ok {
			return true
		}
	}
	return false
}

func hasPathBfsUndirected(g Graph, src, dest string) bool {
	if src == dest {
		return true
	}
	memo := make(map[string]struct{})
	memo[src] = struct{}{}

	q := make(Queue, 0)
	q.Add(src)

	for len(q) != 0 {
		node := q.Get()
		nodes := g[node]
		for _, n := range nodes {
			//visited
			if _, ok := memo[n]; ok == true {
				continue
			}

			if n == dest {
				return true
			}

			q.Add(n)
			memo[n] = struct{}{}
		}
	}

	return false
}

type Queue []string

func (q *Queue) Add(v string) {
	*q = append([]string{v}, (*q)...)
}

func (q *Queue) Get() string {
	last := (*q)[len(*q)-1]
	*q = (*q)[:len(*q)-1]
	return last
}

func fib(n uint64) uint64 {
	memo := make([]uint64, n, n)
	return oldFib(n-1, &memo) + oldFib(n-2, &memo)
}
func oldFib(n uint64, memo *[]uint64) uint64 {
	if n == 1 || n == 2 {
		return 1
	}

	if v := (*memo)[n-1]; v != 0 {
		return v
	}

	result := oldFib(n-1, memo) + oldFib(n-2, memo)
	(*memo)[n-1] = result

	return result
}
