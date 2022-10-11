package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {

	fmt.Printf("'%s'\n", solution("2000 10003 1234000 44444444 9999 11 11 22 123"))

}

func solution(inp string) string {

	spl := strings.Split(inp, " ")
	memo := make(map[int][]string)

	var sortedByValues []int
	for _, s := range spl {
		v := parseInt(s)
		sortedByValues = append(sortedByValues, v)
		memo[v] = append(memo[v], s)
	}

	sort.Ints(sortedByValues)

	done := make(map[int]struct{})
	var out string
	for _, sorted := range sortedByValues {
		for k, memov := range memo {

			if _, ok := done[sorted]; ok == true {
				continue
			}

			if sorted == k {

				sort.Strings(memov)
				for _, mv := range memov {
					out += mv + " "
				}

				done[sorted] = struct{}{}
			}
		}
	}

	return strings.TrimSpace(out)
}

func parseInt(s string) int {
	spl := strings.Split(s, "")
	var n int
	for _, s := range spl {
		res, err := strconv.ParseInt(s, 10, 64)
		if err != nil {
			panic(err)
		}
		n += int(res)
	}
	return n
}
