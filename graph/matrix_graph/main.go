package main

import (
	"fmt"
	"strconv"
)

func main() {
	m := make(Matrix, 0)

	for i := 0; i < 5; i++ {
		m = append(m, []int{0})
		for j := 0; j < 5; j++ {
			row := m[i]
			row = append(row, 0)
			m[i] = row
		}
	}

	m[0][0] = 1
	m[0][1] = 1
	m[1][0] = 1
	m[2][1] = 1
	m[2][0] = 1
	m[2][2] = 1
	m[4][0] = 1
	m[4][1] = 1
	m[4][2] = 1

	m[0][5] = 1
	m[1][5] = 1

	for _, row := range m {
		fmt.Println(row)
	}

	fmt.Println("---")
	visited := make(map[string]struct{})
	var rivers []int
	for i := 0; i < len(m); i++ {
		for j := 0; j < len(m[i]); j++ {
			res := traverse(m, i, j, visited)
			if res == 0 {
				continue
			}
			rivers = append(rivers, res)
		}
	}

	fmt.Println(rivers)
}

type Matrix [][]int

func traverse(m Matrix, i, j int, visited map[string]struct{}) int {
	if i == -1 {
		return 0
	}

	if j == -1 {
		return 0
	}

	v1, v2 := strconv.Itoa(i), strconv.Itoa(j)
	v := v1 + v2
	_, ok := visited[v]

	if ok {
		return 0
	}

	if i+1 > len(m) {
		return 0
	}

	if j+1 > len(m[i]) {
		return 0
	}

	if m[i][j] == 0 {
		return 0
	}

	visited[v] = struct{}{}
	return traverse(m, i+1, j, visited) +
		traverse(m, i, j+1, visited) +
		traverse(m, i-1, j, visited) +
		traverse(m, i, j-1, visited) + 1
}
